package otto

import (
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode/utf16"
)

func builtinJSON_parse(call FunctionCall) Value {
	var walk func(Value, string) Value
	text, err := prepareJSON(call.Argument(0))
	reviver := call.Argument(1)
	var result Value

	if nil != err {
		panic(newSyntaxError(err.Error()))
	}

	{
		p, err := builtinJSON_Parse(text)
		if nil != err {
			switch err := err.(type) {
			case *_syntaxError, *_error, _error:
				panic(err)
			default:
				panic(newSyntaxError(fmt.Sprintf("%v", err)))
			}
		}
		call.runtime.EnterEvalExecutionContext(call)
		result = call.runtime.evaluate(p)
		if result.isEmpty() {
			result = UndefinedValue()
		}
		call.runtime.LeaveExecutionContext()
	}

	walk = func(holder Value, key string) Value {
		value := holder._object().get(key)
		if !value.IsUndefined() && value.IsObject() {
			value._object().enumerate(false, func(name string) bool {
				if builtinJSON_hasOwnProperty(call, value, name) {
					v := walk(value, name)
					if v.IsUndefined() {
						value._object().delete(name, false)
					} else {
						value._object().put(name, v, false)
					}
				}
				return true
			})
		}
		return reviver.call(holder, key, value)
	}

	if reviver.isCallable() {
		obj := newObject(call.runtime, "Object")
		obj.put("", result, true)
		return walk(toValue(obj), "")
	}
	return result
}

func builtinJSON_stringify(call FunctionCall) Value {
	var str func(string, Value) Value
	value := call.Argument(0)
	replacer := call.Argument(1)
	space := call.Argument(2)
	stack := []*_object{nil}
	indent, gap := "", ""

	if !space.IsUndefined() {
		if space.IsNumber() {
			i, _ := space.ToInteger()
			indent = strings.Repeat(" ", int(i))
		}
		if space.IsString() {
			indent = space.String()
		}
	}

	str = func(key string, holder Value) Value {
		value := holder._object().get(key)
		mind := gap
		defer func() { gap = mind }()

		if value.IsObject() && value._object().get("toJSON").IsFunction() {
			value = value._object().get("toJSON").call(value, key)
		}

		if replacer.isCallable() {
			value = replacer.call(holder, key, value)
		}

		switch {
		case value.IsString():
			return toValue(builtinJSON_quote(value.String()))
		case value.IsNumber():
			v := value.toFloat()
			if !math.IsNaN(v) && !math.IsInf(v, 0) {
				return toValue(value.String())
			} else {
				return toValue("null")
			}
		case value.IsBoolean(), value.IsNull():
			return toValue(value.String())
		case value.IsObject():
			if value := value._object(); nil != value {
				for _, object := range stack {
					if value == object {
						panic(newTypeError("Converting circular structure to JSON"))
					}
				}
				stack = append(stack, value)
			}
			partial := []string{}
			obj := value._object()
			gap += indent
			res := ""
			if value.isArray() {
				obj.enumerate(false, func(k string) bool {
					partial = append(partial, str(k, value).String())
					return true
				})
				res = "[]"
				if len(partial) > 0 {
					if "" != gap {
						res = "[\n" + gap + strings.Join(partial, ",\n" + gap) + "\n" + mind + "]"
					} else {
						res = "[" + strings.Join(partial, ",") + "]"
					}
				}
				return toValue(res)
			}
			if replacer.isArray() {
				obj := replacer._object()
				obj.enumerate(false, func(i string) bool {
					if k := obj.get(i); k.IsString() {
						v := builtinJSON_quote(k.String()) + ":"
						if "" != gap {
							v += " "
						}
						partial = append(partial, v + str(k.String(), value).String())
					}
					return true
				})
			} else {
				obj.enumerate(false, func(k string) bool {
					if builtinJSON_hasOwnProperty(call, value, k) {
						v := builtinJSON_quote(k) + ":"
						if "" != gap {
							v += " "
						}
						partial = append(partial, v + str(k, value).String())
					}
					return true
				})
			}
			res = "{}"
			if len(partial) > 0 {
				if "" != gap {
					res = "{\n" + gap + strings.Join(partial, ",\n" + gap) + "\n" + mind + "}"
				} else {
					res = "{" + strings.Join(partial, ",") + "}"
				}
			}
			return toValue(res)
		}
		return UndefinedValue()
	}
	obj := newObject(call.runtime, "Object")
	obj.put("", value, true)
	return str("", toValue(obj))
}

var builtinJSON_prepare_cx = regexp.MustCompile(
	`[\x{0000}\x{00ad}\x{0600}-\x{0604}\x{070f}\x{17b4}\x{17b5}\x{200c}-\x{200f}\x{2028}-\x{202f}\x{2060}-\x{206f}\x{feff}\x{fff0}-\x{ffff}]`)

