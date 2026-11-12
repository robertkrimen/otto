package otto

import (
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// Number

func numberValueFromNumberArgumentList(argumentList []Value) Value {
	if len(argumentList) > 0 {
		return argumentList[0].numberValue()
	}
	return intValue(0)
}

func builtinNumber(call FunctionCall) Value {
	return numberValueFromNumberArgumentList(call.ArgumentList)
}

func builtinNewNumber(obj *object, argumentList []Value) Value {
	return objectValue(obj.runtime.newNumber(numberValueFromNumberArgumentList(argumentList)))
}

func builtinNumberToString(call FunctionCall) Value {
	// Will throw a TypeError if ThisObject is not a Number
	value := call.thisClassObject(classNumberName).primitiveValue()
	radix := 10
	radixArgument := call.Argument(0)
	if radixArgument.IsDefined() {
		integer := toIntegerFloat(radixArgument)
		if integer < 2 || integer > 36 {
			panic(call.runtime.panicRangeError("toString() radix must be between 2 and 36"))
		}
		radix = int(integer)
	}
	if radix == 10 {
		return stringValue(value.string())
	}
	return stringValue(numberToStringRadix(value, radix))
}

func builtinNumberValueOf(call FunctionCall) Value {
	return call.thisClassObject(classNumberName).primitiveValue()
}

func builtinNumberToFixed(call FunctionCall) Value {
	value := call.This.float64()
	precision := toIntegerFloat(call.Argument(0))
	if 100 < precision || 0 > precision {
		panic(call.runtime.panicRangeError("toFixed() digits must be between 0 and 100"))
	}
	if math.IsNaN(value) {
		return stringValue("NaN")
	}
	if math.Abs(value) >= 1e21 {
		return stringValue(floatToString(value, 64))
	}
	return stringValue(numberToFixed(value, int(precision)))
}

func builtinNumberToExponential(call FunctionCall) Value {
	value := call.This.float64()
	fractionDigits := call.Argument(0)
	precision := toIntegerFloat(fractionDigits)
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return stringValue(floatToString(value, 64))
	}
	if 100 < precision || 0 > precision {
		panic(call.runtime.panicRangeError("toExponential() argument must be between 0 and 100"))
	}
	if fractionDigits.IsUndefined() {
		return stringValue(numberToExponential(value, -1))
	}
	return stringValue(numberToExponential(value, int(precision)))
}

func builtinNumberToPrecision(call FunctionCall) Value {
	value := call.This.float64()
	if call.Argument(0).IsUndefined() {
		return stringValue(call.This.string())
	}
	precision := toIntegerFloat(call.Argument(0))
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return stringValue(floatToString(value, 64))
	}
	if 100 < precision || 1 > precision {
		panic(call.runtime.panicRangeError("toPrecision() argument must be between 1 and 100"))
	}
	return stringValue(numberToPrecision(value, int(precision)))
}

func builtinNumberIsNaN(call FunctionCall) Value {
	if len(call.ArgumentList) < 1 {
		return boolValue(false)
	}
	return boolValue(call.Argument(0).IsNaN())
}

func builtinNumberToLocaleString(call FunctionCall) Value {
	value := call.thisClassObject(classNumberName).primitiveValue()
	locale := call.Argument(0)
	lang := defaultLanguage
	if locale.IsDefined() {
		lang = language.MustParse(locale.string())
	}

	p := message.NewPrinter(lang)
	return stringValue(p.Sprintf("%v", number.Decimal(value.value)))
}
