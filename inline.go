package otto

import (
	"math"
)

func (rt *runtime) newContext() {
	{
		rt.global.ObjectPrototype = &object{
			runtime:     rt,
			class:       classObjectName,
			objectClass: classObject,
			prototype:   nil,
			extensible:  true,
			value:       prototypeValueObject,
		}

		rt.global.FunctionPrototype = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       prototypeValueFunction,
		}

		valueOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "valueOf",
				call: builtinObjectValueOf,
			},
		}
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinObjectToString,
			},
		}
		toLocaleString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleString",
				call: builtinObjectToLocaleString,
			},
		}
		hasOwnProperty := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "hasOwnProperty",
				call: builtinObjectHasOwnProperty,
			},
		}
		isPrototypeOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isPrototypeOf",
				call: builtinObjectIsPrototypeOf,
			},
		}
		propertyIsEnumerable := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "propertyIsEnumerable",
				call: builtinObjectPropertyIsEnumerable,
			},
		}
		rt.global.ObjectPrototype.property = map[string]property{
			"valueOf": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: valueOf,
				},
			},
			"toString": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: toString,
				},
			},
			"toLocaleString": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: toLocaleString,
				},
			},
			"hasOwnProperty": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: hasOwnProperty,
				},
			},
			"isPrototypeOf": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: isPrototypeOf,
				},
			},
			"propertyIsEnumerable": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: propertyIsEnumerable,
				},
			},
			"constructor": {
				mode:  0o101,
				value: Value{},
			},
		}
		rt.global.ObjectPrototype.propertyOrder = []string{
			"valueOf",
			"toString",
			"toLocaleString",
			"hasOwnProperty",
			"isPrototypeOf",
			"propertyIsEnumerable",
			"constructor",
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinFunctionToString,
			},
		}
		apply := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "apply",
				call: builtinFunctionApply,
			},
		}
		call := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "call",
				call: builtinFunctionCall,
			},
		}
		bind := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "bind",
				call: builtinFunctionBind,
			},
		}
		rt.global.FunctionPrototype.property = map[string]property{
			"toString": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: toString,
				},
			},
			"apply": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: apply,
				},
			},
			"call": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: call,
				},
			},
			"bind": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: bind,
				},
			},
			"constructor": {
				mode:  0o101,
				value: Value{},
			},
			propertyLength: {
				mode: 0,
				value: Value{
					kind:  valueNumber,
					value: 0,
				},
			},
		}
		rt.global.FunctionPrototype.propertyOrder = []string{
			"toString",
			"apply",
			"call",
			"bind",
			"constructor",
			propertyLength,
		}
	}
	{
		getPrototypeOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getPrototypeOf",
				call: builtinObjectGetPrototypeOf,
			},
		}
		getOwnPropertyDescriptor := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getOwnPropertyDescriptor",
				call: builtinObjectGetOwnPropertyDescriptor,
			},
		}
		defineProperty := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 3,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "defineProperty",
				call: builtinObjectDefineProperty,
			},
		}
		defineProperties := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "defineProperties",
				call: builtinObjectDefineProperties,
			},
		}
		create := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "create",
				call: builtinObjectCreate,
			},
		}
		isExtensible := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isExtensible",
				call: builtinObjectIsExtensible,
			},
		}
		preventExtensions := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "preventExtensions",
				call: builtinObjectPreventExtensions,
			},
		}
		isSealed := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isSealed",
				call: builtinObjectIsSealed,
			},
		}
		seal := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "seal",
				call: builtinObjectSeal,
			},
		}
		isFrozen := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isFrozen",
				call: builtinObjectIsFrozen,
			},
		}
		freeze := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "freeze",
				call: builtinObjectFreeze,
			},
		}
		keys := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "keys",
				call: builtinObjectKeys,
			},
		}
		getOwnPropertyNames := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getOwnPropertyNames",
				call: builtinObjectGetOwnPropertyNames,
			},
		}
		rt.global.Object = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classObjectName,
				call:      builtinObject,
				construct: builtinNewObject,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.ObjectPrototype,
					},
				},
				"getPrototypeOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getPrototypeOf,
					},
				},
				"getOwnPropertyDescriptor": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getOwnPropertyDescriptor,
					},
				},
				"defineProperty": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: defineProperty,
					},
				},
				"defineProperties": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: defineProperties,
					},
				},
				"create": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: create,
					},
				},
				"isExtensible": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: isExtensible,
					},
				},
				"preventExtensions": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: preventExtensions,
					},
				},
				"isSealed": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: isSealed,
					},
				},
				"seal": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: seal,
					},
				},
				"isFrozen": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: isFrozen,
					},
				},
				"freeze": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: freeze,
					},
				},
				"keys": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: keys,
					},
				},
				"getOwnPropertyNames": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getOwnPropertyNames,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
				"getPrototypeOf",
				"getOwnPropertyDescriptor",
				"defineProperty",
				"defineProperties",
				"create",
				"isExtensible",
				"preventExtensions",
				"isSealed",
				"seal",
				"isFrozen",
				"freeze",
				"keys",
				"getOwnPropertyNames",
			},
		}
		rt.global.ObjectPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Object,
			},
		}
	}
	{
		Function := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classFunctionName,
				call:      builtinFunction,
				construct: builtinNewFunction,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.FunctionPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.Function = Function
		rt.global.FunctionPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Function,
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinArrayToString,
			},
		}
		toLocaleString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleString",
				call: builtinArrayToLocaleString,
			},
		}
		concat := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "concat",
				call: builtinArrayConcat,
			},
		}
		join := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "join",
				call: builtinArrayJoin,
			},
		}
		splice := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "splice",
				call: builtinArraySplice,
			},
		}
		shift := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "shift",
				call: builtinArrayShift,
			},
		}
		pop := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "pop",
				call: builtinArrayPop,
			},
		}
		push := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "push",
				call: builtinArrayPush,
			},
		}
		slice := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "slice",
				call: builtinArraySlice,
			},
		}
		unshift := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "unshift",
				call: builtinArrayUnshift,
			},
		}
		reverse := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "reverse",
				call: builtinArrayReverse,
			},
		}
		sort := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "sort",
				call: builtinArraySort,
			},
		}
		indexOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "indexOf",
				call: builtinArrayIndexOf,
			},
		}
		lastIndexOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "lastIndexOf",
				call: builtinArrayLastIndexOf,
			},
		}
		every := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "every",
				call: builtinArrayEvery,
			},
		}
		some := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "some",
				call: builtinArraySome,
			},
		}
		forEach := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "forEach",
				call: builtinArrayForEach,
			},
		}
		mapObj := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "map",
				call: builtinArrayMap,
			},
		}
		filter := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "filter",
				call: builtinArrayFilter,
			},
		}
		reduce := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "reduce",
				call: builtinArrayReduce,
			},
		}
		reduceRight := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "reduceRight",
				call: builtinArrayReduceRight,
			},
		}
		isArray := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isArray",
				call: builtinArrayIsArray,
			},
		}
		rt.global.ArrayPrototype = &object{
			runtime:     rt,
			class:       classArrayName,
			objectClass: classArray,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				propertyLength: {
					mode: 0o100,
					value: Value{
						kind:  valueNumber,
						value: uint32(0),
					},
				},
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"toLocaleString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleString,
					},
				},
				"concat": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: concat,
					},
				},
				"join": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: join,
					},
				},
				"splice": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: splice,
					},
				},
				"shift": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: shift,
					},
				},
				"pop": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: pop,
					},
				},
				"push": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: push,
					},
				},
				"slice": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: slice,
					},
				},
				"unshift": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: unshift,
					},
				},
				"reverse": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: reverse,
					},
				},
				"sort": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: sort,
					},
				},
				"indexOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: indexOf,
					},
				},
				"lastIndexOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: lastIndexOf,
					},
				},
				"every": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: every,
					},
				},
				"some": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: some,
					},
				},
				"forEach": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: forEach,
					},
				},
				"map": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: mapObj,
					},
				},
				"filter": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: filter,
					},
				},
				"reduce": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: reduce,
					},
				},
				"reduceRight": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: reduceRight,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"toString",
				"toLocaleString",
				"concat",
				"join",
				"splice",
				"shift",
				"pop",
				"push",
				"slice",
				"unshift",
				"reverse",
				"sort",
				"indexOf",
				"lastIndexOf",
				"every",
				"some",
				"forEach",
				"map",
				"filter",
				"reduce",
				"reduceRight",
			},
		}
		rt.global.Array = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classArrayName,
				call:      builtinArray,
				construct: builtinNewArray,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.ArrayPrototype,
					},
				},
				"isArray": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: isArray,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
				"isArray",
			},
		}
		rt.global.ArrayPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Array,
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinStringToString,
			},
		}
		valueOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "valueOf",
				call: builtinStringValueOf,
			},
		}
		charAt := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "charAt",
				call: builtinStringCharAt,
			},
		}
		charCodeAt := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "charCodeAt",
				call: builtinStringCharCodeAt,
			},
		}
		concat := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "concat",
				call: builtinStringConcat,
			},
		}
		indexOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "indexOf",
				call: builtinStringIndexOf,
			},
		}
		lastIndexOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "lastIndexOf",
				call: builtinStringLastIndexOf,
			},
		}
		match := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "match",
				call: builtinStringMatch,
			},
		}
		replace := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "replace",
				call: builtinStringReplace,
			},
		}
		search := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "search",
				call: builtinStringSearch,
			},
		}
		split := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "split",
				call: builtinStringSplit,
			},
		}
		slice := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "slice",
				call: builtinStringSlice,
			},
		}
		substring := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "substring",
				call: builtinStringSubstring,
			},
		}
		toLowerCase := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLowerCase",
				call: builtinStringToLowerCase,
			},
		}
		toUpperCase := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toUpperCase",
				call: builtinStringToUpperCase,
			},
		}
		substr := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "substr",
				call: builtinStringSubstr,
			},
		}
		trim := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "trim",
				call: builtinStringTrim,
			},
		}
		trimLeft := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "trimLeft",
				call: builtinStringTrimLeft,
			},
		}
		trimRight := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "trimRight",
				call: builtinStringTrimRight,
			},
		}
		localeCompare := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "localeCompare",
				call: builtinStringLocaleCompare,
			},
		}
		toLocaleLowerCase := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleLowerCase",
				call: builtinStringToLocaleLowerCase,
			},
		}
		toLocaleUpperCase := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleUpperCase",
				call: builtinStringToLocaleUpperCase,
			},
		}
		fromCharCode := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "fromCharCode",
				call: builtinStringFromCharCode,
			},
		}
		rt.global.StringPrototype = &object{
			runtime:     rt,
			class:       classStringName,
			objectClass: classString,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       prototypeValueString,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: int(0),
					},
				},
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"valueOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: valueOf,
					},
				},
				"charAt": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: charAt,
					},
				},
				"charCodeAt": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: charCodeAt,
					},
				},
				"concat": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: concat,
					},
				},
				"indexOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: indexOf,
					},
				},
				"lastIndexOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: lastIndexOf,
					},
				},
				"match": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: match,
					},
				},
				"replace": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: replace,
					},
				},
				"search": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: search,
					},
				},
				"split": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: split,
					},
				},
				"slice": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: slice,
					},
				},
				"substring": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: substring,
					},
				},
				"toLowerCase": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLowerCase,
					},
				},
				"toUpperCase": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toUpperCase,
					},
				},
				"substr": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: substr,
					},
				},
				"trim": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: trim,
					},
				},
				"trimLeft": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: trimLeft,
					},
				},
				"trimRight": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: trimRight,
					},
				},
				"localeCompare": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: localeCompare,
					},
				},
				"toLocaleLowerCase": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleLowerCase,
					},
				},
				"toLocaleUpperCase": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleUpperCase,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"toString",
				"valueOf",
				"charAt",
				"charCodeAt",
				"concat",
				"indexOf",
				"lastIndexOf",
				"match",
				"replace",
				"search",
				"split",
				"slice",
				"substring",
				"toLowerCase",
				"toUpperCase",
				"substr",
				"trim",
				"trimLeft",
				"trimRight",
				"localeCompare",
				"toLocaleLowerCase",
				"toLocaleUpperCase",
			},
		}
		rt.global.String = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classStringName,
				call:      builtinString,
				construct: builtinNewString,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.StringPrototype,
					},
				},
				"fromCharCode": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: fromCharCode,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
				"fromCharCode",
			},
		}
		rt.global.StringPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.String,
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinBooleanToString,
			},
		}
		valueOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "valueOf",
				call: builtinBooleanValueOf,
			},
		}
		rt.global.BooleanPrototype = &object{
			runtime:     rt,
			class:       classBooleanName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       prototypeValueBoolean,
			property: map[string]property{
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"valueOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: valueOf,
					},
				},
			},
			propertyOrder: []string{
				"toString",
				"valueOf",
			},
		}
		rt.global.Boolean = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classBooleanName,
				call:      builtinBoolean,
				construct: builtinNewBoolean,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.BooleanPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.BooleanPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Boolean,
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinNumberToString,
			},
		}
		valueOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "valueOf",
				call: builtinNumberValueOf,
			},
		}
		toFixed := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toFixed",
				call: builtinNumberToFixed,
			},
		}
		toExponential := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toExponential",
				call: builtinNumberToExponential,
			},
		}
		toPrecision := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toPrecision",
				call: builtinNumberToPrecision,
			},
		}
		toLocaleString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleString",
				call: builtinNumberToLocaleString,
			},
		}
		isNaN := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isNaN",
				call: builtinNumberIsNaN,
			},
		}
		rt.global.NumberPrototype = &object{
			runtime:     rt,
			class:       classNumberName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       prototypeValueNumber,
			property: map[string]property{
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"valueOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: valueOf,
					},
				},
				"toFixed": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toFixed,
					},
				},
				"toExponential": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toExponential,
					},
				},
				"toPrecision": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toPrecision,
					},
				},
				"toLocaleString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleString,
					},
				},
			},
			propertyOrder: []string{
				"toString",
				"valueOf",
				"toFixed",
				"toExponential",
				"toPrecision",
				"toLocaleString",
			},
		}
		rt.global.Number = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classNumberName,
				call:      builtinNumber,
				construct: builtinNewNumber,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.NumberPrototype,
					},
				},
				"isNaN": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: isNaN,
					},
				},
				"MAX_VALUE": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.MaxFloat64,
					},
				},
				"MIN_VALUE": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.SmallestNonzeroFloat64,
					},
				},
				"NaN": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.NaN(),
					},
				},
				"NEGATIVE_INFINITY": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Inf(-1),
					},
				},
				"POSITIVE_INFINITY": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Inf(+1),
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
				"isNaN",
				"MAX_VALUE",
				"MIN_VALUE",
				"NaN",
				"NEGATIVE_INFINITY",
				"POSITIVE_INFINITY",
			},
		}
		rt.global.NumberPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Number,
			},
		}
	}
	{
		abs := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "abs",
				call: builtinMathAbs,
			},
		}
		acos := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "acos",
				call: builtinMathAcos,
			},
		}
		asin := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "asin",
				call: builtinMathAsin,
			},
		}
		atan := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "atan",
				call: builtinMathAtan,
			},
		}
		atan2 := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "atan2",
				call: builtinMathAtan2,
			},
		}
		ceil := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "ceil",
				call: builtinMathCeil,
			},
		}
		cos := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "cos",
				call: builtinMathCos,
			},
		}
		exp := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "exp",
				call: builtinMathExp,
			},
		}
		floor := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "floor",
				call: builtinMathFloor,
			},
		}
		log := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "log",
				call: builtinMathLog,
			},
		}
		max := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "max",
				call: builtinMathMax,
			},
		}
		min := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "min",
				call: builtinMathMin,
			},
		}
		pow := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "pow",
				call: builtinMathPow,
			},
		}
		random := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "random",
				call: builtinMathRandom,
			},
		}
		round := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "round",
				call: builtinMathRound,
			},
		}
		sin := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "sin",
				call: builtinMathSin,
			},
		}
		sqrt := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "sqrt",
				call: builtinMathSqrt,
			},
		}
		tan := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "tan",
				call: builtinMathTan,
			},
		}
		rt.global.Math = &object{
			runtime:     rt,
			class:       "Math",
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			property: map[string]property{
				"abs": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: abs,
					},
				},
				"acos": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: acos,
					},
				},
				"asin": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: asin,
					},
				},
				"atan": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: atan,
					},
				},
				"atan2": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: atan2,
					},
				},
				"ceil": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: ceil,
					},
				},
				"cos": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: cos,
					},
				},
				"exp": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: exp,
					},
				},
				"floor": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: floor,
					},
				},
				"log": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: log,
					},
				},
				"max": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: max,
					},
				},
				"min": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: min,
					},
				},
				"pow": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: pow,
					},
				},
				"random": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: random,
					},
				},
				"round": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: round,
					},
				},
				"sin": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: sin,
					},
				},
				"sqrt": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: sqrt,
					},
				},
				"tan": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: tan,
					},
				},
				"E": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.E,
					},
				},
				"LN10": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Ln10,
					},
				},
				"LN2": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Ln2,
					},
				},
				"LOG2E": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Log2E,
					},
				},
				"LOG10E": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Log10E,
					},
				},
				"PI": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Pi,
					},
				},
				"SQRT1_2": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: sqrt1_2,
					},
				},
				"SQRT2": {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: math.Sqrt2,
					},
				},
			},
			propertyOrder: []string{
				"abs",
				"acos",
				"asin",
				"atan",
				"atan2",
				"ceil",
				"cos",
				"exp",
				"floor",
				"log",
				"max",
				"min",
				"pow",
				"random",
				"round",
				"sin",
				"sqrt",
				"tan",
				"E",
				"LN10",
				"LN2",
				"LOG2E",
				"LOG10E",
				"PI",
				"SQRT1_2",
				"SQRT2",
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinDateToString,
			},
		}
		toDateString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toDateString",
				call: builtinDateToDateString,
			},
		}
		toTimeString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toTimeString",
				call: builtinDateToTimeString,
			},
		}
		toUTCString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toUTCString",
				call: builtinDateToUTCString,
			},
		}
		toISOString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toISOString",
				call: builtinDateToISOString,
			},
		}
		toJSON := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toJSON",
				call: builtinDateToJSON,
			},
		}
		toGMTString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toGMTString",
				call: builtinDateToGMTString,
			},
		}
		toLocaleString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleString",
				call: builtinDateToLocaleString,
			},
		}
		toLocaleDateString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleDateString",
				call: builtinDateToLocaleDateString,
			},
		}
		toLocaleTimeString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toLocaleTimeString",
				call: builtinDateToLocaleTimeString,
			},
		}
		valueOf := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "valueOf",
				call: builtinDateValueOf,
			},
		}
		getTime := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getTime",
				call: builtinDateGetTime,
			},
		}
		getYear := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getYear",
				call: builtinDateGetYear,
			},
		}
		getFullYear := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getFullYear",
				call: builtinDateGetFullYear,
			},
		}
		getUTCFullYear := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCFullYear",
				call: builtinDateGetUTCFullYear,
			},
		}
		getMonth := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getMonth",
				call: builtinDateGetMonth,
			},
		}
		getUTCMonth := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCMonth",
				call: builtinDateGetUTCMonth,
			},
		}
		getDate := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getDate",
				call: builtinDateGetDate,
			},
		}
		getUTCDate := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCDate",
				call: builtinDateGetUTCDate,
			},
		}
		getDay := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getDay",
				call: builtinDateGetDay,
			},
		}
		getUTCDay := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCDay",
				call: builtinDateGetUTCDay,
			},
		}
		getHours := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getHours",
				call: builtinDateGetHours,
			},
		}
		getUTCHours := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCHours",
				call: builtinDateGetUTCHours,
			},
		}
		getMinutes := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getMinutes",
				call: builtinDateGetMinutes,
			},
		}
		getUTCMinutes := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCMinutes",
				call: builtinDateGetUTCMinutes,
			},
		}
		getSeconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getSeconds",
				call: builtinDateGetSeconds,
			},
		}
		getUTCSeconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCSeconds",
				call: builtinDateGetUTCSeconds,
			},
		}
		getMilliseconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getMilliseconds",
				call: builtinDateGetMilliseconds,
			},
		}
		getUTCMilliseconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getUTCMilliseconds",
				call: builtinDateGetUTCMilliseconds,
			},
		}
		getTimezoneOffset := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "getTimezoneOffset",
				call: builtinDateGetTimezoneOffset,
			},
		}
		setTime := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setTime",
				call: builtinDateSetTime,
			},
		}
		setMilliseconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setMilliseconds",
				call: builtinDateSetMilliseconds,
			},
		}
		setUTCMilliseconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCMilliseconds",
				call: builtinDateSetUTCMilliseconds,
			},
		}
		setSeconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setSeconds",
				call: builtinDateSetSeconds,
			},
		}
		setUTCSeconds := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCSeconds",
				call: builtinDateSetUTCSeconds,
			},
		}
		setMinutes := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 3,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setMinutes",
				call: builtinDateSetMinutes,
			},
		}
		setUTCMinutes := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 3,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCMinutes",
				call: builtinDateSetUTCMinutes,
			},
		}
		setHours := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 4,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setHours",
				call: builtinDateSetHours,
			},
		}
		setUTCHours := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 4,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCHours",
				call: builtinDateSetUTCHours,
			},
		}
		setDate := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setDate",
				call: builtinDateSetDate,
			},
		}
		setUTCDate := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCDate",
				call: builtinDateSetUTCDate,
			},
		}
		setMonth := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setMonth",
				call: builtinDateSetMonth,
			},
		}
		setUTCMonth := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCMonth",
				call: builtinDateSetUTCMonth,
			},
		}
		setYear := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setYear",
				call: builtinDateSetYear,
			},
		}
		setFullYear := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 3,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setFullYear",
				call: builtinDateSetFullYear,
			},
		}
		setUTCFullYear := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 3,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "setUTCFullYear",
				call: builtinDateSetUTCFullYear,
			},
		}
		parse := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "parse",
				call: builtinDateParse,
			},
		}
		UTC := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 7,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "UTC",
				call: builtinDateUTC,
			},
		}
		now := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "now",
				call: builtinDateNow,
			},
		}
		rt.global.DatePrototype = &object{
			runtime:     rt,
			class:       classDateName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       prototypeValueDate,
			property: map[string]property{
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"toDateString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toDateString,
					},
				},
				"toTimeString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toTimeString,
					},
				},
				"toUTCString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toUTCString,
					},
				},
				"toISOString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toISOString,
					},
				},
				"toJSON": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toJSON,
					},
				},
				"toGMTString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toGMTString,
					},
				},
				"toLocaleString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleString,
					},
				},
				"toLocaleDateString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleDateString,
					},
				},
				"toLocaleTimeString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toLocaleTimeString,
					},
				},
				"valueOf": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: valueOf,
					},
				},
				"getTime": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getTime,
					},
				},
				"getYear": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getYear,
					},
				},
				"getFullYear": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getFullYear,
					},
				},
				"getUTCFullYear": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCFullYear,
					},
				},
				"getMonth": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getMonth,
					},
				},
				"getUTCMonth": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCMonth,
					},
				},
				"getDate": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getDate,
					},
				},
				"getUTCDate": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCDate,
					},
				},
				"getDay": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getDay,
					},
				},
				"getUTCDay": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCDay,
					},
				},
				"getHours": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getHours,
					},
				},
				"getUTCHours": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCHours,
					},
				},
				"getMinutes": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getMinutes,
					},
				},
				"getUTCMinutes": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCMinutes,
					},
				},
				"getSeconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getSeconds,
					},
				},
				"getUTCSeconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCSeconds,
					},
				},
				"getMilliseconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getMilliseconds,
					},
				},
				"getUTCMilliseconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getUTCMilliseconds,
					},
				},
				"getTimezoneOffset": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: getTimezoneOffset,
					},
				},
				"setTime": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setTime,
					},
				},
				"setMilliseconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setMilliseconds,
					},
				},
				"setUTCMilliseconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCMilliseconds,
					},
				},
				"setSeconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setSeconds,
					},
				},
				"setUTCSeconds": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCSeconds,
					},
				},
				"setMinutes": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setMinutes,
					},
				},
				"setUTCMinutes": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCMinutes,
					},
				},
				"setHours": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setHours,
					},
				},
				"setUTCHours": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCHours,
					},
				},
				"setDate": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setDate,
					},
				},
				"setUTCDate": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCDate,
					},
				},
				"setMonth": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setMonth,
					},
				},
				"setUTCMonth": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCMonth,
					},
				},
				"setYear": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setYear,
					},
				},
				"setFullYear": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setFullYear,
					},
				},
				"setUTCFullYear": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: setUTCFullYear,
					},
				},
			},
			propertyOrder: []string{
				"toString",
				"toDateString",
				"toTimeString",
				"toUTCString",
				"toISOString",
				"toJSON",
				"toGMTString",
				"toLocaleString",
				"toLocaleDateString",
				"toLocaleTimeString",
				"valueOf",
				"getTime",
				"getYear",
				"getFullYear",
				"getUTCFullYear",
				"getMonth",
				"getUTCMonth",
				"getDate",
				"getUTCDate",
				"getDay",
				"getUTCDay",
				"getHours",
				"getUTCHours",
				"getMinutes",
				"getUTCMinutes",
				"getSeconds",
				"getUTCSeconds",
				"getMilliseconds",
				"getUTCMilliseconds",
				"getTimezoneOffset",
				"setTime",
				"setMilliseconds",
				"setUTCMilliseconds",
				"setSeconds",
				"setUTCSeconds",
				"setMinutes",
				"setUTCMinutes",
				"setHours",
				"setUTCHours",
				"setDate",
				"setUTCDate",
				"setMonth",
				"setUTCMonth",
				"setYear",
				"setFullYear",
				"setUTCFullYear",
			},
		}
		rt.global.Date = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classDateName,
				call:      builtinDate,
				construct: builtinNewDate,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 7,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.DatePrototype,
					},
				},
				"parse": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: parse,
					},
				},
				"UTC": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: UTC,
					},
				},
				"now": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: now,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
				"parse",
				"UTC",
				"now",
			},
		}
		rt.global.DatePrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Date,
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinRegExpToString,
			},
		}
		exec := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "exec",
				call: builtinRegExpExec,
			},
		}
		test := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "test",
				call: builtinRegExpTest,
			},
		}
		compile := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "compile",
				call: builtinRegExpCompile,
			},
		}
		rt.global.RegExpPrototype = &object{
			runtime:     rt,
			class:       classRegExpName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       prototypeValueRegExp,
			property: map[string]property{
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"exec": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: exec,
					},
				},
				"test": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: test,
					},
				},
				"compile": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: compile,
					},
				},
			},
			propertyOrder: []string{
				"toString",
				"exec",
				"test",
				"compile",
			},
		}
		rt.global.RegExp = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classRegExpName,
				call:      builtinRegExp,
				construct: builtinNewRegExp,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.RegExpPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.RegExpPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.RegExp,
			},
		}
	}
	{
		toString := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "toString",
				call: builtinErrorToString,
			},
		}
		rt.global.ErrorPrototype = &object{
			runtime:     rt,
			class:       classErrorName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"toString": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: toString,
					},
				},
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: classErrorName,
					},
				},
				"message": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "",
					},
				},
			},
			propertyOrder: []string{
				"toString",
				"name",
				"message",
			},
		}
		rt.global.Error = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      classErrorName,
				call:      builtinError,
				construct: builtinNewError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.ErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.ErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.Error,
			},
		}
	}
	{
		rt.global.EvalErrorPrototype = &object{
			runtime:     rt,
			class:       "EvalError",
			objectClass: classObject,
			prototype:   rt.global.ErrorPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "EvalError",
					},
				},
			},
			propertyOrder: []string{
				"name",
			},
		}
		rt.global.EvalError = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      "EvalError",
				call:      builtinEvalError,
				construct: builtinNewEvalError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.EvalErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.EvalErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.EvalError,
			},
		}
	}
	{
		rt.global.TypeErrorPrototype = &object{
			runtime:     rt,
			class:       "TypeError",
			objectClass: classObject,
			prototype:   rt.global.ErrorPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "TypeError",
					},
				},
			},
			propertyOrder: []string{
				"name",
			},
		}
		rt.global.TypeError = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      "TypeError",
				call:      builtinTypeError,
				construct: builtinNewTypeError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.TypeErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.TypeErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.TypeError,
			},
		}
	}
	{
		rt.global.RangeErrorPrototype = &object{
			runtime:     rt,
			class:       "RangeError",
			objectClass: classObject,
			prototype:   rt.global.ErrorPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "RangeError",
					},
				},
			},
			propertyOrder: []string{
				"name",
			},
		}
		rt.global.RangeError = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      "RangeError",
				call:      builtinRangeError,
				construct: builtinNewRangeError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.RangeErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.RangeErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.RangeError,
			},
		}
	}
	{
		rt.global.ReferenceErrorPrototype = &object{
			runtime:     rt,
			class:       "ReferenceError",
			objectClass: classObject,
			prototype:   rt.global.ErrorPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "ReferenceError",
					},
				},
			},
			propertyOrder: []string{
				"name",
			},
		}
		rt.global.ReferenceError = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      "ReferenceError",
				call:      builtinReferenceError,
				construct: builtinNewReferenceError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.ReferenceErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.ReferenceErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.ReferenceError,
			},
		}
	}
	{
		rt.global.SyntaxErrorPrototype = &object{
			runtime:     rt,
			class:       "SyntaxError",
			objectClass: classObject,
			prototype:   rt.global.ErrorPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "SyntaxError",
					},
				},
			},
			propertyOrder: []string{
				"name",
			},
		}
		rt.global.SyntaxError = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      "SyntaxError",
				call:      builtinSyntaxError,
				construct: builtinNewSyntaxError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.SyntaxErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.SyntaxErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.SyntaxError,
			},
		}
	}
	{
		rt.global.URIErrorPrototype = &object{
			runtime:     rt,
			class:       "URIError",
			objectClass: classObject,
			prototype:   rt.global.ErrorPrototype,
			extensible:  true,
			value:       nil,
			property: map[string]property{
				"name": {
					mode: 0o101,
					value: Value{
						kind:  valueString,
						value: "URIError",
					},
				},
			},
			propertyOrder: []string{
				"name",
			},
		}
		rt.global.URIError = &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			value: nativeFunctionObject{
				name:      "URIError",
				call:      builtinURIError,
				construct: builtinNewURIError,
			},
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
				"prototype": {
					mode: 0,
					value: Value{
						kind:  valueObject,
						value: rt.global.URIErrorPrototype,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
				"prototype",
			},
		}
		rt.global.URIErrorPrototype.property["constructor"] = property{
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.URIError,
			},
		}
	}
	{
		parse := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "parse",
				call: builtinJSONParse,
			},
		}
		stringify := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 3,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "stringify",
				call: builtinJSONStringify,
			},
		}
		rt.global.JSON = &object{
			runtime:     rt,
			class:       "JSON",
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			property: map[string]property{
				"parse": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: parse,
					},
				},
				"stringify": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: stringify,
					},
				},
			},
			propertyOrder: []string{
				"parse",
				"stringify",
			},
		}
	}
	{
		eval := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "eval",
				call: builtinGlobalEval,
			},
		}
		parseInt := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 2,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "parseInt",
				call: builtinGlobalParseInt,
			},
		}
		parseFloat := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "parseFloat",
				call: builtinGlobalParseFloat,
			},
		}
		isNaN := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isNaN",
				call: builtinGlobalIsNaN,
			},
		}
		isFinite := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "isFinite",
				call: builtinGlobalIsFinite,
			},
		}
		decodeURI := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "decodeURI",
				call: builtinGlobalDecodeURI,
			},
		}
		decodeURIComponent := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "decodeURIComponent",
				call: builtinGlobalDecodeURIComponent,
			},
		}
		encodeURI := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "encodeURI",
				call: builtinGlobalEncodeURI,
			},
		}
		encodeURIComponent := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "encodeURIComponent",
				call: builtinGlobalEncodeURIComponent,
			},
		}
		escape := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "escape",
				call: builtinGlobalEscape,
			},
		}
		unescape := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 1,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "unescape",
				call: builtinGlobalUnescape,
			},
		}
		rt.globalObject.property = map[string]property{
			"eval": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: eval,
				},
			},
			"parseInt": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: parseInt,
				},
			},
			"parseFloat": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: parseFloat,
				},
			},
			"isNaN": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: isNaN,
				},
			},
			"isFinite": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: isFinite,
				},
			},
			"decodeURI": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: decodeURI,
				},
			},
			"decodeURIComponent": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: decodeURIComponent,
				},
			},
			"encodeURI": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: encodeURI,
				},
			},
			"encodeURIComponent": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: encodeURIComponent,
				},
			},
			"escape": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: escape,
				},
			},
			"unescape": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: unescape,
				},
			},
			classObjectName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Object,
				},
			},
			classFunctionName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Function,
				},
			},
			classArrayName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Array,
				},
			},
			classStringName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.String,
				},
			},
			classBooleanName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Boolean,
				},
			},
			classNumberName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Number,
				},
			},
			"Math": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Math,
				},
			},
			classDateName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Date,
				},
			},
			classRegExpName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.RegExp,
				},
			},
			classErrorName: {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.Error,
				},
			},
			"EvalError": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.EvalError,
				},
			},
			"TypeError": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.TypeError,
				},
			},
			"RangeError": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.RangeError,
				},
			},
			"ReferenceError": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.ReferenceError,
				},
			},
			"SyntaxError": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.SyntaxError,
				},
			},
			"URIError": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.URIError,
				},
			},
			"JSON": {
				mode: 0o101,
				value: Value{
					kind:  valueObject,
					value: rt.global.JSON,
				},
			},
			"undefined": {
				mode: 0,
				value: Value{
					kind: valueUndefined,
				},
			},
			"NaN": {
				mode: 0,
				value: Value{
					kind:  valueNumber,
					value: math.NaN(),
				},
			},
			"Infinity": {
				mode: 0,
				value: Value{
					kind:  valueNumber,
					value: math.Inf(+1),
				},
			},
		}
		rt.globalObject.propertyOrder = []string{
			"eval",
			"parseInt",
			"parseFloat",
			"isNaN",
			"isFinite",
			"decodeURI",
			"decodeURIComponent",
			"encodeURI",
			"encodeURIComponent",
			"escape",
			"unescape",
			classObjectName,
			classFunctionName,
			classArrayName,
			classStringName,
			classBooleanName,
			classNumberName,
			"Math",
			classDateName,
			classRegExpName,
			classErrorName,
			"EvalError",
			"TypeError",
			"RangeError",
			"ReferenceError",
			"SyntaxError",
			"URIError",
			"JSON",
			"undefined",
			"NaN",
			"Infinity",
		}
	}
}

