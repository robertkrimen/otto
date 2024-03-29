package otto

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// FIXME terst, Review tests

func TestOperator(t *testing.T) {
	tt(t, func() {
		test, vm := test()

		test("xyzzy = 1")
		test("xyzzy", 1)

		if true {
			vm.Set("twoPlusTwo", func(FunctionCall) Value {
				return toValue(5)
			})
			test("twoPlusTwo( 1 )", 5)

			test("1 + twoPlusTwo( 1 )", 6)

			test("-1 + twoPlusTwo( 1 )", 4)
		}

		test("result = 4")
		test("result", 4)

		test("result += 1")
		test("result", 5)

		test("result *= 2")
		test("result", 10)

		test("result /= 2")
		test("result", 5)

		test("result = 112.51 % 3.1")
		test("result", 0.9100000000000019)

		test("result = 'Xyzzy'")
		test("result", "Xyzzy")

		test("result = 'Xyz' + 'zy'")
		test("result", "Xyzzy")

		test("result = \"Xyzzy\"")
		test("result", "Xyzzy")

		test("result = 1; result = result")
		test("result", 1)

		test(`
            var result64
            =
            64
            , result10 =
            10
        `)
		test("result64", 64)
		test("result10", 10)

		test(`
            result = 1;
            result += 1;
        `)
		test("result", 2)
	})
}

func TestFunction_(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            result = 2
            xyzzy = function() {
                result += 1
            }
            xyzzy()
            result;
        `, 3)

		test(`
            xyzzy = function() {
                return 1
            }
            result = xyzzy()
        `, 1)

		test(`
            xyzzy = function() {}
            result = xyzzy()
        `, "undefined")

		test(`
            xyzzy = function() {
                return 64
                return 1
            }
            result = xyzzy()
        `, 64)

		test(`
            result = 4
            xyzzy = function() {
                result = 2
            }
            xyzzy();
            result;
        `, 2)

		test(`
            result = 4
            xyzzy = function() {
                var result
                result = 2
            }
            xyzzy();
            result;
        `, 4)

		test(`
            xyzzy = function() {
                var result = 4
                return result
            }
            result = xyzzy()
        `, 4)

		test(`
            xyzzy = function() {
                    function test() {
                        var result = 1
                    return result
                }
                    return test() + 1
                }
            result = xyzzy() + 1
        `, 3)

		test(`
            xyzzy = function() {
                function test() {
                    var result = 1
                    return result
                }
                _xyzzy = 2
                    var result = _xyzzy + test() + 1
                    return result
            }
            result = xyzzy() + 1;
            [ result, _xyzzy ];
        `, "5,2")

		test(`
            xyzzy = function(apple) {
                return 1
            }
            result = xyzzy(1)
        `, 1)

		test(`
            xyzzy = function(apple) {
                return apple + 1
            }
            result = xyzzy(2)
        `, 3)

		test(`
            {
                result = 1
                result += 1;
            }
        `, 2)

		test(`
            var global = 1
            outer = function() {
                var global = 2
                var inner = function(){
                    return global
                }
                return inner()
            }
            result = outer()
        `, 2)

		test(`
            var apple = 1
            var banana = function() {
                return apple
            }
            var cherry = function() {
                var apple = 2
                return banana()
            }
            result = cherry()
        `, 1)

		test(`
            function xyz() {
            };
            delete xyz;
        `, false)

		test(`
            var abc = function __factorial(def){
                if (def === 1) {
                    return def;
                } else {
                    return __factorial(def-1)*def;
                }
            };
            abc(3);
        `, 6)
	})
}

func TestDoWhile(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            limit = 4;
            result = 0;
            do { 
                result = result + 1;
                limit = limit - 1;
            } while (limit);
            result;
        `, 4)

		test(`
            result = eval("do {abc=1; break; abc=2;} while (0);");
            [ result, abc ];
        `, "1,1")
	})
}

