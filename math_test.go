package otto

import (
	"math"
	"testing"
)

var (
	naN      = math.NaN()
	infinity = math.Inf(1)
)

func TestMath_toString(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.toString()`, "[object Math]")
	})
}

func TestMath_abs(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.abs(NaN)`, naN)
		test(`Math.abs(2)`, 2)
		test(`Math.abs(-2)`, 2)
		test(`Math.abs(-Infinity)`, infinity)

		test(`Math.acos(0.5)`, 1.0471975511965976)

		test(`Math.abs('-1')`, 1)
		test(`Math.abs(-2)`, 2)
		test(`Math.abs(null)`, 0)
		test(`Math.abs("string")`, naN)
		test(`Math.abs()`, naN)
	})
}

func TestMath_acos(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.acos(NaN)`, naN)
		test(`Math.acos(2)`, naN)
		test(`Math.acos(-2)`, naN)
		test(`1/Math.acos(1)`, infinity)

		test(`Math.acos(0.5)`, 1.0471975511965976)
	})
}

func TestMath_acosh(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.acosh(-1)`, naN)
		test(`Math.acosh(0)`, naN)
		test(`Math.acosh(0.999999999999)`, naN)
		test(`1/Math.acosh(1)`, infinity)
		test(`Math.acosh(Infinity)`, infinity)
		test(`Math.acosh(2)`, 1.3169578969248166)
		test(`Math.acosh(2.5)`, 1.566799236972411)
	})
}

func TestMath_asin(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.asin(NaN)`, naN)
		test(`Math.asin(2)`, naN)
		test(`Math.asin(-2)`, naN)
		test(`1/Math.asin(0)`, infinity)
		test(`1/Math.asin(-0)`, -infinity)

		test(`Math.asin(0.5)`, 0.5235987755982989)
	})
}

func TestMath_asinh(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.asinh(-1)`, -0.881373587019543)
		test(`Math.asinh(1)`, 0.881373587019543)
		test(`Math.asinh(-0)`, -0)
		test(`Math.asinh(0)`, 0)
		test(`Math.asinh(-Infinity)`, -infinity)
		test(`Math.asinh(Infinity)`, infinity)
		test(`Math.asinh(2)`, 1.4436354751788103)
	})
}

func TestMath_atan(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.atan(NaN)`, naN)
		test(`1/Math.atan(0)`, infinity)
		test(`1/Math.atan(-0)`, -infinity)
		test(`Math.atan(Infinity)`, 1.5707963267948966)
		test(`Math.atan(-Infinity)`, -1.5707963267948966)

		// freebsd/386 1.03 => 0.4636476090008061
		// darwin 1.03 => 0.46364760900080604
		test(`Math.atan(0.5).toPrecision(10)`, "0.463647609")
	})
}

func TestMath_atan2(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.atan2()`, naN)
		test(`Math.atan2(NaN)`, naN)
		test(`Math.atan2(0, NaN)`, naN)

		test(`Math.atan2(1, 0)`, 1.5707963267948966)
		test(`Math.atan2(1, -0)`, 1.5707963267948966)

		test(`1/Math.atan2(0, 1)`, infinity)
		test(`1/Math.atan2(0, 0)`, infinity)
		test(`Math.atan2(0, -0)`, 3.141592653589793)
		test(`Math.atan2(0, -1)`, 3.141592653589793)

		test(`1/Math.atan2(-0, 1)`, -infinity)
		test(`1/Math.atan2(-0, 0)`, -infinity)
		test(`Math.atan2(-0, -0)`, -3.141592653589793)
		test(`Math.atan2(-0, -1)`, -3.141592653589793)

		test(`Math.atan2(-1, 0)`, -1.5707963267948966)
		test(`Math.atan2(-1, -0)`, -1.5707963267948966)

		test(`1/Math.atan2(1, Infinity)`, infinity)
		test(`Math.atan2(1, -Infinity)`, 3.141592653589793)
		test(`1/Math.atan2(-1, Infinity)`, -infinity)
		test(`Math.atan2(-1, -Infinity)`, -3.141592653589793)

		test(`Math.atan2(Infinity, 1)`, 1.5707963267948966)
		test(`Math.atan2(-Infinity, 1)`, -1.5707963267948966)

		test(`Math.atan2(Infinity, Infinity)`, 0.7853981633974483)
		test(`Math.atan2(Infinity, -Infinity)`, 2.356194490192345)
		test(`Math.atan2(-Infinity, Infinity)`, -0.7853981633974483)
		test(`Math.atan2(-Infinity, -Infinity)`, -2.356194490192345)
	})
}

