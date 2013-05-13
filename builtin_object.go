package otto

import (
	"fmt"
)

// Object

func builtinObject(call FunctionCall) Value {
	value := call.Argument(0)
	switch value._valueType {
	case valueUndefined, valueNull:
		return toValue(call.runtime.newObject())
	}

	return toValue(call.runtime.toObject(value))
}

func builtinNewObject(self *_object, _ Value, argumentList []Value) Value {
	value := valueOfArrayIndex(argumentList, 0)
	switch value._valueType {
	case valueNull, valueUndefined:
	case valueNumber, valueString, valueBoolean:
		return toValue(self.runtime.toObject(value))
	case valueObject:
		return value
	default:
	}
	return toValue(self.runtime.newObject())
}

func builtinObject_toString(call FunctionCall) Value {
	result := ""
	if call.This.IsUndefined() {
		result = "[object Undefined]"
	} else if call.This.IsNull() {
		result = "[object Null]"
	} else {
		result = fmt.Sprintf("[object %s]", call.thisObject().class)
	}
	return toValue(result)
}

func builtinObject_toLocaleString(call FunctionCall) Value {
	toString := call.thisObject().get("toString")
	if !toString.isCallable() {
		panic(newTypeError())
	}
	return toString.call(call.This)
}

func builtinObject_getPrototypeOf(call FunctionCall) Value {
	objectValue := call.Argument(0)
	object := objectValue._object()
	if object == nil {
		panic(newTypeError())
	}

	if object.prototype == nil {
		return NullValue()
	}

	return toValue(object.prototype)
}

func builtinObject_getOwnPropertyDescriptor(call FunctionCall) Value {
	objectValue := call.Argument(0)
	object := objectValue._object()
	if object == nil {
		panic(newTypeError())
	}

	name := toString(call.Argument(1))
	descriptor := object.getOwnProperty(name)
	if descriptor == nil {
		return UndefinedValue()
	}
	return toValue(call.runtime.fromPropertyDescriptor(*descriptor))
}

func builtinObject_getOwnPropertyNames(call FunctionCall) Value {
	if object := call.Argument(0)._object(); nil != object {
		var walk func(_stash, func(string))
		var names []Value

		walk = func(stash _stash, each func(string)) {
			switch stash := stash.(type) {
			case *_objectStash:
				for _, name := range stash._order {
					each(name)
				}
			case *_arrayStash:
				for index, _ := range stash.valueArray {
					if stash.valueArray[index]._valueType == valueEmpty {
						continue // A sparse array
					}
					each(arrayIndexToString(uint(index)))
				}
				walk(stash._stash, each)
			}
		}

		walk(object.stash, func(name string) {
			if p := object.getOwnProperty(name); nil != p {
				names = append(names, toValue(name))
			}
		})
		return toValue(call.runtime.newArray(names))
	}
	panic(newTypeError())
}

func builtinObject_defineProperty(call FunctionCall) Value {
	objectValue := call.Argument(0)
	object := objectValue._object()
	if object == nil {
		panic(newTypeError())
	}
	name := toString(call.Argument(1))
	descriptor := toPropertyDescriptor(call.Argument(2))
	object.defineOwnProperty(name, descriptor, true)
	return objectValue
}

func builtinObject_defineProperties(call FunctionCall) Value {
	objectValue := call.Argument(0)
	object := objectValue._object()
	if object == nil {
		panic(newTypeError())
	}

	properties := call.runtime.toObject(call.Argument(1))
	properties.enumerate(func(name string) {
		descriptor := toPropertyDescriptor(properties.get(name))
		object.defineOwnProperty(name, descriptor, true)
	})

	return objectValue
}

func builtinObject_create(call FunctionCall) Value {
	prototypeValue := call.Argument(0)
	if !prototypeValue.IsNull() && !prototypeValue.IsObject() {
		panic(newTypeError())
	}

	object := call.runtime.newObject()
	object.prototype = prototypeValue._object()

	propertiesValue := call.Argument(1)
	if propertiesValue.IsDefined() {
		properties := call.runtime.toObject(propertiesValue)
		properties.enumerate(func(name string) {
			descriptor := toPropertyDescriptor(properties.get(name))
			object.defineOwnProperty(name, descriptor, true)
		})
	}

	return toValue(object)
}