func TestContinueBreak(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            limit = 4
            result = 0
            while (limit) {
                limit = limit - 1
                if (limit) {
                }
                else {
                    break
                }
                result = result + 1
            }
            [ result, limit ];
        `, "3,0")

		test(`
            limit = 4
            result = 0
            while (limit) {
                limit = limit - 1
                if (limit) {
                    continue
                }
                else {
                    break
                }
                result = result + 1
            }
            result;
        `, 0)

		test(`
            limit = 4
            result = 0
            do {
                limit = limit - 1
                if (limit) {
                    continue
                }
                else {
                    break
                }
                result = result + 1
            } while (limit)
            result;
        `, 0)
	})
}

func TestTryCatchError(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var abc
            try {
                1()
            }
            catch (def) {
                abc = def
            }
            abc;
        `, "TypeError: 1 is not a function")
	})
}

func TestPositiveNegativeZero(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`1/0`, infinity)
		test(`1/-0`, -infinity)
		test(`
            abc = -0
            1/abc
        `, -infinity)
	})
}

func TestComparison(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            undefined = 1; undefined;
        `, "undefined")

		test("undefined == undefined", true)

		test("undefined != undefined", false)

		test("null == null", true)

		test("null != null", false)

		test("0 == 1", false)

		is(negativeZero, -0)
		is(positiveZero, 0)
		is(math.Signbit(negativeZero), true)
		is(positiveZero == negativeZero, true)

		test("1 == 1", true)

		test("'Hello, World.' == 'Goodbye, World.'", false)

		test("'Hello, World.' == true", false)

		test("'Hello, World.' == false", false)

		test("'Hello, World.' == 1", false)

		test("1 == 'Hello, World.'", false)

		is(parseNumber("-1"), -1)

		test("0+Object", "0function Object() { [native code] }")
	})
}

func TestComparisonRelational(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test("0 < 0", false)

		test("0 > 0", false)

		test("0 <= 0", true)

		test("0 >= 0", true)

		test("'   0' >= 0", true)

		test("'_   0' >= 0", false)
	})
}

func TestArguments(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            xyzzy = function() {
                return arguments[0]
            }
            result = xyzzy("xyzzy");
        `, "xyzzy")

		test(`
            xyzzy = function() {
                arguments[0] = "abcdef"
                return arguments[0]
            }
            result = xyzzy("xyzzy");
        `, "abcdef")

		test(`
            xyzzy = function(apple) {
                apple = "abcdef"
                return arguments[0]
            }
            result = xyzzy("xyzzy");
        `, "abcdef")

		test(`
            (function(){
                return arguments
            })()
        `, "[object Arguments]")

		test(`
            (function(){
                return arguments.length
            })()
        `, 0)

		test(`
            (function(){
                return arguments.length
            })(1, 2, 4, 8, 10)
        `, 5)
	})
}

func TestObjectLiteral(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            ({});
        `, "[object Object]")

		test(`
            var abc = {
                xyzzy: "Nothing happens.",
                get 1e2() {
                    return 3.14159;
                },
                get null() {
                    return true;
                },
                get "[\n]"() {
                    return "<>";
                }
            };
            [ abc["1e2"], abc.null, abc["[\n]"] ]; 
        `, "3.14159,true,<>")

		test(`
            var abc = {
                xyzzy: "Nothing happens.",
                set 1e2() {
                    this[3.14159] = 100;
                    return Math.random();
                },
                set null(def) {
                    this.def = def;
                    return Math.random();
                },
            };
            [ abc["1e2"] = Infinity, abc[3.14159], abc.null = "xyz", abc.def ];
        `, "Infinity,100,xyz,xyz")
	})
}

func TestUnaryPrefix(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var result = 0;
            [++result, result];
        `, "1,1")

		test(`
            result = 0;
            [--result, result];
        `, "-1,-1")

		test(`
            var object = { valueOf: function() { return 1; } };
            result = ++object;
            [ result, typeof result ];
        `, "2,number")

		test(`
            var object = { valueOf: function() { return 1; } };
            result = --object;
            [ result, typeof result ];
        `, "0,number")
	})
}

