package otto

import (
	. "./terst"
	"testing"
)

func TestObject_(t *testing.T) {
	Terst(t)

	object := newObject(nil, "")
	IsTrue(object != nil)

	object.put("xyzzy", toValue("Nothing happens."), true)
	Is(object.get("xyzzy"), "Nothing happens.")
}

func TestStringObject(t *testing.T) {
	Terst(t)

	object := New().runtime.newStringObject(toValue("xyzzy"))
	Is(object.get("1"), "y")
	Is(object.get("10"), "undefined")
	Is(object.get("2"), "z")
}

func TestObject_getPrototypeOf(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`
        abc = {};
        def = Object.getPrototypeOf(abc);
        ghi = Object.getPrototypeOf(def);
        [abc,def,ghi,ghi+""];
    `, "[object Object],[object Object],,null")
}

func TestObject_new(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`
        [ new Object("abc"), new Object(2+2) ];
    `, "abc,4")
}

func TestObject_create(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: Object.create()`, "TypeError")
	test(`
		var abc = Object.create(null)
		var def = Object.create({x: 10, y: 20})
		var ghi = Object.create(Object.prototype)

		var jkl = Object.create({x: 10, y: 20}, {
			z: {
				value: 30,
				writable: true
			},
			sum: {
				get: function() {
					return this.x + this.y + this.z
				},
			}
		})

		var xyz = [ abc.prototype, def.x, def.y, ghi, jkl.x, jkl.y, jkl.z, jkl.sum ]
		xyz
	`, ",10,20,[object Object],10,20,30,60")
}

func TestObject_keys(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`Object.keys({x:1,y:2})`, "x,y")
	test(`
		function o() { this.a = 1; this.b = 2 }
		Object.keys(new o())
	`, "a,b")
	test(`
		function p() { this.c = 2 }
		p.prototype = new o()
		Object.keys(new p())
	`, "c")
	test(`
		var obj = Object.create({a: 10, b: 20}, {
			x: { value: 30, enumerable: true },
			y: { value: 40, enumerable: false }
		})
		Object.keys(obj)
	`, "x")
}

func TestObject_getOwnPropertyNames(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`Object.getOwnPropertyNames({a:1})`, "a")
	test(`
		var obj = Object.create({a: 10, b: 20}, {
			x: { value: 30, enumerable: true },
			y: { value: 40, enumerable: false}
		})
		Object.getOwnPropertyNames(obj)
	`, "x,y")
}

func TestObject_toLocaleString(t *testing.T) {
	Terst(t)

	test := runTest()
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
}

func TestObject_isExtensible(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise:
        Object.isExtensible();
    `, "TypeError")
	test(`raise:
        Object.isExtensible({});
    `, "true")

	test(`Object.isExtensible.length`, "1")
	test(`Object.isExtensible.prototype`, "undefined")
}

func TestObject_preventExtensions(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise:
        Object.preventExtensions()
    `, "TypeError")

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

	test(`Object.preventExtensions.length`, "1")
	test(`Object.preventExtensions.prototype`, "undefined")
}

func TestObject_isSealed(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: Object.isSealed()`, "TypeError")
	test(`Object.isSealed(Object.preventExtensions({a:1}))`, "false")
	test(`Object.isSealed({})`, "false")

	test(`Object.isSealed.length`, "1")
	test(`Object.isSealed.prototype`, "undefined")
}

func TestObject_seal(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: Object.seal()`, "TypeError")
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

	test(`Object.seal.length`, "1")
	test(`Object.seal.prototype`, "undefined")
}

func TestObject_isFrozen(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: Object.isFrozen()`, "TypeError")
	test(`Object.isFrozen(Object.preventExtensions({a:1}))`, "false")
	test(`Object.isFrozen({})`, "false")

	test(`Object.isFrozen.length`, "1")
	test(`Object.isFrozen.prototype`, "undefined")
}

func TestObject_freeze(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: Object.freeze()`, "TypeError")
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

	test(`Object.freeze.length`, "1")
	test(`Object.freeze.prototype`, "undefined")
}

func TestObject_GetterSetter(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: Object.create({},{a:{get:function(){return "true"},writable:true}}).a`, "TypeError")
	test(`raise: Object.create({},{a:{get:function(){return "true"},writable:false}}).a`, "TypeError")
	test(`Object.create({},{a:{get:function(){return "true"}}}).a`, "true")
	test(`Object.create({x:true},{a:{get:function(){return this.x}}}).a`, "true")
	test(`
		var _val = false
		var o = Object.create({}, {
			val: {
				set: function(v) {
					_val = v
				},
			}
		})
	o.val = true
	_val
	`, "true")

	test(`
		var object = {
			_x: 10, _y: 20,
			get x() { return this._x },
			set x(v) { this._x = v },
			get y() { return this._y },
			set y(v) { this._y = v }
		};
		[object.x, object.y, (object.x = 20)+(object.y = 10), object.x, object.y]
	`, "10,20,30,20,10")
}

func TestObject_defineGetter(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: this.__defineGetter__("a", "x")`, "TypeError")
	test(`
		var object = {}
		object.__defineGetter__("a", function() { return true })
		object.a
	`, "true")
	test(`
		var object = {x: 10, y: 20, z: false};
		object.__defineGetter__("xy", function() { return this.x + this.y });
		object.__defineGetter__("z", function() { return true });
		[object.xy, object.z];
	`, "30,true")
}

func TestObject_lookupGetter(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`
		var object = {};
		var x = object.__lookupGetter__("a") === undefined;
		object.__defineGetter__("a", function() { return true });
		[x, object.a, object.__lookupGetter__("a") !== undefined];
	`, "true,true,true")
}

func TestObject_defineSetter(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: this.__defineSetter__("a", "x")`, "TypeError")
	test(`
		var value = false
		var object = {}
		object.__defineSetter__("a", function(v) { value = v })
		object.a = true
	`, "true")
	test(`
		var object = {x: 20, y: 10, z: 30};
		object.__defineSetter__("xyz", function(v) {
			if (Array.isArray(v)) {
				this.x = v[0]
				this.y = v[1]
				this.z = v[2]
				return
			}
			throw new TypeError()
		})
		object.xyz = [30, 20, 10];
		var caught = false;
		try {
			object.xyz = "abc"
		} catch (e) {
			caught = e instanceof TypeError
		}
		[object.x, object.y, object.z, caught];
	`, "30,20,10,true")
}

func TestObject_lookupSetter(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`
		var value = false;
		var object = {};
		var x = object.__lookupSetter__("a") === undefined;
		object.__defineSetter__("a", function(v) { value = v });
		object.a = true;
		[x, value, object.__lookupSetter__("a") !== undefined];
	`, "true,true,true")
}
