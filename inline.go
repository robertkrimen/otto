package otto

import (
	"math"
)

func (rt *runtime) newContext() {
	// Order here is import as definitions depend on each other.

	// Object prototype.
	rt.global.ObjectPrototype = &object{
		runtime:     rt,
		class:       classObjectName,
		objectClass: classObject,
		prototype:   nil,
		extensible:  true,
		value:       prototypeValueObject,
	}

	// Function prototype.
	rt.global.FunctionPrototype = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		value:       prototypeValueFunction,
	}

	// Object prototype property definition.
	rt.global.ObjectPrototype.property = map[string]property{
		"hasOwnProperty": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "hasOwnProperty",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "hasOwnProperty",
						call: builtinObjectHasOwnProperty,
					},
				},
			},
		},
		"isPrototypeOf": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "isPrototypeOf",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "isPrototypeOf",
						call: builtinObjectIsPrototypeOf,
					},
				},
			},
		},
		"propertyIsEnumerable": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "propertyIsEnumerable",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "propertyIsEnumerable",
						call: builtinObjectPropertyIsEnumerable,
					},
				},
			},
		},
		methodToString: {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "toString",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: methodToString,
						call: builtinObjectToString,
					},
				},
			},
		},
		"valueOf": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "valueOf",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "valueOf",
						call: builtinObjectValueOf,
					},
				},
			},
		},
		"toLocaleString": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "toLocaleString",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "toLocaleString",
						call: builtinObjectToLocaleString,
					},
				},
			},
		},
	}
	rt.global.ObjectPrototype.propertyOrder = []string{
		propertyConstructor,
		"hasOwnProperty",
		"isPrototypeOf",
		"propertyIsEnumerable",
		methodToString,
		"valueOf",
		"toLocaleString",
	}

	// Function prototype property definition.
	rt.global.FunctionPrototype.property = map[string]property{
		methodToString: {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "toString",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: methodToString,
						call: builtinFunctionToString,
					},
				},
			},
		},
		"apply": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "apply",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "apply",
						call: builtinFunctionApply,
					},
				},
			},
		},
		"call": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "call",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "call",
						call: builtinFunctionCall,
					},
				},
			},
		},
		"bind": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "bind",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "bind",
						call: builtinFunctionBind,
					},
				},
			},
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
		methodToString,
		"apply",
		"call",
		"bind",
		propertyConstructor,
		propertyLength,
	}

	// Object definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.ObjectPrototype,
				},
			},
			"getPrototypeOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getPrototypeOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getPrototypeOf",
							call: builtinObjectGetPrototypeOf,
						},
					},
				},
			},
			"getOwnPropertyDescriptor": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getOwnPropertyDescriptor",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getOwnPropertyDescriptor",
							call: builtinObjectGetOwnPropertyDescriptor,
						},
					},
				},
			},
			"defineProperty": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "defineProperty",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "defineProperty",
							call: builtinObjectDefineProperty,
						},
					},
				},
			},
			"defineProperties": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "defineProperties",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "defineProperties",
							call: builtinObjectDefineProperties,
						},
					},
				},
			},
			"create": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "create",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "create",
							call: builtinObjectCreate,
						},
					},
				},
			},
			"isExtensible": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "isExtensible",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "isExtensible",
							call: builtinObjectIsExtensible,
						},
					},
				},
			},
			"preventExtensions": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "preventExtensions",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "preventExtensions",
							call: builtinObjectPreventExtensions,
						},
					},
				},
			},
			"isSealed": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "isSealed",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "isSealed",
							call: builtinObjectIsSealed,
						},
					},
				},
			},
			"seal": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "seal",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "seal",
							call: builtinObjectSeal,
						},
					},
				},
			},
			"isFrozen": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "isFrozen",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "isFrozen",
							call: builtinObjectIsFrozen,
						},
					},
				},
			},
			"freeze": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "freeze",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "freeze",
							call: builtinObjectFreeze,
						},
					},
				},
			},
			"keys": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "keys",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "keys",
							call: builtinObjectKeys,
						},
					},
				},
			},
			"getOwnPropertyNames": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getOwnPropertyNames",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getOwnPropertyNames",
							call: builtinObjectGetOwnPropertyNames,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
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

	// Object constructor definition.
	rt.global.ObjectPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Object,
		},
	}

	// Function definition.
	rt.global.Function = &object{
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.FunctionPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// Function constructor definition.
	rt.global.FunctionPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Function,
		},
	}

	// Array prototype.
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
			"concat": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "concat",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "concat",
							call: builtinArrayConcat,
						},
					},
				},
			},
			"lastIndexOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "lastIndexOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "lastIndexOf",
							call: builtinArrayLastIndexOf,
						},
					},
				},
			},
			"pop": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "pop",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "pop",
							call: builtinArrayPop,
						},
					},
				},
			},
			"push": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "push",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "push",
							call: builtinArrayPush,
						},
					},
				},
			},
			"reverse": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "reverse",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "reverse",
							call: builtinArrayReverse,
						},
					},
				},
			},
			"shift": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "shift",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "shift",
							call: builtinArrayShift,
						},
					},
				},
			},
			"unshift": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "unshift",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "unshift",
							call: builtinArrayUnshift,
						},
					},
				},
			},
			"slice": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "slice",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "slice",
							call: builtinArraySlice,
						},
					},
				},
			},
			"sort": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "sort",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "sort",
							call: builtinArraySort,
						},
					},
				},
			},
			"splice": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "splice",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "splice",
							call: builtinArraySplice,
						},
					},
				},
			},
			"indexOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "indexOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "indexOf",
							call: builtinArrayIndexOf,
						},
					},
				},
			},
			"join": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "join",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "join",
							call: builtinArrayJoin,
						},
					},
				},
			},
			"forEach": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "forEach",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "forEach",
							call: builtinArrayForEach,
						},
					},
				},
			},
			"filter": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "filter",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "filter",
							call: builtinArrayFilter,
						},
					},
				},
			},
			"map": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "map",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "map",
							call: builtinArrayMap,
						},
					},
				},
			},
			"every": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "every",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "every",
							call: builtinArrayEvery,
						},
					},
				},
			},
			"some": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "some",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "some",
							call: builtinArraySome,
						},
					},
				},
			},
			"reduce": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "reduce",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "reduce",
							call: builtinArrayReduce,
						},
					},
				},
			},
			"reduceRight": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "reduceRight",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "reduceRight",
							call: builtinArrayReduceRight,
						},
					},
				},
			},
			"toLocaleString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleString",
							call: builtinArrayToLocaleString,
						},
					},
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinArrayToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyConstructor,
			"concat",
			"lastIndexOf",
			"pop",
			"push",
			"reverse",
			"shift",
			"unshift",
			"slice",
			"sort",
			"splice",
			"indexOf",
			"join",
			"forEach",
			"filter",
			"map",
			"every",
			"some",
			"reduce",
			"reduceRight",
			"toLocaleString",
			methodToString,
		},
	}

	// Array definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.ArrayPrototype,
				},
			},
			"isArray": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "isArray",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "isArray",
							call: builtinArrayIsArray,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
			"isArray",
		},
	}

	// Array constructor definition.
	rt.global.ArrayPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Array,
		},
	}

	// String prototype.
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
			"charAt": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "charAt",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "charAt",
							call: builtinStringCharAt,
						},
					},
				},
			},
			"charCodeAt": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "charCodeAt",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "charCodeAt",
							call: builtinStringCharCodeAt,
						},
					},
				},
			},
			"concat": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "concat",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "concat",
							call: builtinStringConcat,
						},
					},
				},
			},
			"indexOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "indexOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "indexOf",
							call: builtinStringIndexOf,
						},
					},
				},
			},
			"lastIndexOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "lastIndexOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "lastIndexOf",
							call: builtinStringLastIndexOf,
						},
					},
				},
			},
			"localeCompare": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "localeCompare",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "localeCompare",
							call: builtinStringLocaleCompare,
						},
					},
				},
			},
			"match": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "match",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "match",
							call: builtinStringMatch,
						},
					},
				},
			},
			"replace": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "replace",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "replace",
							call: builtinStringReplace,
						},
					},
				},
			},
			"search": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "search",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "search",
							call: builtinStringSearch,
						},
					},
				},
			},
			"slice": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "slice",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "slice",
							call: builtinStringSlice,
						},
					},
				},
			},
			"split": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "split",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "split",
							call: builtinStringSplit,
						},
					},
				},
			},
			"substr": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "substr",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "substr",
							call: builtinStringSubstr,
						},
					},
				},
			},
			"substring": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "substring",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "substring",
							call: builtinStringSubstring,
						},
					},
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinStringToString,
						},
					},
				},
			},
			"trim": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "trim",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "trim",
							call: builtinStringTrim,
						},
					},
				},
			},
			"trimLeft": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "trimLeft",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "trimLeft",
							call: builtinStringTrimLeft,
						},
					},
				},
			},
			"trimRight": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "trimRight",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "trimRight",
							call: builtinStringTrimRight,
						},
					},
				},
			},
			"toLocaleLowerCase": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleLowerCase",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleLowerCase",
							call: builtinStringToLocaleLowerCase,
						},
					},
				},
			},
			"toLocaleUpperCase": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleUpperCase",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleUpperCase",
							call: builtinStringToLocaleUpperCase,
						},
					},
				},
			},
			"toLowerCase": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLowerCase",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLowerCase",
							call: builtinStringToLowerCase,
						},
					},
				},
			},
			"toUpperCase": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toUpperCase",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toUpperCase",
							call: builtinStringToUpperCase,
						},
					},
				},
			},
			"valueOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "valueOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinStringValueOf,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyConstructor,
			"charAt",
			"charCodeAt",
			"concat",
			"indexOf",
			"lastIndexOf",
			"localeCompare",
			"match",
			"replace",
			"search",
			"slice",
			"split",
			"substr",
			"substring",
			methodToString,
			"trim",
			"trimLeft",
			"trimRight",
			"toLocaleLowerCase",
			"toLocaleUpperCase",
			"toLowerCase",
			"toUpperCase",
			"valueOf",
		},
	}

	// String definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.StringPrototype,
				},
			},
			"fromCharCode": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "fromCharCode",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "fromCharCode",
							call: builtinStringFromCharCode,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
			"fromCharCode",
		},
	}

	// String constructor definition.
	rt.global.StringPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.String,
		},
	}

	// Boolean prototype.
	rt.global.BooleanPrototype = &object{
		runtime:     rt,
		class:       classBooleanName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		value:       prototypeValueBoolean,
		property: map[string]property{
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinBooleanToString,
						},
					},
				},
			},
			"valueOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "valueOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinBooleanValueOf,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			methodToString,
			"valueOf",
		},
	}

	// Boolean definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.BooleanPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// Boolean constructor definition.
	rt.global.BooleanPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Boolean,
		},
	}

	// Number prototype.
	rt.global.NumberPrototype = &object{
		runtime:     rt,
		class:       classNumberName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		value:       prototypeValueNumber,
		property: map[string]property{
			"toExponential": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toExponential",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toExponential",
							call: builtinNumberToExponential,
						},
					},
				},
			},
			"toFixed": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toFixed",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toFixed",
							call: builtinNumberToFixed,
						},
					},
				},
			},
			"toPrecision": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toPrecision",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toPrecision",
							call: builtinNumberToPrecision,
						},
					},
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinNumberToString,
						},
					},
				},
			},
			"valueOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "valueOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinNumberValueOf,
						},
					},
				},
			},
			"toLocaleString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleString",
							call: builtinNumberToLocaleString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"toExponential",
			"toFixed",
			"toPrecision",
			methodToString,
			"valueOf",
			"toLocaleString",
		},
	}

	// Number definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.NumberPrototype,
				},
			},
			"isNaN": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "isNaN",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "isNaN",
							call: builtinNumberIsNaN,
						},
					},
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
			propertyPrototype,
			"isNaN",
			"MAX_VALUE",
			"MIN_VALUE",
			"NaN",
			"NEGATIVE_INFINITY",
			"POSITIVE_INFINITY",
		},
	}

	// Number constructor definition.
	rt.global.NumberPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Number,
		},
	}

	// Math definition.
	rt.global.Math = &object{
		runtime:     rt,
		class:       classMathName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		property: map[string]property{
			"abs": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "abs",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "abs",
							call: builtinMathAbs,
						},
					},
				},
			},
			"acos": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "acos",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "acos",
							call: builtinMathAcos,
						},
					},
				},
			},
			"asin": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "asin",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "asin",
							call: builtinMathAsin,
						},
					},
				},
			},
			"atan": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "atan",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "atan",
							call: builtinMathAtan,
						},
					},
				},
			},
			"atan2": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "atan2",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "atan2",
							call: builtinMathAtan2,
						},
					},
				},
			},
			"ceil": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "ceil",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "ceil",
							call: builtinMathCeil,
						},
					},
				},
			},
			"cos": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "cos",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "cos",
							call: builtinMathCos,
						},
					},
				},
			},
			"exp": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "exp",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "exp",
							call: builtinMathExp,
						},
					},
				},
			},
			"floor": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "floor",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "floor",
							call: builtinMathFloor,
						},
					},
				},
			},
			"log": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "log",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "log",
							call: builtinMathLog,
						},
					},
				},
			},
			"max": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "max",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "max",
							call: builtinMathMax,
						},
					},
				},
			},
			"min": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "min",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "min",
							call: builtinMathMin,
						},
					},
				},
			},
			"pow": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "pow",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "pow",
							call: builtinMathPow,
						},
					},
				},
			},
			"random": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "random",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "random",
							call: builtinMathRandom,
						},
					},
				},
			},
			"round": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "round",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "round",
							call: builtinMathRound,
						},
					},
				},
			},
			"sin": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "sin",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "sin",
							call: builtinMathSin,
						},
					},
				},
			},
			"sqrt": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "sqrt",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "sqrt",
							call: builtinMathSqrt,
						},
					},
				},
			},
			"tan": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "tan",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "tan",
							call: builtinMathTan,
						},
					},
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
			"LOG10E": {
				mode: 0,
				value: Value{
					kind:  valueNumber,
					value: math.Log10E,
				},
			},
			"LOG2E": {
				mode: 0,
				value: Value{
					kind:  valueNumber,
					value: math.Log2E,
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
			"LOG10E",
			"LOG2E",
			"PI",
			"SQRT1_2",
			"SQRT2",
		},
	}

	// Date prototype.
	rt.global.DatePrototype = &object{
		runtime:     rt,
		class:       classDateName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		value:       prototypeValueDate,
		property: map[string]property{
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinDateToString,
						},
					},
				},
			},
			"toDateString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toDateString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toDateString",
							call: builtinDateToDateString,
						},
					},
				},
			},
			"toTimeString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toTimeString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toTimeString",
							call: builtinDateToTimeString,
						},
					},
				},
			},
			"toISOString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toISOString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toISOString",
							call: builtinDateToISOString,
						},
					},
				},
			},
			"toUTCString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toUTCString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toUTCString",
							call: builtinDateToUTCString,
						},
					},
				},
			},
			"toGMTString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toGMTString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toGMTString",
							call: builtinDateToGMTString,
						},
					},
				},
			},
			"getDate": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getDate",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getDate",
							call: builtinDateGetDate,
						},
					},
				},
			},
			"setDate": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setDate",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setDate",
							call: builtinDateSetDate,
						},
					},
				},
			},
			"getDay": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getDay",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getDay",
							call: builtinDateGetDay,
						},
					},
				},
			},
			"getFullYear": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getFullYear",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getFullYear",
							call: builtinDateGetFullYear,
						},
					},
				},
			},
			"setFullYear": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setFullYear",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setFullYear",
							call: builtinDateSetFullYear,
						},
					},
				},
			},
			"getHours": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getHours",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getHours",
							call: builtinDateGetHours,
						},
					},
				},
			},
			"setHours": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setHours",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setHours",
							call: builtinDateSetHours,
						},
					},
				},
			},
			"getMilliseconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getMilliseconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getMilliseconds",
							call: builtinDateGetMilliseconds,
						},
					},
				},
			},
			"setMilliseconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setMilliseconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setMilliseconds",
							call: builtinDateSetMilliseconds,
						},
					},
				},
			},
			"getMinutes": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getMinutes",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getMinutes",
							call: builtinDateGetMinutes,
						},
					},
				},
			},
			"setMinutes": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setMinutes",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setMinutes",
							call: builtinDateSetMinutes,
						},
					},
				},
			},
			"getMonth": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getMonth",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getMonth",
							call: builtinDateGetMonth,
						},
					},
				},
			},
			"setMonth": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setMonth",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setMonth",
							call: builtinDateSetMonth,
						},
					},
				},
			},
			"getSeconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getSeconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getSeconds",
							call: builtinDateGetSeconds,
						},
					},
				},
			},
			"setSeconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setSeconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setSeconds",
							call: builtinDateSetSeconds,
						},
					},
				},
			},
			"getTime": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getTime",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getTime",
							call: builtinDateGetTime,
						},
					},
				},
			},
			"setTime": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setTime",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setTime",
							call: builtinDateSetTime,
						},
					},
				},
			},
			"getTimezoneOffset": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getTimezoneOffset",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getTimezoneOffset",
							call: builtinDateGetTimezoneOffset,
						},
					},
				},
			},
			"getUTCDate": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCDate",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCDate",
							call: builtinDateGetUTCDate,
						},
					},
				},
			},
			"setUTCDate": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCDate",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCDate",
							call: builtinDateSetUTCDate,
						},
					},
				},
			},
			"getUTCDay": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCDay",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCDay",
							call: builtinDateGetUTCDay,
						},
					},
				},
			},
			"getUTCFullYear": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCFullYear",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCFullYear",
							call: builtinDateGetUTCFullYear,
						},
					},
				},
			},
			"setUTCFullYear": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCFullYear",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCFullYear",
							call: builtinDateSetUTCFullYear,
						},
					},
				},
			},
			"getUTCHours": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCHours",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCHours",
							call: builtinDateGetUTCHours,
						},
					},
				},
			},
			"setUTCHours": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCHours",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCHours",
							call: builtinDateSetUTCHours,
						},
					},
				},
			},
			"getUTCMilliseconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCMilliseconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCMilliseconds",
							call: builtinDateGetUTCMilliseconds,
						},
					},
				},
			},
			"setUTCMilliseconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCMilliseconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCMilliseconds",
							call: builtinDateSetUTCMilliseconds,
						},
					},
				},
			},
			"getUTCMinutes": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCMinutes",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCMinutes",
							call: builtinDateGetUTCMinutes,
						},
					},
				},
			},
			"setUTCMinutes": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCMinutes",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCMinutes",
							call: builtinDateSetUTCMinutes,
						},
					},
				},
			},
			"getUTCMonth": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCMonth",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCMonth",
							call: builtinDateGetUTCMonth,
						},
					},
				},
			},
			"setUTCMonth": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCMonth",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCMonth",
							call: builtinDateSetUTCMonth,
						},
					},
				},
			},
			"getUTCSeconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getUTCSeconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getUTCSeconds",
							call: builtinDateGetUTCSeconds,
						},
					},
				},
			},
			"setUTCSeconds": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setUTCSeconds",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setUTCSeconds",
							call: builtinDateSetUTCSeconds,
						},
					},
				},
			},
			"valueOf": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "valueOf",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinDateValueOf,
						},
					},
				},
			},
			"getYear": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "getYear",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "getYear",
							call: builtinDateGetYear,
						},
					},
				},
			},
			"setYear": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "setYear",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "setYear",
							call: builtinDateSetYear,
						},
					},
				},
			},
			"toJSON": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toJSON",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toJSON",
							call: builtinDateToJSON,
						},
					},
				},
			},
			"toLocaleString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleString",
							call: builtinDateToLocaleString,
						},
					},
				},
			},
			"toLocaleDateString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleDateString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleDateString",
							call: builtinDateToLocaleDateString,
						},
					},
				},
			},
			"toLocaleTimeString": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toLocaleTimeString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "toLocaleTimeString",
							call: builtinDateToLocaleTimeString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			methodToString,
			"toDateString",
			"toTimeString",
			"toISOString",
			"toUTCString",
			"toGMTString",
			"getDate",
			"setDate",
			"getDay",
			"getFullYear",
			"setFullYear",
			"getHours",
			"setHours",
			"getMilliseconds",
			"setMilliseconds",
			"getMinutes",
			"setMinutes",
			"getMonth",
			"setMonth",
			"getSeconds",
			"setSeconds",
			"getTime",
			"setTime",
			"getTimezoneOffset",
			"getUTCDate",
			"setUTCDate",
			"getUTCDay",
			"getUTCFullYear",
			"setUTCFullYear",
			"getUTCHours",
			"setUTCHours",
			"getUTCMilliseconds",
			"setUTCMilliseconds",
			"getUTCMinutes",
			"setUTCMinutes",
			"getUTCMonth",
			"setUTCMonth",
			"getUTCSeconds",
			"setUTCSeconds",
			"valueOf",
			"getYear",
			"setYear",
			"toJSON",
			"toLocaleString",
			"toLocaleDateString",
			"toLocaleTimeString",
		},
	}

	// Date definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.DatePrototype,
				},
			},
			"parse": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "parse",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "parse",
							call: builtinDateParse,
						},
					},
				},
			},
			"UTC": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "UTC",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "UTC",
							call: builtinDateUTC,
						},
					},
				},
			},
			"now": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "now",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "now",
							call: builtinDateNow,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
			"parse",
			"UTC",
			"now",
		},
	}

	// Date constructor definition.
	rt.global.DatePrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Date,
		},
	}

	// RegExp prototype.
	rt.global.RegExpPrototype = &object{
		runtime:     rt,
		class:       classRegExpName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		value:       prototypeValueRegExp,
		property: map[string]property{
			"exec": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "exec",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "exec",
							call: builtinRegExpExec,
						},
					},
				},
			},
			"compile": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "compile",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "compile",
							call: builtinRegExpCompile,
						},
					},
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinRegExpToString,
						},
					},
				},
			},
			"test": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "test",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "test",
							call: builtinRegExpTest,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"exec",
			"compile",
			methodToString,
			"test",
		},
	}

	// RegExp definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.RegExpPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// RegExp constructor definition.
	rt.global.RegExpPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.RegExp,
		},
	}

	// Error prototype.
	rt.global.ErrorPrototype = &object{
		runtime:     rt,
		class:       classErrorName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
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
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// Error definition.
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.ErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// Error constructor definition.
	rt.global.ErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Error,
		},
	}

	// EvalError prototype.
	rt.global.EvalErrorPrototype = &object{
		runtime:     rt,
		class:       classEvalErrorName,
		objectClass: classObject,
		prototype:   rt.global.ErrorPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
			"name": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: classEvalErrorName,
				},
			},
			"message": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: "",
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// EvalError definition.
	rt.global.EvalError = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.FunctionPrototype,
		extensible:  true,
		value: nativeFunctionObject{
			name:      classEvalErrorName,
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.EvalErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// EvalError constructor definition.
	rt.global.EvalErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.EvalError,
		},
	}

	// TypeError prototype.
	rt.global.TypeErrorPrototype = &object{
		runtime:     rt,
		class:       classTypeErrorName,
		objectClass: classObject,
		prototype:   rt.global.ErrorPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
			"name": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: classTypeErrorName,
				},
			},
			"message": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: "",
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// TypeError definition.
	rt.global.TypeError = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.FunctionPrototype,
		extensible:  true,
		value: nativeFunctionObject{
			name:      classTypeErrorName,
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.TypeErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// TypeError constructor definition.
	rt.global.TypeErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.TypeError,
		},
	}

	// RangeError prototype.
	rt.global.RangeErrorPrototype = &object{
		runtime:     rt,
		class:       classRangeErrorName,
		objectClass: classObject,
		prototype:   rt.global.ErrorPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
			"name": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: classRangeErrorName,
				},
			},
			"message": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: "",
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// RangeError definition.
	rt.global.RangeError = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.FunctionPrototype,
		extensible:  true,
		value: nativeFunctionObject{
			name:      classRangeErrorName,
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.RangeErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// RangeError constructor definition.
	rt.global.RangeErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.RangeError,
		},
	}

	// ReferenceError prototype.
	rt.global.ReferenceErrorPrototype = &object{
		runtime:     rt,
		class:       classReferenceErrorName,
		objectClass: classObject,
		prototype:   rt.global.ErrorPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
			"name": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: classReferenceErrorName,
				},
			},
			"message": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: "",
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// ReferenceError definition.
	rt.global.ReferenceError = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.FunctionPrototype,
		extensible:  true,
		value: nativeFunctionObject{
			name:      classReferenceErrorName,
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.ReferenceErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// ReferenceError constructor definition.
	rt.global.ReferenceErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.ReferenceError,
		},
	}

	// SyntaxError prototype.
	rt.global.SyntaxErrorPrototype = &object{
		runtime:     rt,
		class:       classSyntaxErrorName,
		objectClass: classObject,
		prototype:   rt.global.ErrorPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
			"name": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: classSyntaxErrorName,
				},
			},
			"message": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: "",
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// SyntaxError definition.
	rt.global.SyntaxError = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.FunctionPrototype,
		extensible:  true,
		value: nativeFunctionObject{
			name:      classSyntaxErrorName,
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.SyntaxErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// SyntaxError constructor definition.
	rt.global.SyntaxErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.SyntaxError,
		},
	}

	// URIError prototype.
	rt.global.URIErrorPrototype = &object{
		runtime:     rt,
		class:       classURIErrorName,
		objectClass: classObject,
		prototype:   rt.global.ErrorPrototype,
		extensible:  true,
		value:       nil,
		property: map[string]property{
			"name": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: classURIErrorName,
				},
			},
			"message": {
				mode: 0o101,
				value: Value{
					kind:  valueString,
					value: "",
				},
			},
			methodToString: {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "toString",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyConstructor,
			"name",
			"message",
			methodToString,
		},
	}

	// URIError definition.
	rt.global.URIError = &object{
		runtime:     rt,
		class:       classFunctionName,
		objectClass: classObject,
		prototype:   rt.global.FunctionPrototype,
		extensible:  true,
		value: nativeFunctionObject{
			name:      classURIErrorName,
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
			propertyPrototype: {
				mode: 0,
				value: Value{
					kind:  valueObject,
					value: rt.global.URIErrorPrototype,
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			propertyPrototype,
		},
	}

	// URIError constructor definition.
	rt.global.URIErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.URIError,
		},
	}

	// JSON definition.
	rt.global.JSON = &object{
		runtime:     rt,
		class:       classJSONName,
		objectClass: classObject,
		prototype:   rt.global.ObjectPrototype,
		extensible:  true,
		property: map[string]property{
			"parse": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "parse",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "parse",
							call: builtinJSONParse,
						},
					},
				},
			},
			"stringify": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "stringify",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "stringify",
							call: builtinJSONStringify,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			"parse",
			"stringify",
		},
	}

	// Global properties.
	rt.globalObject.property = map[string]property{
		"eval": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "eval",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "eval",
						call: builtinGlobalEval,
					},
				},
			},
		},
		"parseInt": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "parseInt",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "parseInt",
						call: builtinGlobalParseInt,
					},
				},
			},
		},
		"parseFloat": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "parseFloat",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "parseFloat",
						call: builtinGlobalParseFloat,
					},
				},
			},
		},
		"isNaN": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "isNaN",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "isNaN",
						call: builtinGlobalIsNaN,
					},
				},
			},
		},
		"isFinite": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "isFinite",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "isFinite",
						call: builtinGlobalIsFinite,
					},
				},
			},
		},
		"decodeURI": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "decodeURI",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "decodeURI",
						call: builtinGlobalDecodeURI,
					},
				},
			},
		},
		"decodeURIComponent": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "decodeURIComponent",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "decodeURIComponent",
						call: builtinGlobalDecodeURIComponent,
					},
				},
			},
		},
		"encodeURI": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "encodeURI",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "encodeURI",
						call: builtinGlobalEncodeURI,
					},
				},
			},
		},
		"encodeURIComponent": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "encodeURIComponent",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "encodeURIComponent",
						call: builtinGlobalEncodeURIComponent,
					},
				},
			},
		},
		"escape": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "escape",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "escape",
						call: builtinGlobalEscape,
					},
				},
			},
		},
		"unescape": {
			mode: 0o101,
			value: Value{
				kind: valueObject,
				value: &object{
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
						propertyName: {
							mode: 0,
							value: Value{
								kind:  valueString,
								value: "unescape",
							},
						},
					},
					propertyOrder: []string{
						propertyLength,
						propertyName,
					},
					value: nativeFunctionObject{
						name: "unescape",
						call: builtinGlobalUnescape,
					},
				},
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
		classMathName: {
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
		classEvalErrorName: {
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.EvalError,
			},
		},
		classTypeErrorName: {
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.TypeError,
			},
		},
		classRangeErrorName: {
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.RangeError,
			},
		},
		classReferenceErrorName: {
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.ReferenceError,
			},
		},
		classSyntaxErrorName: {
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.SyntaxError,
			},
		},
		classURIErrorName: {
			mode: 0o101,
			value: Value{
				kind:  valueObject,
				value: rt.global.URIError,
			},
		},
		classJSONName: {
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

	// Global property order.
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
		classMathName,
		classDateName,
		classRegExpName,
		classErrorName,
		classEvalErrorName,
		classTypeErrorName,
		classRangeErrorName,
		classReferenceErrorName,
		classSyntaxErrorName,
		classURIErrorName,
		classJSONName,
		"undefined",
		"NaN",
		"Infinity",
	}
}