func builtinObject_keys(call FunctionCall) Value {
	if object, elements := call.Argument(0)._object(), []Value{}; nil != object {
		object.enumerate(func(name string) {
			elements = append(elements, toValue(name))
		})
		return toValue(call.runtime.newArray(elements))
	}
	panic(newTypeError())
}

func builtinObject_isExtensible(call FunctionCall) Value {
	object := call.Argument(0)
	if object := object._object(); object != nil {
		return toValue(object.extensible)
	}
	panic(newTypeError())
}

func builtinObject_preventExtensions(call FunctionCall) Value {
	object := call.Argument(0)
	if object := object._object(); object != nil {
		object.extensible = false
	} else {
		panic(newTypeError())
	}
	return object
}

func builtinObject_isSealed(call FunctionCall) Value {
	object := call.Argument(0)
	if object := object._object(); object != nil {
		if object.extensible {
			return toValue(false)
		}
		result := true
		object.enumerate(func(name string) {
			property := object.getProperty(name)
			if property.configurable() {
				result = false
			}
		})
		return toValue(result)
	}
	panic(newTypeError())
}

func builtinObject_seal(call FunctionCall) Value {
	object := call.Argument(0)
	if object := object._object(); object != nil {
		object.enumerate(func(name string) {
			if p := object.getOwnProperty(name);
					nil != p && p.configurable() {
				p.mode &= ^propertyMode_configure
				object.defineOwnProperty(name, *p, true)
			}
		})
		object.stash.lock()
	} else {
		panic(newTypeError())
	}
	return object
}

func builtinObject_isFrozen(call FunctionCall) Value {
	object := call.Argument(0)
	if object := object._object(); object != nil {
		if object.extensible {
			return toValue(false)
		}
		result := true
		object.enumerate(func(name string) {
			property := object.getProperty(name)
			if property.configurable() || property.writable() {
				result = false
			}
		})
		return toValue(result)
	}
	panic(newTypeError())
}

func builtinObject_freeze(call FunctionCall) Value {
	object := call.Argument(0)
	if object := object._object(); object != nil {
		object.enumerate(func(name string) {
			if property, update := object.getOwnProperty(name), false; nil != property {
				if property.isDataDescriptor() && property.writable() {
					property.mode &= ^propertyMode_write
					update = true
				}
				if property.configurable() {
					property.mode &= ^propertyMode_configure
					update = true
				}
				if update {
					object.defineOwnProperty(name, *property, true)
				}
			}
		})
		object.extensible = false
	} else {
		panic(newTypeError())
	}
	return object
}

func builtinObject_defineGetterSetter(index int) func(FunctionCall) Value {
	if index == 0 || index == 1 {
		return func(call FunctionCall) Value {
			if fn := call.Argument(1); fn.isCallable() {
				this, key := call.thisObject(), call.Argument(0).String()
				getset := _propertyGetSet{nil, nil}
				getset[index] = fn._object()
				if p := this.getProperty(key); nil != p {
					if current, test := p.value.(_propertyGetSet); test {
						if index == 0 {
							index = 1
						} else {
							index = 0
						}
						getset[index] = current[index]
					}
				}
				descriptor := _property{getset, _propertyMode(0011)}
				this.defineOwnProperty(key, descriptor, false)
				return UndefinedValue()
			}
			panic(newTypeError())
		}
	}
	panic(hereBeDragons())
}

func builtinObject_lookupGetterSetter(index int) func(FunctionCall) Value {
	if index == 0 || index == 1 {
		return func(call FunctionCall) Value {
			if p := call.thisObject().getProperty(call.Argument(0).String()); nil != p {
				if getset, test := p.value.(_propertyGetSet); test {
					return toValue(getset[index])
				}
			}
			return UndefinedValue()
		}
	}
	panic(hereBeDragons())
}
