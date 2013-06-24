package otto

import (
	"encoding/hex"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// Global
func builtinGlobal_eval(call FunctionCall) Value {
	source := call.Argument(0)
	if !source.IsString() {
		return source
	}
	program, err := parse(toString(source))
	if err != nil {
		switch err := err.(type) {
		case *_syntaxError, *_error, _error:
			panic(err)
		default:
			panic(&_syntaxError{Message: fmt.Sprintf("%v", err)})
		}
	}
	runtime := call.runtime
	if call.evalHint {
		runtime.EnterEvalExecutionContext(call)
		defer runtime.LeaveExecutionContext()
	}
	returnValue := runtime.evaluate(program)
	if returnValue.isEmpty() {
		return UndefinedValue()
	}
	return returnValue
}

func builtinGlobal_isNaN(call FunctionCall) Value {
	value := toFloat(call.Argument(0))
	return toValue_bool(math.IsNaN(value))
}

func builtinGlobal_isFinite(call FunctionCall) Value {
	value := toFloat(call.Argument(0))
	return toValue_bool(!math.IsNaN(value) && !math.IsInf(value, 0))
}

// radix 3 => 2 (ASCII 50) +47
// radix 11 => A/a (ASCII 65/97) +54/+86
var parseInt_alphabetTable = func() []string {
	table := []string{"", "", "01"}
	for radix := 3; radix <= 36; radix += 1 {
		alphabet := table[radix-1]
		if radix <= 10 {
			alphabet += string(radix + 47)
		} else {
			alphabet += string(radix+54) + string(radix+86)
		}
		table = append(table, alphabet)
	}
	return table
}()

func builtinGlobal_parseInt(call FunctionCall) Value {
	input := strings.TrimSpace(toString(call.Argument(0)))
	if len(input) == 0 {
		return NaNValue()
	}

	radix := int(toInt32(call.Argument(1)))

	sign := int64(1)
	switch input[0] {
	case '+':
		sign = 1
		input = input[1:]
	case '-':
		sign = -1
		input = input[1:]
	}

	strip := true
	if radix == 0 {
		radix = 10
	} else {
		if radix < 2 || radix > 36 {
			return NaNValue()
		} else if radix != 16 {
			strip = false
		}
	}

	switch len(input) {
	case 0:
		return NaNValue()
	case 1:
	default:
		if strip {
			if input[0] == '0' && (input[1] == 'x' || input[1] == 'X') {
				input = input[2:]
				radix = 16
			}
		}
	}

	alphabet := parseInt_alphabetTable[radix]
	if index := strings.IndexFunc(input, func(chr rune) bool {
		return !strings.ContainsRune(alphabet, chr)
	}); index != -1 {
		input = input[0:index]
	}

	value, err := strconv.ParseInt(input, radix, 64)
	if err != nil {
		return NaNValue()
	}
	value *= sign

	return toValue_int64(value)
}

var parseFloat_matchBadSpecial = regexp.MustCompile(`[\+\-]?(?:[Ii]nf$|infinity)`)
var parseFloat_matchValid = regexp.MustCompile(`[0-9eE\+\-\.]|Infinity`)

func builtinGlobal_parseFloat(call FunctionCall) Value {
	// Caveat emptor: This implementation does NOT match the specification
	input := strings.TrimSpace(toString(call.Argument(0)))
	if parseFloat_matchBadSpecial.MatchString(input) {
		return NaNValue()
	}
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		for end := len(input); end > 0; end -= 1 {
			input := input[0:end]
			if !parseFloat_matchValid.MatchString(input) {
				return NaNValue()
			}
			value, err = strconv.ParseFloat(input, 64)
			if err == nil {
				break
			}
		}
		if err != nil {
			return NaNValue()
		}
	}
	return toValue_float64(value)
}

// encodeURI/decodeURI

func _builtinGlobal_encodeURI(call FunctionCall, escape *regexp.Regexp) Value {
	value := call.Argument(0)
	var input []uint16
	switch vl := value.value.(type) {
	case []uint16:
		input = vl
	default:
		input = utf16.Encode([]rune(toString(value)))
	}
	if len(input) == 0 {
		return toValue_string("")
	}
	output := []byte{}
	length := len(input)
	encode := make([]byte, 4)
	for index := 0; index < length; {
		value := input[index]
		decode := utf16.Decode(input[index : index+1])
		if value >= 0xDC00 && value <= 0xDFFF {
			panic(newURIError("URI malformed"))
		}
		if value >= 0xD800 && value <= 0xDBFF {
			index += 1
			if index >= length {
				panic(newURIError("URI malformed"))
			}
			// input = ..., value, value1, ...
			value = value
			value1 := input[index]
			if value1 < 0xDC00 || value1 > 0xDFFF {
				panic(newURIError("URI malformed"))
			}
			decode = []rune{((rune(value) - 0xD800) * 0x400) + (rune(value1) - 0xDC00) + 0x10000}
		}
		index += 1
		size := utf8.EncodeRune(encode, decode[0])
		encode := encode[0:size]
		output = append(output, encode...)
	}
	{
		value := escape.ReplaceAllFunc(output, func(target []byte) []byte {
			// Probably a better way of doing this
			if target[0] == ' ' {
				return []byte("%20")
			}
			return []byte(url.QueryEscape(string(target)))
		})
		return toValue_string(string(value))
	}
}

