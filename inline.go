package otto

import (
	"math"
)

func (rt *runtime) newContext() {
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
	rt.global.ObjectPrototype.property = map[string]property{
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
					},
					propertyOrder: []string{
						propertyLength,
					},
					value: nativeFunctionObject{
						name: "valueOf",
						call: builtinObjectValueOf,
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
					},
					propertyOrder: []string{
						propertyLength,
					},
					value: nativeFunctionObject{
						name: methodToString,
						call: builtinObjectToString,
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
					},
					propertyOrder: []string{
						propertyLength,
					},
					value: nativeFunctionObject{
						name: "toLocaleString",
						call: builtinObjectToLocaleString,
					},
				},
			},
		},
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
					},
					value: nativeFunctionObject{
						name: "propertyIsEnumerable",
						call: builtinObjectPropertyIsEnumerable,
					},
				},
			},
		},
		propertyConstructor: {
			mode:  0o101,
			value: Value{},
		},
	}
	rt.global.ObjectPrototype.propertyOrder = []string{
		"valueOf",
		methodToString,
		"toLocaleString",
		"hasOwnProperty",
		"isPrototypeOf",
		"propertyIsEnumerable",
		propertyConstructor,
	}
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
					},
					value: nativeFunctionObject{
						name: "bind",
						call: builtinFunctionBind,
					},
				},
			},
		},
		propertyConstructor: {
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
		methodToString,
		"apply",
		"call",
		"bind",
		propertyConstructor,
		propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
	rt.global.ObjectPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Object,
		},
	}
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
	rt.global.FunctionPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Function,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinArrayToString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toLocaleString",
							call: builtinArrayToLocaleString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "concat",
							call: builtinArrayConcat,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "join",
							call: builtinArrayJoin,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "splice",
							call: builtinArraySplice,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "shift",
							call: builtinArrayShift,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "push",
							call: builtinArrayPush,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "slice",
							call: builtinArraySlice,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "unshift",
							call: builtinArrayUnshift,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "reverse",
							call: builtinArrayReverse,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "sort",
							call: builtinArraySort,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "indexOf",
							call: builtinArrayIndexOf,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "lastIndexOf",
							call: builtinArrayLastIndexOf,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "some",
							call: builtinArraySome,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "forEach",
							call: builtinArrayForEach,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "map",
							call: builtinArrayMap,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "filter",
							call: builtinArrayFilter,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "reduceRight",
							call: builtinArrayReduceRight,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			methodToString,
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
						},
						propertyOrder: []string{
							propertyLength,
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
	rt.global.ArrayPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Array,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinStringToString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinStringValueOf,
						},
					},
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "lastIndexOf",
							call: builtinStringLastIndexOf,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "search",
							call: builtinStringSearch,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "split",
							call: builtinStringSplit,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "slice",
							call: builtinStringSlice,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "substring",
							call: builtinStringSubstring,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toUpperCase",
							call: builtinStringToUpperCase,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "substr",
							call: builtinStringSubstr,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "trimRight",
							call: builtinStringTrimRight,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "localeCompare",
							call: builtinStringLocaleCompare,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toLocaleUpperCase",
							call: builtinStringToLocaleUpperCase,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			propertyLength,
			methodToString,
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
						},
						propertyOrder: []string{
							propertyLength,
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
	rt.global.StringPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.String,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
			methodToString,
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
	rt.global.BooleanPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Boolean,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinNumberValueOf,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toFixed",
							call: builtinNumberToFixed,
						},
					},
				},
			},
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toExponential",
							call: builtinNumberToExponential,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toPrecision",
							call: builtinNumberToPrecision,
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
						},
						propertyOrder: []string{
							propertyLength,
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
			methodToString,
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
						},
						propertyOrder: []string{
							propertyLength,
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
	rt.global.NumberPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Number,
		},
	}
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toTimeString",
							call: builtinDateToTimeString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toUTCString",
							call: builtinDateToUTCString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toISOString",
							call: builtinDateToISOString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toJSON",
							call: builtinDateToJSON,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toGMTString",
							call: builtinDateToGMTString,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "toLocaleTimeString",
							call: builtinDateToLocaleTimeString,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "valueOf",
							call: builtinDateValueOf,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getTime",
							call: builtinDateGetTime,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getYear",
							call: builtinDateGetYear,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getFullYear",
							call: builtinDateGetFullYear,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCFullYear",
							call: builtinDateGetUTCFullYear,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getMonth",
							call: builtinDateGetMonth,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCMonth",
							call: builtinDateGetUTCMonth,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getDate",
							call: builtinDateGetDate,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCDate",
							call: builtinDateGetUTCDate,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getDay",
							call: builtinDateGetDay,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCDay",
							call: builtinDateGetUTCDay,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getHours",
							call: builtinDateGetHours,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCHours",
							call: builtinDateGetUTCHours,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getMinutes",
							call: builtinDateGetMinutes,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCMinutes",
							call: builtinDateGetUTCMinutes,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getSeconds",
							call: builtinDateGetSeconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCSeconds",
							call: builtinDateGetUTCSeconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getMilliseconds",
							call: builtinDateGetMilliseconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getUTCMilliseconds",
							call: builtinDateGetUTCMilliseconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "getTimezoneOffset",
							call: builtinDateGetTimezoneOffset,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setTime",
							call: builtinDateSetTime,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setMilliseconds",
							call: builtinDateSetMilliseconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCMilliseconds",
							call: builtinDateSetUTCMilliseconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setSeconds",
							call: builtinDateSetSeconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCSeconds",
							call: builtinDateSetUTCSeconds,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setMinutes",
							call: builtinDateSetMinutes,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCMinutes",
							call: builtinDateSetUTCMinutes,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setHours",
							call: builtinDateSetHours,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCHours",
							call: builtinDateSetUTCHours,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setDate",
							call: builtinDateSetDate,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCDate",
							call: builtinDateSetUTCDate,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setMonth",
							call: builtinDateSetMonth,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCMonth",
							call: builtinDateSetUTCMonth,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setYear",
							call: builtinDateSetYear,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setFullYear",
							call: builtinDateSetFullYear,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "setUTCFullYear",
							call: builtinDateSetUTCFullYear,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			methodToString,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
	rt.global.DatePrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Date,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinRegExpToString,
						},
					},
				},
			},
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "exec",
							call: builtinRegExpExec,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "test",
							call: builtinRegExpTest,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: "compile",
							call: builtinRegExpCompile,
						},
					},
				},
			},
		},
		propertyOrder: []string{
			methodToString,
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
	rt.global.RegExpPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.RegExp,
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
						},
						propertyOrder: []string{
							propertyLength,
						},
						value: nativeFunctionObject{
							name: methodToString,
							call: builtinErrorToString,
						},
					},
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
			methodToString,
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
	rt.global.ErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.Error,
		},
	}
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
	rt.global.EvalErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.EvalError,
		},
	}
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
	rt.global.TypeErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.TypeError,
		},
	}
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
	rt.global.RangeErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.RangeError,
		},
	}
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
	rt.global.ReferenceErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.ReferenceError,
		},
	}
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
	rt.global.SyntaxErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.SyntaxError,
		},
	}
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
	rt.global.URIErrorPrototype.property[propertyConstructor] = property{
		mode: 0o101,
		value: Value{
			kind:  valueObject,
			value: rt.global.URIError,
		},
	}
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
					},
					propertyOrder: []string{
						propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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
						},
						propertyOrder: []string{
							propertyLength,
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

func int8Value(value int8) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func int16Value(value int16) Value {
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

func uintValue(value uint) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func uint8Value(value uint8) Value {
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

func uint64Value(value uint64) Value {
	return Value{
		kind:  valueNumber,
		value: value,
	}
}

func float32Value(value float32) Value {
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