func newConsoleObject(rt *runtime) *object {
	{
		log := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "log",
				call: builtinConsoleLog,
			},
		}
		debug := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "debug",
				call: builtinConsoleLog,
			},
		}
		info := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "info",
				call: builtinConsoleLog,
			},
		}
		errorObj := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "error",
				call: builtinConsoleError,
			},
		}
		warn := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "warn",
				call: builtinConsoleError,
			},
		}
		dir := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "dir",
				call: builtinConsoleDir,
			},
		}
		time := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "time",
				call: builtinConsoleTime,
			},
		}
		timeEnd := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "timeEnd",
				call: builtinConsoleTimeEnd,
			},
		}
		trace := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "trace",
				call: builtinConsoleTrace,
			},
		}
		assert := &object{
			runtime:     rt,
			class:       classFunctionName,
			objectClass: classObject,
			prototype:   rt.global.FunctionPrototype,
			extensible:  true,
			property: map[string]property{
				propertyLength: {
					mode: 0,
					value: Value{
						kind:  valueNumber,
						value: 0,
					},
				},
			},
			propertyOrder: []string{
				propertyLength,
			},
			value: nativeFunctionObject{
				name: "assert",
				call: builtinConsoleAssert,
			},
		}
		return &object{
			runtime:     rt,
			class:       classObjectName,
			objectClass: classObject,
			prototype:   rt.global.ObjectPrototype,
			extensible:  true,
			property: map[string]property{
				"log": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: log,
					},
				},
				"debug": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: debug,
					},
				},
				"info": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: info,
					},
				},
				"error": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: errorObj,
					},
				},
				"warn": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: warn,
					},
				},
				"dir": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: dir,
					},
				},
				"time": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: time,
					},
				},
				"timeEnd": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: timeEnd,
					},
				},
				"trace": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: trace,
					},
				},
				"assert": {
					mode: 0o101,
					value: Value{
						kind:  valueObject,
						value: assert,
					},
				},
			},
			propertyOrder: []string{
				"log",
				"debug",
				"info",
				"error",
				"warn",
				"dir",
				"time",
				"timeEnd",
				"trace",
				"assert",
			},
		}
	}
}

func intValue(value int) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func int32Value(value int32) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func int64Value(value int64) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func uint16Value(value uint16) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func uint32Value(value uint32) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func float64Value(value float64) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func stringValue(value string) Value {
	return Value{
		kind:  valueString,
		value: value,
	}
}

func string16Value(value []uint16) Value {
	return Value{
		kind:  valueString,
		value: value,
	}
}

func boolValue(value bool) Value {
	return Value{
		kind:  valueBoolean,
		value: value,
	}
}

func objectValue(value *object) Value {
	return Value{
		kind:  valueObject,
		value: value,
	}
}
