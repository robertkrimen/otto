package otto

import (
	"testing"
)

func TestObject_(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		obj := newObject(nil, "")
		is(obj != nil, true)

		obj.put("xyzzy", toValue("Nothing happens."), true)
		is(obj.get("xyzzy"), "Nothing happens.")

		test(`
            var abc = Object.getOwnPropertyDescriptor(Object, "prototype");
            [ [ typeof Object.prototype, abc.writable, abc.enumerable, abc.configurable ],
            ];
        `, "object,false,false,false")
	})
}

func TestStringObject(t *testing.T) {
	tt(t, func() {
		obj := New().runtime.newStringObject(toValue("xyzzy"))
		is(obj.get("1"), "y")
		is(obj.get("10"), "undefined")
		is(obj.get("2"), "z")
	})
}

func TestObject_getPrototypeOf(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = {};
            def = Object.getPrototypeOf(abc);
            ghi = Object.getPrototypeOf(def);
            [abc,def,ghi,ghi+""];
        `, "[object Object],[object Object],,null")

		test(`
            abc = Object.getOwnPropertyDescriptor(Object, "getPrototypeOf");
            [ abc.value === Object.getPrototypeOf, abc.writable, abc.enumerable, abc.configurable ];
        `, "true,true,false,true")
	})
}

func TestObject_new(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            [ new Object("abc"), new Object(2+2) ];
        `, "abc,4")
	})
}

func TestObject_create(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise: Object.create()`, "TypeError: Object.Create is nil")

		test(`
            var abc = Object.create(null)
            var def = Object.create({x: 10, y: 20})
            var ghi = Object.create(Object.prototype)

            var jkl = Object.create({x: 10, y: 20}, {
                z: {
                    value: 30,
                    writable: true
                },
                // sum: {
                //     get: function() {
                //         return this.x + this.y + this.z
                //     }
                // }
            });
            [ abc.prototype, def.x, def.y, ghi, jkl.x, jkl.y, jkl.z ]
        `, ",10,20,[object Object],10,20,30")

		test(`
            var properties = {};
            Object.defineProperty(properties, "abc", {
                value: {},
                enumerable: false
            });
            var mno = Object.create({}, properties);
            mno.hasOwnProperty("abc");
        `, false)
	})
}

func TestObject_toLocaleString(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            ({}).toLocaleString();
        `, "[object Object]")

		test(`
            object = {
                toString: function() {
                    return "Nothing happens.";
                }
            };
            object.toLocaleString();
        `, "Nothing happens.")
	})
}

func TestObject_isExtensible(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise:
            Object.isExtensible();
        `, "TypeError: Object.IsExtensible is nil")

		// FIXME terst, Why raise?
		test(`raise:
            Object.isExtensible({});
        `, true)

		test(`Object.isExtensible.length`, 1)
		test(`Object.isExtensible.prototype`, "undefined")
	})
}

func TestObject_preventExtensions(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise:
            Object.preventExtensions()
        `, "TypeError: Object.PreventExtensions is nil")

		test(`raise:
            var abc = { def: true };
            var ghi = Object.preventExtensions(abc);
            [ ghi.def === true, Object.isExtensible(abc), Object.isExtensible(ghi) ];
        `, "true,false,false")

		test(`
            var abc = new String();
            var def = Object.isExtensible(abc);
            Object.preventExtensions(abc);
            var ghi = false;
            try {
                Object.defineProperty(abc, "0", { value: "~" });
            } catch (err) {
                ghi = err instanceof TypeError;
            }
            [ def, ghi, abc.hasOwnProperty("0"), typeof abc[0] ];
        `, "true,true,false,undefined")

		test(`Object.preventExtensions.length`, 1)
		test(`Object.preventExtensions.prototype`, "undefined")
	})
}

func TestObject_isSealed(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Object.isSealed.length`, 1)
		test(`Object.isSealed.prototype`, "undefined")
	})
}

func TestObject_seal(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise: Object.seal()`, "TypeError: Object.Seal is nil")

		test(`
            var abc = {a:1,b:1,c:3};
            var sealed = Object.isSealed(abc);
            Object.seal(abc);
            [sealed, Object.isSealed(abc)];
        `, "false,true")

		test(`
            var abc = {a:1,b:1,c:3};
            var sealed = Object.isSealed(abc);
            var caught = false;
            Object.seal(abc);
            abc.b = 5;
            Object.defineProperty(abc, "a", {value:4});
            try {
                Object.defineProperty(abc, "a", {value:42,enumerable:false});
            } catch (e) {
                caught = e instanceof TypeError;
            }
            [sealed, Object.isSealed(abc), caught, abc.a, abc.b];
        `, "false,true,true,4,5")

		test(`Object.seal.length`, 1)
		test(`Object.seal.prototype`, "undefined")
	})
}

