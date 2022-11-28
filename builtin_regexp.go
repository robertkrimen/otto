package otto

import (
	"fmt"
)

// RegExp

func builtinRegExp(call FunctionCall) Value {
	pattern := call.Argument(0)
	flags := call.Argument(1)
	if object := pattern._object(); object != nil {
		if object.class == classRegExp && flags.IsUndefined() {
			return pattern
		}
	}
	return toValue_object(call.runtime.newRegExp(pattern, flags))
}

func builtinNewRegExp(self *_object, argumentList []Value) Value {
	return toValue_object(self.runtime.newRegExp(
		valueOfArrayIndex(argumentList, 0),
		valueOfArrayIndex(argumentList, 1),
	))
}

func builtinRegExp_toString(call FunctionCall) Value {
	thisObject := call.thisObject()
	source := thisObject.get("source").string()
	flags := []byte{}
	if thisObject.get("global").bool() {
		flags = append(flags, 'g')
	}
	if thisObject.get("ignoreCase").bool() {
		flags = append(flags, 'i')
	}
	if thisObject.get("multiline").bool() {
		flags = append(flags, 'm')
	}
	return toValue_string(fmt.Sprintf("/%s/%s", source, flags))
}

func builtinRegExp_exec(call FunctionCall) Value {
	thisObject := call.thisObject()
	target := call.Argument(0).string()
	match, result := execRegExp(thisObject, target)
	if !match {
		return nullValue
	}
	return toValue_object(execResultToArray(call.runtime, target, result))
}

func builtinRegExp_test(call FunctionCall) Value {
	thisObject := call.thisObject()
	target := call.Argument(0).string()
	match, result := execRegExp(thisObject, target)

	if !match {
		return toValue_bool(match)
	}

	// Match extract and assign input, $_ and $1 -> $9 on global RegExp.
	input := toValue_string(target)
	call.runtime.global.RegExp.defineProperty("$_", input, 0100, false)
	call.runtime.global.RegExp.defineProperty("input", input, 0100, false)

	var start int
	n := 1
	re := call.runtime.global.RegExp
	for i, v := range result[2:] {
		if i%2 == 0 {
			start = v
		} else {
			re.defineProperty(fmt.Sprintf("$%d", n), toValue_string(target[start:v]), 0100, false)
			n++
			if n == 10 {
				break
			}
		}
	}

	if n <= 9 {
		// Erase remaining.
		empty := toValue_string("")
		for i := n; i <= 9; i++ {
			re.defineProperty(fmt.Sprintf("$%d", i), empty, 0100, false)
		}
	}

	return toValue_bool(match)
}

func builtinRegExp_compile(call FunctionCall) Value {
	// This (useless) function is deprecated, but is here to provide some
	// semblance of compatibility.
	// Caveat emptor: it may not be around for long.
	return Value{}
}
