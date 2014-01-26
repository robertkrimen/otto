package otto

import (
	"regexp"
	"strings"
	"unicode"
)

// Function

func builtinFunction(call FunctionCall) Value {
	return toValue_object(builtinNewFunctionNative(call.runtime, call.ArgumentList))
}

func builtinNewFunction(self *_object, _ Value, argumentList []Value) Value {
	return toValue_object(builtinNewFunctionNative(self.runtime, argumentList))
}

func argumentList2parameterList(argumentList []Value) []string {
	parameterList := make([]string, 0, len(argumentList))
	for _, value := range argumentList {
		tmp := strings.FieldsFunc(toString(value), func(chr rune) bool {
			return chr == ',' || unicode.IsSpace(chr)
		})
		parameterList = append(parameterList, tmp...)
	}
	return parameterList
}

var matchIdentifier = regexp.MustCompile(`^[$_\p{L}][$_\p{L}\d}]*$`)

func builtinNewFunctionNative(runtime *_runtime, argumentList []Value) *_object {
	parameterList := []string(nil)
	bodySource := ""
	argumentCount := len(argumentList)
	if argumentCount > 0 {
		parameterList = argumentList2parameterList(argumentList[0 : argumentCount-1])
		for _, value := range parameterList {
			// TODO: This is extremely hacky... maybe go through the parser directly?
			if !matchIdentifier.MatchString(value) || keywordTable[value] {
				panic(newSyntaxError(value))
			}
		}
		bodySource = toString(argumentList[argumentCount-1])
	}

	parser := newParser()
	parser.lexer.Source = bodySource
	_programNode := parser.ParseAsFunction()
	return runtime.newNodeFunction(_programNode.toFunction(parameterList), runtime.GlobalEnvironment)
}

func builtinFunction_toString(call FunctionCall) Value {
	call.thisClassObject("Function") // Should throw a TypeError unless Function
	return toValue_string("[function]")
}

func builtinFunction_apply(call FunctionCall) Value {
	if !call.This.isCallable() {
		panic(newTypeError())
	}
	this := call.Argument(0)
	if this.IsUndefined() {
		// FIXME Not ECMA5
		this = toValue_object(call.runtime.GlobalObject)
	}
	argumentList := call.Argument(1)
	switch argumentList._valueType {
	case valueUndefined, valueNull:
		return call.thisObject().Call(this, []Value{})
	case valueObject:
	default:
		panic(newTypeError())
	}

	arrayObject := argumentList._object()
	thisObject := call.thisObject()
	length := int64(toUint32(arrayObject.get("length")))
	valueArray := make([]Value, length)
	for index := int64(0); index < length; index++ {
		valueArray[index] = arrayObject.get(arrayIndexToString(index))
	}
	return thisObject.Call(this, valueArray)
}

func builtinFunction_call(call FunctionCall) Value {
	if !call.This.isCallable() {
		panic(newTypeError())
	}
	thisObject := call.thisObject()
	this := call.Argument(0)
	if this.IsUndefined() {
		// FIXME Not ECMA5
		this = toValue_object(call.runtime.GlobalObject)
	}
	if len(call.ArgumentList) >= 1 {
		return thisObject.Call(this, call.ArgumentList[1:])
	}
	return thisObject.Call(this, []Value{})
}

func builtinFunction_bind(call FunctionCall) Value {
	target := call.This
	if !target.isCallable() {
		panic(newTypeError())
	}
	targetObject := target._object()

	this := call.Argument(0)
	argumentList := call.slice(1)
	if this.IsUndefined() {
		// FIXME Do this elsewhere?
		this = toValue_object(call.runtime.GlobalObject)
	}

	return toValue_object(call.runtime.newBoundFunction(targetObject, this, argumentList))
}