func TestObject_isFrozen(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise: Object.isFrozen()`, "TypeError: Object.IsFrozen is nil")
		test(`Object.isFrozen(Object.preventExtensions({a:1}))`, false)
		test(`Object.isFrozen({})`, false)

		test(`
            var abc = {};
            Object.defineProperty(abc, "def", {
                value: "def",
                writable: true,
                configurable: false
            });
            Object.preventExtensions(abc);
            !Object.isFrozen(abc);
        `, true)

		test(`Object.isFrozen.length`, 1)
		test(`Object.isFrozen.prototype`, "undefined")
	})
}

func TestObject_freeze(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise: Object.freeze()`, "TypeError: Object.Freeze is nil")

		test(`
            var abc = {a:1,b:2,c:3};
            var frozen = Object.isFrozen(abc);
            Object.freeze(abc);
            abc.b = 5;
            [frozen, Object.isFrozen(abc), abc.b];
        `, "false,true,2")

		test(`
            var abc = {a:1,b:2,c:3};
            var frozen = Object.isFrozen(abc);
            var caught = false;
            Object.freeze(abc);
            abc.b = 5;
            try {
                Object.defineProperty(abc, "a", {value:4});
            } catch (e) {
                caught = e instanceof TypeError;
            }
            [frozen, Object.isFrozen(abc), caught, abc.a, abc.b];
        `, "false,true,true,1,2")

		test(`Object.freeze.length`, 1)
		test(`Object.freeze.prototype`, "undefined")
	})
}

func TestObject_defineProperty(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            (function(abc, def, ghi){
                Object.defineProperty(arguments, "0", {
                    enumerable: false
                });
                return true;
            })(0, 1, 2);
        `, true)

		test(`
            var abc = {};
            abc.def = 3.14; // Default: writable: true, enumerable: true, configurable: true

            Object.defineProperty(abc, "def", {
                value: 42
            });

            var ghi = Object.getOwnPropertyDescriptor(abc, "def");
            [ ghi.value, ghi.writable, ghi.enumerable, ghi.configurable ];
        `, "42,true,true,true")

		// Test that we handle the case of DefineOwnProperty
		// where [[Writable]] is something but [[Value]] is not
		test(`
            var abc = [];
            Object.defineProperty(abc, "0", { writable: false });
            0 in abc;
        `, true)

		// Test that we handle the case of DefineOwnProperty
		// where [[Writable]] is something but [[Value]] is not
		// (and the property originally had something for [[Value]]
		test(`
            abc = {
                def: 42
            };
            Object.defineProperty(abc, "def", { writable: false });
            abc.def;
        `, 42)
	})
}

func TestObject_assign(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		// Test 1: Assigning two empty objects should return an empty object.
		test(`JSON.stringify(Object.assign({}, {}))`, "{}")

		// Test 2: Single source assignment.
		test(`JSON.stringify(Object.assign({}, {a: 1}))`, "{\"a\":1}")

		// Test 3: Multiple sources with later properties overriding earlier ones.
		test(`JSON.stringify(Object.assign({a: 1, c: 5}, {a: 2}, {b: 3}))`, "{\"a\":2,\"b\":3,\"c\":5}")

		// Test 4: Merging objects with overlapping keys.
		test(`JSON.stringify(Object.assign({a: 1, b: 2}, {b: 3, c: 4}))`, "{\"a\":1,\"b\":3,\"c\":4}")

		// Test 5: Sources that are null or undefined should be ignored.
		test(`JSON.stringify(Object.assign({a: 1}, null, undefined, {b: 2}))`, "{\"a\":1,\"b\":2}")

		// Test 6: When a string is used as a source, its characters are assigned as indexed properties.
		test(`JSON.stringify(Object.assign({}, "abc"))`, "{\"0\":\"a\",\"1\":\"b\",\"2\":\"c\"}")

		// // Test 7: The return value should be the target object.
		test(`(function(){ var o = {x:1}; var r = Object.assign(o, {y:2}); return (o === r).toString(); })()`, "true")

		// Test 8: Non-enumerable properties should not be copied.
		test(`(function(){
            var target = {};
            var source = {};
            Object.defineProperty(source, "hidden", { value: 42, enumerable: false });
            Object.assign(target, source);
            return target.hasOwnProperty("hidden").toString();
		})()`, "false")

		// Test 9: Using a number as a source should not add any properties.
		test(`JSON.stringify(Object.assign({}, 123))`, "{}")

		// Test 10: Using a boolean as a source should not add any properties.
		test(`JSON.stringify(Object.assign({}, true))`, "{}")

		// Test 11: Arrays are objects, so their indexed elements are copied.
		test(`JSON.stringify(Object.assign({}, [1,2,3]))`, "{\"0\":1,\"1\":2,\"2\":3}")
	})
}

func TestObject_keys(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Object.keys({ abc:undefined, def:undefined })`, "abc,def")

		test(`
            function abc() {
                this.abc = undefined;
                this.def = undefined;
            }
            Object.keys(new abc())
        `, "abc,def")

		test(`
            function def() {
                this.ghi = undefined;
            }
            def.prototype = new abc();
            Object.keys(new def());
        `, "ghi")

		test(`
            var ghi = Object.create(
                {
                    abc: undefined,
                    def: undefined
                },
                {
                    ghi: { value: undefined, enumerable: true },
                    jkl: { value: undefined, enumerable: false }
                }
            );
            Object.keys(ghi);
        `, "ghi")

		test(`
            (function(abc, def, ghi){
                return Object.keys(arguments)
            })(undefined, undefined);
        `, "0,1")

		test(`
            (function(abc, def, ghi){
                return Object.keys(arguments)
            })(undefined, undefined, undefined, undefined);
        `, "0,1,2,3")
	})
}

