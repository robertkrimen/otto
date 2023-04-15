package otto

// newConsoleManual is a manually implemented "fork" of newConsole which is generated code.
func (rt *runtime) newConsoleManual() *object {
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