func TestUnaryPostfix(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var result = 0;
            result++;
            [ result++, result ];
        `, "1,2")

		test(`
            result = 0;
            result--;
            [ result--, result ];
        `, "-1,-2")

		test(`
            var object = { valueOf: function() { return 1; } };
            result = object++;
            [ result, typeof result ];
        `, "1,number")

		test(`
            var object = { valueOf: function() { return 1; } };
            result = object--
            [ result, typeof result ];
        `, "1,number")
	})
}

func TestBinaryLogicalOperation(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = true
            def = false
            ghi = false
            jkl = false
            result = abc && def || ghi && jkl
        `, false)

		test(`
            abc = true
            def = true
            ghi = false
            jkl = false
            result = abc && def || ghi && jkl
        `, true)
	})
}

func TestBinaryBitwiseOperation(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = 1 & 2;
            def = 1 & 3;
            ghi = 1 | 3;
            jkl = 1 ^ 2;
            mno = 1 ^ 3;
            [ abc, def, ghi, jkl, mno ];
        `, "0,1,3,3,2")
	})
}

func TestBinaryShiftOperation(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            high = (1 << 30) - 1 + (1 << 30)
            low = -high - 1
            abc = 23 << 1
            def = -105 >> 1
            ghi = 23 << 2
            jkl = 1 >>> 31
            mno = 1 << 64
            pqr = 1 >> 2
            stu = -2 >> 4
            vwx = low >> 1
            yz = low >>> 1
        `)
		test("abc", 46)
		test("def", -53)
		test("ghi", 92)
		test("jkl", 0)
		test("mno", 1)
		test("pqr", 0)
		test("stu", -1)
		test("vwx", -1073741824)
		test("yz", 1073741824)
	})
}

func TestParenthesizing(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = 1 + 2 * 3
            def = (1 + 2) * 3
            ghi = !(false || true)
            jkl = !false || true
        `)
		test("abc", 7)
		test("def", 9)
		test("ghi", false)
		test("jkl", true)
	})
}

func Test_instanceof(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = {} instanceof Object;
        `, true)

		test(`
            abc = "abc" instanceof Object;
        `, false)

		test(`raise:
            abc = {} instanceof "abc";
        `, "TypeError: invalid kind String for instanceof (expected object)")

		test(`raise:
            "xyzzy" instanceof Math;
        `, "TypeError: Object.hasInstance not callable")
	})
}

func TestIn(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = "prototype" in Object;
            def = "xyzzy" in Object;
            [ abc, def ];
        `, "true,false")
	})
}

func Test_new(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = new Boolean;
            def = new Boolean(1);
            [ abc, def ];
        `, "false,true")
	})
}

func TestNewFunction(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            new Function("return 11")()
        `, 11)

		test(`
            abc = 10
            new Function("abc += 1")()
            abc
        `, 11)

		test(`
            new Function("a", "b", "c", "return b + 2")(10, 11, 12)
        `, 13)

		test(`raise:
            new 1
        `, "TypeError: 1 is not a function")

		// TODO Better error reporting: new this
		test(`raise:
            new this
        `, "TypeError: [object environment] is not a function")

		test(`raise:
            new {}
        `, "TypeError: [object Object] is not a function")
	})
}

func TestNewPrototype(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            abc = { 'xyzzy': 'Nothing happens.' }
            function Xyzzy(){}
            Xyzzy.prototype = abc;
            (new Xyzzy()).xyzzy
        `, "Nothing happens.")
	})
}

func TestBlock(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var abc=0;
            var ghi;
            def: {
                do {
                    abc++;
                    if (!(abc < 10)) {
                        break def;
                        ghi = "ghi";
                    }
                } while (true);
            }
            [ abc,ghi ];
        `, "10,")
	})
}

func Test_toString(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            [undefined+""]
        `, "undefined")
	})
}