func TestObject_values(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Object.values({ abc:"first_example", def:"second_example" })`, "first_example,second_example")

		test(`
			function abc() {
				this.abc = "first_example";
				this.def = "second_example";
			}
			Object.values(new abc())
		`, "first_example,second_example")

		test(`
			function def() {
				this.ghi = "third_example"
			}
			def.prototype = new abc();
			Object.values(new def());
		`, "third_example")

		test(`
			var arr = [1, 2, 3];
			Object.values(arr);
		`, "1,2,3")

		test(`
			var arr = [{"abc": "first_example"}, {"def": "second_example"}];
			Object.values(arr);
		`, "[object Object],[object Object]")
	})
}

func TestObject_getOwnPropertyNames(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Object.getOwnPropertyNames({ abc:undefined, def:undefined })`, "abc,def")

		test(`
            var ghi = Object.create(
                {
                    abc: undefined,
                    def: undefined
                },
                {
                    ghi: { value: undefined, enumerable: true },
                    jkl: { value: undefined, enumerable: false }
                }
            );
            Object.getOwnPropertyNames(ghi)
        `, "ghi,jkl")
	})
}

func TestObjectGetterSetter(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`raise:
            Object.create({}, {
                abc: {
                    get: function(){
                        return "true";
                    },
                    writable: true
                }
            }).abc;
        `, "TypeError: toPropertyDescriptor descriptor writeSet")

		test(`raise:
            Object.create({}, {
                abc: {
                    get: function(){
                        return "true";
                    },
                    writable: false
                }
            }).abc;
        `, "TypeError: toPropertyDescriptor descriptor writeSet")

		test(`
            Object.create({}, {
                abc: {
                    get: function(){
                        return "true";
                    }
                }
            }).abc;
        `, "true")

		test(`
            Object.create({xyz:true},{abc:{get:function(){return this.xyx}}}).abc;
            Object.create({
                    xyz: true
                }, {
                abc: {
                    get: function(){
                        return this.xyz;
                    }
                }
            }).abc;
        `, true)

		test(`
            var abc = false;
            var def = Object.create({}, {
                xyz: {
                    set: function(value) {
                        abc = value;
                    }
                }
            });
            def.xyz = true;
            [ abc ];
        `, "true")

		test(`
            var abc = {};
            Object.defineProperty(abc, "def", {
                value: "xyzzy",
                configurable: true
            });
            Object.preventExtensions(abc);
            Object.defineProperty(abc, "def", {
                get: function() {
                    return 5;
                }
            });
            var def = Object.getOwnPropertyDescriptor(abc, "def");
            [ abc.def, typeof def.get, typeof def.set, typeof def.value, def.configurable, def.enumerable, typeof def.writable ];
        `, "5,function,undefined,undefined,true,false,undefined")

		test(`
            var abc = {};
            Object.defineProperty(abc, "def", {
                get: function() {
                    return 5;
                }
                configurable: true
            });
            Object.preventExtensions(abc);
            Object.defineProperty(abc, "def", {
                value: "xyzzy",
            });
            var def = Object.getOwnPropertyDescriptor(abc, "def");
            [ abc.def, typeof def.get, typeof def.set, def.value, def.configurable, def.enumerable, def.writable ];
        `, "xyzzy,undefined,undefined,xyzzy,true,false,false")

		test(`
            var abc = {};

            function _get0() {
                return 10;
            }

            function _set(value) {
                abc.def = value;
            }

            Object.defineProperty(abc, "ghi", {
                get: _get0,
                set: _set,
                configurable: true
            });

            function _get1() {
                return 20;
            }

            Object.defineProperty(abc, "ghi", {
                get: _get0
            });

            var descriptor = Object.getOwnPropertyDescriptor(abc, "ghi");
            [ typeof descriptor.set ];
        `, "function")

		test(`raise:
            var abc = [];
            Object.defineProperty(abc, "length", {
                get: function () {
                    return 2;
                }
            });
        `, `TypeError: Array.DefineOwnProperty ["function () {\n                    return 2;\n                }" <nil>] is not a value`)

		test(`
            var abc = {};

            var getter = function() {
                return 1;
            }

            Object.defineProperty(abc, "def", {
                get: getter,
                configurable: false
            });

            var jkl = undefined;
            try {
                Object.defineProperty(abc, "def", {
                    get: undefined
                });
            }
            catch (err) {
                jkl = err;
            }
            var ghi = Object.getOwnPropertyDescriptor(abc, "def");
            [ jkl instanceof TypeError, ghi.get === getter, ghi.configurable, ghi.enumerable ];
        `, "true,true,false,false")

		test(`
            var abc = {};

            var getter = function() {
                return 1;
            };

            Object.defineProperty(abc, "def", {
                get: getter
            });

            Object.defineProperty(abc, "def", {
                set: undefined
            });

            var ghi = Object.getOwnPropertyDescriptor(abc, "def");
            [ ghi.get === getter, ghi.set === undefined, ghi.configurable, ghi.enumerable ];
        `, "true,true,false,false")

		test(`
            var abc = {};

            var getter = function() {
                return 1;
            };

            Object.defineProperty(abc, "def", {
                get: getter
            });

            var jkl = undefined;
            try {
                Object.defineProperty(abc, "def", {
                    set: function() {}
                });
            }
            catch (err) {
                jkl = err;
            }

            var ghi = Object.getOwnPropertyDescriptor(abc, "def");
            [ jkl instanceof TypeError, ghi.get === getter, ghi.set, ghi.configurable, ghi.enumerable ];
        `, "true,true,,false,false")

		test(`
            var abc = {};
            var def = "xyzzy";

            Object.defineProperty(abc, "ghi", {
                get: undefined,
                set: function(value) {
                    def = value;
                },
                enumerable: true,
                configurable: true
            });

            var hasOwn = abc.hasOwnProperty("ghi");
            var descriptor = Object.getOwnPropertyDescriptor(abc, "ghi");

            [ hasOwn, typeof descriptor.get ];
        `, "true,undefined")

		test(`
            var abc = "xyzzy";
            Object.defineProperty(Array.prototype, "abc", {
                get: function () {
                    return abc;
                },
                set: function (value) {
                    abc = value;
                },
                enumerable: true,
                configurable: true
            });
            var def = [];
            def.abc = 3.14159;
            [ def.hasOwnProperty("abc"), def.abc, abc ];
        `, "false,3.14159,3.14159")
	})
}

func TestProperty(t *testing.T) {
	tt(t, func() {
		prop := property{}
		prop.writeOn()
		is(prop.writeSet(), true)

		prop.writeClear()
		is(prop.writeSet(), false)

		prop.writeOff()
		is(prop.writeSet(), true)

		prop.writeClear()
		is(prop.writeSet(), false)
	})
}