func TestMath_atanh(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.atanh(-2)`, naN)
		test(`Math.atanh(2)`, naN)
		test(`Math.atanh(-1)`, -infinity)
		test(`Math.atanh(1)`, infinity)
		test(`Math.atanh(0)`, 0)
		test(`Math.atanh(0.5)`, 0.5493061443340548)
	})
}

func TestMath_cbrt(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.cbrt(NaN)`, naN)
		test(`Math.cbrt(-1)`, -1)
		test(`Math.cbrt(1)`, 1)
		test(`Math.cbrt(-0)`, -0)
		test(`Math.cbrt(0)`, 0)
		test(`Math.cbrt(-Infinity)`, -infinity)
		test(`Math.cbrt(Infinity)`, infinity)
		test(`Math.cbrt(null)`, 0)
		test(`Math.cbrt(2)`, 1.2599210498948732)
	})
}

func TestMath_ceil(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.ceil(NaN)`, naN)
		test(`Math.ceil(+0)`, 0)
		test(`1/Math.ceil(-0)`, -infinity)
		test(`Math.ceil(Infinity)`, infinity)
		test(`Math.ceil(-Infinity)`, -infinity)
		test(`1/Math.ceil(-0.5)`, -infinity)

		test(`Math.ceil(-11)`, -11)
		test(`Math.ceil(-0.5)`, 0)
		test(`Math.ceil(1.5)`, 2)
	})
}

func TestMath_cos(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.cos(NaN)`, naN)
		test(`Math.cos(+0)`, 1)
		test(`Math.cos(-0)`, 1)
		test(`Math.cos(Infinity)`, naN)
		test(`Math.cos(-Infinity)`, naN)

		test(`Math.cos(0.5)`, 0.8775825618903728)
	})
}

func TestMath_cosh(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.cosh(0)`, 1)
		test(`Math.cosh(1)`, 1.5430806348152437)
		test(`Math.cosh(-1)`, 1.5430806348152437)
	})
}

func TestMath_exp(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.exp(NaN)`, naN)
		test(`Math.exp(+0)`, 1)
		test(`Math.exp(-0)`, 1)
		test(`Math.exp(Infinity)`, infinity)
		test(`Math.exp(-Infinity)`, 0)
	})
}

func TestMath_expm1(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.expm1(0)`, 0)
		test(`Math.expm1(1)`, 1.718281828459045)
		test(`Math.expm1(-1)`, -0.6321205588285577)
		test(`Math.expm1(2)`, 6.38905609893065)
		test(`Math.expm1("foo")`, naN)
	})
}

func TestMath_floor(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.floor(NaN)`, naN)
		test(`Math.floor(+0)`, 0)
		test(`1/Math.floor(-0)`, -infinity)
		test(`Math.floor(Infinity)`, infinity)
		test(`Math.floor(-Infinity)`, -infinity)

		test(`Math.floor(-11)`, -11)
		test(`Math.floor(-0.5)`, -1)
		test(`Math.floor(1.5)`, 1)
	})
}

func TestMath_log(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.log(NaN)`, naN)
		test(`Math.log(-1)`, naN)
		test(`Math.log(+0)`, -infinity)
		test(`Math.log(-0)`, -infinity)
		test(`1/Math.log(1)`, infinity)
		test(`Math.log(Infinity)`, infinity)

		test(`Math.log(0.5)`, -0.6931471805599453)
	})
}

func TestMath_log10(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.log10(100000)`, 5)
		test(`Math.log10(-2)`, naN)
		test(`Math.log10(2)`, 0.3010299956639812)
		test(`Math.log10(1)`, 0)
		test(`Math.log10(-0)`, -infinity)
		test(`Math.log10(0)`, -infinity)
		test(`Math.log10(Infinity)`, infinity)
	})
}

func TestMath_log1p(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.log1p(-2)`, naN)
		test(`Math.log1p(-1)`, -infinity)
		test(`Math.log1p(1)`, 0.6931471805599453)
		test(`Math.log1p(-0)`, -0)
		test(`Math.log1p(0)`, 0)
		test(`Math.log1p(Infinity)`, infinity)
	})
}

func TestMath_log2(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.log2(-2)`, naN)
		test(`Math.log2(-0)`, -infinity)
		test(`Math.log2(0)`, -infinity)
		test(`Math.log2(1)`, 0)
		test(`Math.log2(2)`, 1)
		test(`Math.log2(5)`, 2.321928094887362)
		test(`Math.log2(1024)`, 10)
		test(`Math.log2(Infinity)`, infinity)
	})
}

func TestMath_max(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.max(-11, -1, 0, 1, 2, 3, 11)`, 11)
	})
}

func TestMath_min(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.min(-11, -1, 0, 1, 2, 3, 11)`, -11)
	})
}

