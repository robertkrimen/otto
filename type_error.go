package otto

func (rt *_runtime) newErrorObject(name string, message Value, stackFramesToPop int) *_object {
	self := rt.newClassObject(classError)
	if message.IsDefined() {
		msg := message.string()
		self.defineProperty("message", toValue_string(msg), 0111, false)
		self.value = newError(rt, name, stackFramesToPop, msg)
	} else {
		self.value = newError(rt, name, stackFramesToPop)
	}

	self.defineOwnProperty("stack", _property{
		value: _propertyGetSet{
			rt.newNativeFunction("get", "internal", 0, func(FunctionCall) Value {
				return toValue_string(self.value.(_error).formatWithStack())
			}),
			&_nilGetSetObject,
		},
		mode: modeConfigureMask & modeOnMask,
	}, false)

	return self
}

func (rt *_runtime) newErrorObjectError(err _error) *_object {
	self := rt.newClassObject(classError)
	self.defineProperty("message", err.messageValue(), 0111, false)
	self.value = err
	switch err.name {
	case "EvalError":
		self.prototype = rt.global.EvalErrorPrototype
	case "TypeError":
		self.prototype = rt.global.TypeErrorPrototype
	case "RangeError":
		self.prototype = rt.global.RangeErrorPrototype
	case "ReferenceError":
		self.prototype = rt.global.ReferenceErrorPrototype
	case "SyntaxError":
		self.prototype = rt.global.SyntaxErrorPrototype
	case "URIError":
		self.prototype = rt.global.URIErrorPrototype
	default:
		self.prototype = rt.global.ErrorPrototype
	}

	self.defineOwnProperty("stack", _property{
		value: _propertyGetSet{
			rt.newNativeFunction("get", "internal", 0, func(FunctionCall) Value {
				return toValue_string(self.value.(_error).formatWithStack())
			}),
			&_nilGetSetObject,
		},
		mode: modeConfigureMask & modeOnMask,
	}, false)

	return self
}