var encodeURI_Regexp = regexp.MustCompile(`([^~!@#$&*()=:/,;?+'])`)

func builtinGlobal_encodeURI(call FunctionCall) Value {
	return _builtinGlobal_encodeURI(call, encodeURI_Regexp)
}

var encodeURIComponent_Regexp = regexp.MustCompile(`([^~!*()'])`)

func builtinGlobal_encodeURIComponent(call FunctionCall) Value {
	return _builtinGlobal_encodeURI(call, encodeURIComponent_Regexp)
}

// 3B/2F/3F/3A/40/26/3D/2B/24/2C/23
var decodeURI_guard = regexp.MustCompile(`(?i)(?:%)(3B|2F|3F|3A|40|26|3D|2B|24|2C|23)`)

func _decodeURI(input string, reserve bool) (string, bool) {
	if reserve {
		input = decodeURI_guard.ReplaceAllString(input, "%25$1")
	}
	input = strings.Replace(input, "+", "%2B", -1) // Ugly hack to make QueryUnescape work with our use case
	output, err := url.QueryUnescape(input)
	if err != nil || !utf8.ValidString(output) {
		return "", true
	}
	return output, false
}

func builtinGlobal_decodeURI(call FunctionCall) Value {
	output, err := _decodeURI(toString(call.Argument(0)), true)
	if err {
		panic(newURIError("URI malformed"))
	}
	return toValue_string(output)
}

func builtinGlobal_decodeURIComponent(call FunctionCall) Value {
	output, err := _decodeURI(toString(call.Argument(0)), false)
	if err {
		panic(newURIError("URI malformed"))
	}
	return toValue_string(output)
}

// escape/unescape

func builtin_shouldEscape(chr byte) bool {
	if 'A' <= chr && chr <= 'Z' || 'a' <= chr && chr <= 'z' || '0' <= chr && chr <= '9' {
		return false
	}
	return !strings.ContainsRune("*_+-./", rune(chr))
}

const escapeBase16 = "0123456789ABCDEF"

func builtin_escape(input string) string {
	output := make([]byte, 0, len(input))
	length := len(input)
	for index := 0; index < length; {
		if builtin_shouldEscape(input[index]) {
			chr, width := utf8.DecodeRuneInString(input[index:])
			chr16 := utf16.Encode([]rune{chr})[0]
			if 256 > chr16 {
				output = append(output, '%',
					escapeBase16[chr16>>4],
					escapeBase16[chr16&15],
				)
			} else {
				output = append(output, '%', 'u',
					escapeBase16[chr16>>12],
					escapeBase16[(chr16>>8)&15],
					escapeBase16[(chr16>>4)&15],
					escapeBase16[chr16&15],
				)
			}
			index += width

		} else {
			output = append(output, input[index])
			index += 1
		}
	}
	return string(output)
}

func builtin_unescape(input string) string {
	output := make([]rune, 0, len(input))
	length := len(input)
	for index := 0; index < length; {
		if input[index] == '%' {
			if index <= length-6 && input[index+1] == 'u' {
				byte16, err := hex.DecodeString(input[index+2 : index+6])
				if err == nil {
					value := uint16(byte16[0])<<8 + uint16(byte16[1])
					chr := utf16.Decode([]uint16{value})[0]
					output = append(output, chr)
					index += 6
					continue
				}
			}
			if index <= length-3 {
				byte8, err := hex.DecodeString(input[index+1 : index+3])
				if err == nil {
					value := uint16(byte8[0])
					chr := utf16.Decode([]uint16{value})[0]
					output = append(output, chr)
					index += 3
					continue
				}
			}
		}
		output = append(output, rune(input[index]))
		index += 1
	}
	return string(output)
}

func builtinGlobal_escape(call FunctionCall) Value {
	return toValue_string(builtin_escape(toString(call.Argument(0))))
}

func builtinGlobal_unescape(call FunctionCall) Value {
	return toValue_string(builtin_unescape(toString(call.Argument(0))))
}

// Error

func builtinError(call FunctionCall) Value {
	return toValue_object(call.runtime.newError("", call.Argument(0)))
}

func builtinNewError(self *_object, _ Value, argumentList []Value) Value {
	return toValue_object(self.runtime.newError("", valueOfArrayIndex(argumentList, 0)))
}

func builtinError_toString(call FunctionCall) Value {
	thisObject := call.thisObject()
	if thisObject == nil {
		panic(newTypeError())
	}

	name := "Error"
	nameValue := thisObject.get("name")
	if nameValue.IsDefined() {
		name = toString(nameValue)
	}

	message := ""
	messageValue := thisObject.get("message")
	if messageValue.IsDefined() {
		message = toString(messageValue)
	}

	if len(name) == 0 {
		return toValue_string(message)
	}

	if len(message) == 0 {
		return toValue_string(name)
	}

	return toValue_string(fmt.Sprintf("%s: %s", name, message))
}