func (rt *runtime) newConsole() *object {
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
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "log",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "log",
							call: builtinConsoleLog,
						},
					},
				},
			},
			"debug": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "debug",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "debug",
							call: builtinConsoleLog,
						},
					},
				},
			},
			"info": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "info",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "info",
							call: builtinConsoleLog,
						},
					},
				},
			},
			"error": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "error",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "error",
							call: builtinConsoleError,
						},
					},
				},
			},
			"warn": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "warn",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "warn",
							call: builtinConsoleError,
						},
					},
				},
			},
			"dir": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "dir",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "dir",
							call: builtinConsoleDir,
						},
					},
				},
			},
			"time": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "time",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "time",
							call: builtinConsoleTime,
						},
					},
				},
			},
			"timeEnd": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "timeEnd",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "timeEnd",
							call: builtinConsoleTimeEnd,
						},
					},
				},
			},
			"trace": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "trace",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "trace",
							call: builtinConsoleTrace,
						},
					},
				},
			},
			"assert": {
				mode: 0o101,
				value: Value{
					kind: valueObject,
					value: &object{
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
							propertyName: {
								mode: 0,
								value: Value{
									kind:  valueString,
									value: "assert",
								},
							},
						},
						propertyOrder: []string{
							propertyLength,
							propertyName,
						},
						value: nativeFunctionObject{
							name: "assert",
							call: builtinConsoleAssert,
						},
					},
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