func TestMath_pow(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.pow(0, NaN)`, naN)
		test(`Math.pow(0, 0)`, 1)
		test(`Math.pow(NaN, 0)`, 1)
		test(`Math.pow(0, -0)`, 1)
		test(`Math.pow(NaN, -0)`, 1)
		test(`Math.pow(NaN, 1)`, naN)
		test(`Math.pow(2, Infinity)`, infinity)
		test(`1/Math.pow(2, -Infinity)`, infinity)
		test(`Math.pow(1, Infinity)`, naN)
		test(`Math.pow(1, -Infinity)`, naN)
		test(`1/Math.pow(0.1, Infinity)`, infinity)
		test(`Math.pow(0.1, -Infinity)`, infinity)
		test(`Math.pow(Infinity, 1)`, infinity)
		test(`1/Math.pow(Infinity, -1)`, infinity)
		test(`Math.pow(-Infinity, 1)`, -infinity)
		test(`Math.pow(-Infinity, 2)`, infinity)
		test(`1/Math.pow(-Infinity, -1)`, -infinity)
		test(`1/Math.pow(-Infinity, -2)`, infinity)
		test(`1/Math.pow(0, 1)`, infinity)
		test(`Math.pow(0, -1)`, infinity)
		test(`1/Math.pow(-0, 1)`, -infinity)
		test(`1/Math.pow(-0, 2)`, infinity)
		test(`Math.pow(-0, -1)`, -infinity)
		test(`Math.pow(-0, -2)`, infinity)
		test(`Math.pow(-1, 0.1)`, naN)

		test(`
            [ Math.pow(-1, +Infinity), Math.pow(1, Infinity) ];
        `, "NaN,NaN")
	})
}

func TestMath_round(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.round(NaN)`, naN)
		test(`1/Math.round(0)`, infinity)
		test(`1/Math.round(-0)`, -infinity)
		test(`Math.round(Infinity)`, infinity)
		test(`Math.round(-Infinity)`, -infinity)
		test(`1/Math.round(0.1)`, infinity)
		test(`1/Math.round(-0.1)`, -infinity)

		test(`Math.round(3.5)`, 4)
		test(`Math.round(-3.5)`, -3)
	})
}

func TestMath_sin(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.sin(NaN)`, naN)
		test(`1/Math.sin(+0)`, infinity)
		test(`1/Math.sin(-0)`, -infinity)
		test(`Math.sin(Infinity)`, naN)
		test(`Math.sin(-Infinity)`, naN)

		test(`Math.sin(0.5)`, 0.479425538604203)
	})
}

func TestMath_sinh(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.sinh(-Infinity)`, -infinity)
		test(`Math.sinh(Infinity)`, infinity)
		test(`Math.sinh(-0)`, -0)
		test(`Math.sinh(0)`, 0)
		test(`Math.sinh(-1)`, -1.1752011936438014)
		test(`Math.sinh(1)`, 1.1752011936438014)
		test(`Math.sinh(2)`, 3.626860407847019)
	})
}

func TestMath_sqrt(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.sqrt(NaN)`, naN)
		test(`Math.sqrt(-1)`, naN)
		test(`1/Math.sqrt(+0)`, infinity)
		test(`1/Math.sqrt(-0)`, -infinity)
		test(`Math.sqrt(Infinity)`, infinity)

		test(`Math.sqrt(2)`, 1.4142135623730951)
	})
}

func TestMath_tan(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.tan(NaN)`, naN)
		test(`1/Math.tan(+0)`, infinity)
		test(`1/Math.tan(-0)`, -infinity)
		test(`Math.tan(Infinity)`, naN)
		test(`Math.tan(-Infinity)`, naN)

		test(`Math.tan(0.5)`, 0.5463024898437905)
	})
}

func TestMath_tanh(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.tanh(Infinity)`, 1)
		test(`Math.tanh(-Infinity)`, -1)
		test(`Math.tanh(-1)`, -0.7615941559557649)
		test(`Math.tanh(1)`, 0.7615941559557649)
		test(`Math.tanh(-0)`, -0)
		test(`Math.tanh(0)`, 0)
	})
}

func TestMath_trunc(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`Math.trunc(-Infinity)`, -infinity)
		test(`Math.trunc(Infinity)`, infinity)
		test(`Math.trunc(-0.123)`, -0)
		test(`Math.trunc(0.123)`, 0)
		test(`Math.trunc(-0)`, -0)
		test(`Math.trunc(0)`, 0)
		test(`Math.trunc("-1.123")`, -1)
		test(`Math.trunc(13.37)`, 13)
		test(`Math.trunc(42.84)`, 42)
	})
}