func TestEvaluationOrder(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var abc = 0;
            abc < (abc = 1) === true;
        `, true)
	})
}

func TestClone(t *testing.T) {
	tt(t, func() {
		vm1 := New()
		_, err := vm1.Run(`
            var abc = 1;
        `)
		require.NoError(t, err)

		vm2 := vm1.clone()
		_, err = vm1.Run(`
            abc += 2;
        `)
		require.NoError(t, err)
		_, err = vm2.Run(`
            abc += 4;
        `)
		require.NoError(t, err)

		is(vm1.getValue("abc"), 3)
		is(vm2.getValue("abc"), 5)
	})
}

func Test_debugger(t *testing.T) {
	tt(t, func() {
		called := false

		vm := New()
		vm.SetDebuggerHandler(func(o *Otto) {
			is(o, vm)
			called = true
		})

		_, err := vm.Run(`debugger`)
		require.NoError(t, err)
		is(called, true)
	})

	tt(t, func() {
		called := false

		vm := New()
		vm.SetDebuggerHandler(func(o *Otto) {
			is(o, vm)
			called = true
		})

		_, err := vm.Run(`null`)
		require.NoError(t, err)
		is(called, false)
	})

	tt(t, func() {
		vm := New()

		_, err := vm.Run(`debugger`)
		require.NoError(t, err)
	})
}

func Test_random(t *testing.T) {
	tt(t, func() {
		vm := New()
		vm.SetRandomSource(func() float64 { return 1 })

		r, err := vm.Run(`Math.random()`)
		require.NoError(t, err)
		f, err := r.ToFloat()
		require.NoError(t, err)
		is(f, 1)
	})

	tt(t, func() {
		vm := New()

		r1, err := vm.Run(`Math.random()`)
		require.NoError(t, err)
		f1, err := r1.ToFloat()
		require.NoError(t, err)

		r2, err := vm.Run(`Math.random()`)
		require.NoError(t, err)
		f2, err := r2.ToFloat()
		require.NoError(t, err)

		is(f1 == f2, false)
	})
}

func Test_stringArray(t *testing.T) {
	getStrings := func() []string {
		return []string{"these", "are", "strings"}
	}
	concatStrings := func(a []string) string {
		if len(a) == 0 {
			return ""
		}
		r := a[0]
		for i := 1; i < len(a); i++ {
			r += " "
			r += a[i]
		}
		return r
	}
	tt(t, func() {
		vm := New()
		err := vm.Set("getStrings", getStrings)
		require.NoError(t, err)
		err = vm.Set("concatStrings", concatStrings)
		require.NoError(t, err)
		r1, err := vm.Run(`var a = getStrings(); concatStrings(a)`)
		require.NoError(t, err)
		is(r1, "these are strings")
	})
}

type goByteArrayWithMethodsTest [8]byte

func (g goByteArrayWithMethodsTest) S() string    { return string(g[:]) }
func (g goByteArrayWithMethodsTest) F(i int) byte { return g[i] }

func Test_goByteArrayWithMethods_typeof_S(t *testing.T) {
	a := goByteArrayWithMethodsTest{97, 98, 99, 100, 101, 102, 103, 104}

	tt(t, func() {
		test, vm := test()
		vm.Set("a", a)
		is(test("typeof a.S").export(), "function")
	})
}

func Test_goByteArrayWithMethods_S(t *testing.T) {
	a := goByteArrayWithMethodsTest{97, 98, 99, 100, 101, 102, 103, 104}

	tt(t, func() {
		test, vm := test()
		vm.Set("a", a)
		is(test("a.S()").export(), "abcdefgh")
	})
}

func Test_goByteArrayWithMethods_F0(t *testing.T) {
	a := goByteArrayWithMethodsTest{97, 98, 99, 100, 101, 102, 103, 104}

	tt(t, func() {
		test, vm := test()
		vm.Set("a", a)
		is(test("a.F(0)").export(), 97)
	})
}

func Test_goByteArrayWithMethods_F1(t *testing.T) {
	a := goByteArrayWithMethodsTest{97, 98, 99, 100, 101, 102, 103, 104}

	tt(t, func() {
		test, vm := test()
		vm.Set("a", a)
		is(test("a.F(1)").export(), 98)
	})
}