func prepareJSON(v Value) (string, error) {
	var b = []byte(v.toString())
	if builtinJSON_prepare_cx.Match(b) {
		b = builtinJSON_prepare_cx.ReplaceAllFunc(b, func(a []byte) []byte {
			return []byte(builtinJSON_encodeChar(string(a), 0))
		})
		v = toValue(string(b))
	}
	return v.String(), nil
}

var builtinJSON_quote_escapable = regexp.MustCompile(
	`[\\\"\x00-\x1f\x7f-\x9f\x{00ad}\x{0600}-\x{0604}\x{070f}\x{17b4}\x{17b5}\x{200c}-\x{200f}\x{2028}-\x{202f}\x{2060}-\x{206f}\x{feff}\x{fff0}-\x{ffff}]`)

var builtinJSON_quote_chartable = map[string]string{
	"\b": "\\b",
	"\t": "\\t",
	"\n": "\\n",
	"\f": "\\f",
	"\r": "\\r",
	`"` : `\\"`,
	"\\": "\\\\",
}

func builtinJSON_quote(s string) string {
	if builtinJSON_quote_escapable.MatchString(s) {
		return builtinJSON_quote_escapable.ReplaceAllStringFunc(s, func(a string) string {
			if char, found := builtinJSON_quote_chartable[a]; found {
				return char
			}
			return builtinJSON_encodeChar(a, 0)
		})
	}
	return `"` + s + `"`
}

func builtinJSON_encodeChar(s string, i int) string {
	si16 := hex.EncodeToString([]byte(string(utf16.Encode([]rune(s))[i])))
	res := "\\u"
	if len(si16) < 4 {
		res += strings.Repeat("0", 4-len(si16)) + si16
	} else {
		res += si16
	}
	return res
}

func builtinJSON_hasOwnProperty(call FunctionCall, obj Value, name string) bool {
	p, err := call.runtime.Global.ObjectPrototype.get("hasOwnProperty").Call(obj, name)
	return (nil == err) && p.toBoolean()
}

type builtinJSON_Parser struct {
	_parser
}

func builtinJSON_NewParser() *builtinJSON_Parser {
	parser := &builtinJSON_Parser{
		_parser{
			history: make([]_token, 0, 4),
		},
	}
	parser.lexer.readIn = make([]rune, 0)
	return parser
}

func builtinJSON_Parse(source string) (result _node, err interface{}) {
	defer func() {
		if caught := recover(); caught != nil {
			switch caught := caught.(type) {
			case *_syntaxError, _error, *_error:
				err = caught
				return
			}
			panic(caught)
		}
	}()
	parser := builtinJSON_NewParser()
	parser.lexer.Source = source
	result = parser.ParsePrimaryExpression()
	return
}

func (self *builtinJSON_Parser) ParsePrimaryExpression() _node {
	token := self.Peek()
	switch token.Kind {
	case "string":
		return self.ConsumeString()
	case "boolean":
		return self.ConsumeBoolean()
	case "number":
		return self.ConsumeNumber()
	case "null":
		return self.ConsumeNull()
	case "{":
		return self.ParseObjectLiteral()
	case "[":
		return self.ParseArrayLiteral()
	case "(":
		self.Expect("(")
		result := self.ParsePrimaryExpression()
		self.Expect(")")
		return result
	}
	panic(self.Unexpected(token))
}

func (self *builtinJSON_Parser) ParseArrayLiteral() *_arrayNode {
	self.Expect("[")
	nodeList := []_node{}
	for !self.Match("]") {
		if self.Accept(",") {
			nodeList = append(nodeList, newEmptyNode())
			continue
		}
		nodeList = append(nodeList, self.ParsePrimaryExpression())
		if self.Accept(",") {
			continue
		}
	}
	self.Expect("]")
	node := newArrayNode(nodeList)
	self.markNode(node)
	return node
}

func (self *builtinJSON_Parser) ParseObjectLiteral() *_objectNode {
	node := newObjectNode()
	self.markNode(node)
	self.Expect("{")
	for !self.Match("}") {
		if !self.Match("string") {
			panic(self.Unexpected(self.Next()))
		}
		key := toString(self.ConsumeString().Value)
		self.Expect(":")
		property := newObjectPropertyNode(key, self.ParsePrimaryExpression())
		node.AddProperty(property)
		self.markNode(property)
		if self.Accept(",") {
			continue
		}
	}
	self.Expect("}")
	return node
}

func (self *builtinJSON_Parser) Unexpected(token _token) *_syntaxError {
	switch token.Kind {
	case "EOF":
		if len(self.history) > 0 {
			return self.History(-1).newSyntaxError("Unexpected end of input")
		}
		return token.newSyntaxError("Unexpected end of input")
	case "illegal":
		return token.newSyntaxError("Unexpected token ILLEGAL (%s)", token.Text)
	}
	return token.newSyntaxError("Unexpected token %s", token.Text)
}
