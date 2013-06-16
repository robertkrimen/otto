package otto

import (
	. "./terst"
	"testing"
)

func TestNumber(t *testing.T) {
	Terst(t)
	test := runTest()
	test(`
        var abc = Object.getOwnPropertyDescriptor(Number, "prototype");
        [   [ typeof Number.prototype ],
            [ abc.writable, abc.enumerable, abc.configurable ] ];
    `, "object,false,false,false")
}

func TestNumber_toString(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
        new Number(451).toString();
    `, "451")

	test(`
        new Number(451).toString(10);
    `, "451")

	test(`
        new Number(451).toString(8);
    `, "703")

	test(`raise:
        new Number(451).toString(1);
    `, "RangeError: RangeError: toString() radix must be between 2 and 36")

	test(`raise:
        new Number(451).toString(Infinity);
    `, "RangeError: RangeError: toString() radix must be between 2 and 36")

	test(`
        new Number(NaN).toString()
    `, "NaN")

	test(`
        new Number(Infinity).toString()
    `, "Infinity")

	test(`
        new Number(Infinity).toString(16)
    `, "Infinity")
}

func TestNumber_toFixed(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`new Number(451).toFixed(2)`, "451.00")
	test(`12345.6789.toFixed()`, "12346")
	test(`12345.6789.toFixed(1)`, "12345.7")
	test(`12345.6789.toFixed(6)`, "12345.678900")
	test(`(1.23e-20).toFixed(2)`, "0.00")
	test(`2.34.toFixed(1)`, "2.3")
	test(`-2.34.toFixed(1)`, "-2.3")
	test(`(-2.34).toFixed(1)`, "-2.3")
}

func TestNumber_toExponential(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`new Number(451).toExponential(2)`, "4.51e+02")
	test(`77.1234.toExponential()`, "7.71234e+01")
	test(`77.1234.toExponential(4)`, "7.7123e+01")
	test(`77.1234.toExponential(2)`, "7.71e+01")
	test(`77 .toExponential()`, "7.7e+01")
}

func TestNumber_toPrecision(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`new Number(451).toPrecision()`, "451")
	test(`new Number(451).toPrecision(1)`, "5e+02")
	test(`5.123456.toPrecision()`, "5.123456")
	test(`5.123456.toPrecision(5)`, "5.1235")
	test(`5.123456.toPrecision(2)`, "5.1")
	test(`5.123456.toPrecision(1)`, "5")
}
