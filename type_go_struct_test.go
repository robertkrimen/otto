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

		vm.Set("v", B{A{"a1", "a2", "a3"}, "b1"})

		test(`[v.a1,v.a2,v.a3,v.b1]`, "a1,a2,a3,b1")
	})
}

func TestGoStructNilBoolPointerField(t *testing.T) {
	type S struct {
		C interface{} `json:"c"`
		B *bool       `json:"b"`
		A int         `json:"a"`
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

func TestGoStructMapPointerField(t *testing.T) {
	type S struct {
		A *map[string]interface{} `json:"a"`
	}

	tt(t, func() {
		test, vm := test()
		data := map[string]interface{}{"one": 1}
		vm.Set("s", S{A: &data})

		test(`'a' in s`, true)
		test(`s.a.one`, 1)
		test(`s.a.one = "updated"`, "updated")
		is(data["one"], "updated")
		test(`s.a.two = "added"`, "added")
		is(data["two"], "added")
		test(`delete s.a.one`, true)
		_, exists := data["one"]
		is(exists, false)
		test(`Object.keys(s.a).join(",")`, "two")
	})
}

func TestGoStructNilMapPointerField(t *testing.T) {
	type S struct {
		Pointer *map[string]interface{} `json:"pointer"`
		NilMap  *map[string]interface{} `json:"nilMap"`
		Map     map[string]interface{}  `json:"map"`
	}

	tt(t, func() {
		test, vm := test()
		var data map[string]interface{}
		vm.Set("s", S{NilMap: &data})

		test(`'pointer' in s`, true)
		test(`typeof s.pointer`, "undefined")
		test(`typeof s.nilMap`, "object")
		test(`s.nilMap.missing === undefined`, true)
		test(`Object.keys(s.nilMap).length`, 0)
		test(`typeof s.map`, "object")
		test(`s.map.missing === undefined`, true)
		test(`Object.keys(s.map).length`, 0)
		is(data == nil, true)

		data = map[string]interface{}{"initialized": true}
		test(`s.nilMap.initialized`, true)
		test(`Object.keys(s.map).length`, 0)
	})
}

func TestGoStructError(t *testing.T) {
	type S1 struct {
		A string `json:"a"`
		B string `json:"b"`
	}

	type S2 struct {
		B S1   `json:"b"`
		A []S1 `json:"a"`
	}

	type S3 struct {
		B S2   `json:"b"`
		A []S2 `json:"a"`
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
