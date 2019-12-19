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
