package otto

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
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
		p, err := parse(text)
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
			value._object().enumerate(func(name string) {
				if builtinJSON_hasOwnProperty(call, value, name) {
					v := walk(value, name)
					if v.IsUndefined() {
						value._object().delete(name, false)
					} else {
						value._object().put(name, v, false)
					}
				}
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
			partial := []string{}
			obj := value._object()
			gap += indent
			res := ""
			if value.isArray() {
				obj.enumerate(func(k string) {
					partial = append(partial, str(k, value).String())
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
				obj.enumerate(func(i string) {
					if k := obj.get(i); k.IsString() {
						v := builtinJSON_quote(k.String()) + ":"
						if "" != gap {
							v += " "
						}
						partial = append(partial, v + str(k.String(), value).String())
					}
				})
			} else {
				obj.enumerate(func(k string) {
					if builtinJSON_hasOwnProperty(call, value, k) {
						v := builtinJSON_quote(k) + ":"
						if "" != gap {
							v += " "
						}
						partial = append(partial, v + str(k, value).String())
					}
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

var builtinJSON_prepare_replace = []*regexp.Regexp{
	// Replace the JSON backslash pairs with '@' (a non-JSON character)
	regexp.MustCompile(`\\(?:["\\\/bfnrt]|u[0-9a-fA-F]{4})`),
	// Replace all simple value tokens with ']' characters
	regexp.MustCompile(`"[^"\\\n\r]*"|true|false|null|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?`),
	// Delete all open brackets that follow a colon or comma or that begin the text
	regexp.MustCompile(`(?:^|:|,)(?:\s*\[)+`),
}

// Look to see that the remaining characters are only whitespace or ']' or ',' or ':'
//  or '{' or '}'. If that is so, then the text is safe for eval.
var builtinJSON_prepare_validate = regexp.MustCompile(`^[\],:{}\s]*$`)

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
	b = builtinJSON_prepare_replace[0].ReplaceAll(b, []byte("@"))
	b = builtinJSON_prepare_replace[1].ReplaceAll(b, []byte("]"))
	b = builtinJSON_prepare_replace[2].ReplaceAll(b, []byte(""))
	if builtinJSON_prepare_validate.Match(b) {
		return "(" + v.String() + ")", nil
	}
	return "", errors.New("JSON.parse")
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

func builtinJSON_toJSON(call FunctionCall) Value {
	value := call.thisObject()
	if value.class == "Date" {
		if d := dateObjectOf(value); !d.isNaN {
			f := func(n int) string {
				s := strconv.FormatInt(int64(n), 10)
				if len(s) < 2 {
					s = "0" + s
				}
				return s
			}
			return toValue(
				f(d.Time().Year())       + "-" +
				f(int(d.Time().Month())) + "-" +
				f(d.Time().Day())        + "T" +
				f(d.Time().Hour())       + ":" +
				f(d.Time().Minute())     + ":" +
				f(d.Time().Second())     + "Z")
		}
		return NullValue()
	}
	return value.DefaultValue(
		defaultValueNoHint)
}
