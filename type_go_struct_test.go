package otto

import (
	"testing"
)

func TestGoStructEmbeddedFields(t *testing.T) {
	type A struct {
		A1 string `json:"a1"`
		A2 string `json:"a2"`
		A3 string `json:"a3"`
	}

	type B struct {
		A
		B1 string `json:"b1"`
	}

	tt(t, func() {
		test, vm := test()

		var b B

		b.A1 = "a1"
		b.A2 = "a2"
		b.A3 = "a3"
		b.B1 = "b1"

		vm.Set("v", B{A{"a1", "a2", "a3"}, "b1"})

		test(`[v.a1,v.a2,v.a3,v.b1]`, "a1,a2,a3,b1")
	})
}

func TestGoStructNilBoolPointerField(t *testing.T) {
	type S struct {
		A int         `json:"a"`
		B *bool       `json:"b"`
		C interface{} `json:"c"`
	}

	tt(t, func() {
		test, vm := test()
		vm.Set("s", S{A: 1, B: nil, C: nil})
		test(`'a' in s`, true)
		test(`typeof s.a`, "number")
		test(`'b' in s`, true)
		test(`typeof s.b`, "undefined")
		test(`'c' in s`, true)
		test(`typeof s.c`, "undefined")
	})
}

func TestGoStructError(t *testing.T) {
	type S1 struct {
		A string `json:"a"`
		B string `json:"b"`
	}

	type S2 struct {
		A []S1 `json:"a"`
		B S1   `json:"b"`
	}

	type S3 struct {
		A []S2 `json:"a"`
		B S2   `json:"b"`
	}

	tt(t, func() {
		test, vm := test()
		vm.Set("fn", func(s *S3) string { return "cool" })
		test(
			`(function() { try { fn({a:[{a:[{c:"x"}]}]}) } catch (ex) { return ex } })()`,
			`TypeError: can't convert to *otto.S3: couldn't convert property "a" of otto.S3: couldn't convert element 0 of []otto.S2: couldn't convert property "a" of otto.S2: couldn't convert element 0 of []otto.S1: can't convert property "c" of otto.S1: field does not exist`,
		)
	})
}
