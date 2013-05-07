package otto

import (
	Time "time"
)

// Date

func builtinDate(call FunctionCall) Value {
	date := &_dateObject{}
	if len(call.ArgumentList) == 0 {
		// TODO Should make this prettier
		date.Set(newDateTime([]Value{}, Time.Local))
		return toValue(date.Time().Format(Time.RFC1123))
	}
	date.Set(newDateTime(call.ArgumentList, Time.Local))
	return toValue(date.Time().Local().Format(Time.RFC1123))
}

func builtinNewDate(self *_object, _ Value, argumentList []Value) Value {
	return toValue(self.runtime.newDate(newDateTime(argumentList, Time.Local)))
}

func builtinDate_toString(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return toValue("Invalid Date")
	}
	return toValue(date.Time().Local().Format(Time.RFC1123))
}

func builtinDate_toUTCString(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return toValue("Invalid Date")
	}
	return toValue(date.Time().Format(Time.RFC1123))
}

func builtinDate_toGMTString(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return toValue("Invalid Date")
	}
	return toValue(date.Time().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
}

func builtinDate_getTime(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return NaNValue()
	}
	// We do this (convert away from a float) so the user
	// does not get something back in exponential notation
	return toValue(int64(date.Epoch()))
}

func builtinDate_setTime(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	date.Set(toFloat(call.Argument(0)))
	return date.Value()
}

func _builtinDate_set(call FunctionCall, argumentCap int, dateLocal bool) (*_dateObject, *_ecmaTime) {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return nil, nil
	}
	for index := 0; index < len(call.ArgumentList) && index < argumentCap; index++ {
		value := call.Argument(index)
		if value.IsNaN() {
			date.SetNaN()
			return date, nil
		}
	}
	baseTime := date.Time()
	if dateLocal {
		baseTime = baseTime.Local()
	}
	ecmaTime := ecmaTime(baseTime)
	return date, &ecmaTime
}

func builtinDate_parse(call FunctionCall) Value {
	date := toString(call.Argument(0))
	return toValue(dateParse(date))
}

func builtinDate_UTC(call FunctionCall) Value {
	return toValue(newDateTime(call.ArgumentList, Time.UTC))
}

func builtinDate_now(call FunctionCall) Value {
	call.ArgumentList = []Value{}
	return builtinDate_UTC(call)
}

// This is a placeholder
func builtinDate_toLocaleString(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return toValue("Invalid Date")
	}
	return toValue(date.Time().Local().Format("2006-01-02 15:04:05"))
}

// This is a placeholder
func builtinDate_toLocaleDateString(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return toValue("Invalid Date")
	}
	return toValue(date.Time().Local().Format("2006-01-02"))
}

// This is a placeholder
func builtinDate_toLocaleTimeString(call FunctionCall) Value {
	date := dateObjectOf(call.thisObject())
	if date.isNaN {
		return toValue("Invalid Date")
	}
	return toValue(date.Time().Local().Format("15:04:05"))
}
