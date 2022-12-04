package otto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// this is its own file because the tests in it rely on the line numbers of
// some of the functions defined here. putting it in with the rest of the
// tests would probably be annoying.

func TestErrorContextNative(t *testing.T) {
	tt(t, func() {
		vm := New()

		err := vm.Set("N", func(c FunctionCall) Value {
			v, err := c.Argument(0).Call(NullValue())
			if err != nil {
				panic(err)
			}
			return v
		})
		require.NoError(t, err)

		s, err := vm.Compile("test.js", `
			function F() { throw new Error('wow'); }
			function G() { return N(F); }
		`)
		require.NoError(t, err)

		_, err = vm.Run(s)
		require.NoError(t, err)

		f1, err := vm.Get("G")
		require.NoError(t, err)
		_, err = f1.Call(NullValue())
		require.Error(t, err)
		err1 := asError(t, err)
		is(err1.message, "wow")
		is(len(err1.trace), 3)
		is(err1.trace[0].location(), "F (test.js:2:29)")
		is(err1.trace[1].location(), "github.com/robertkrimen/otto.TestErrorContextNative.func1.1 (error_native_test.go:17)")
		is(err1.trace[2].location(), "G (test.js:3:26)")
	})
}
