package otto

import "testing"

func TestMathNaNArgumentCoercion(t *testing.T) {
	tt(t, func() {
		test, _ := test()
		for _, method := range []string{"max", "min", "atan2"} {
			test(`(function () {
				var converted = false;
				Math.`+method+`(NaN, { valueOf: function () { converted = true; return 1; } });
				return converted;
			}())`, true)
		}
		for _, method := range []string{"max", "min"} {
			test(`(function () {
				var converted = false;
				Math.`+method+`(1, NaN, { valueOf: function () { converted = true; return 2; } });
				return converted;
			}())`, true)
		}
		test(`Math.max(NaN, Infinity)`, naN)
		test(`Math.min(NaN, -Infinity)`, naN)
	})
}
