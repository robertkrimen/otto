package otto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// this is its own file because the tests in it rely on the line numbers of
// some of the functions defined here. putting it in with the rest of the
// tests would probably be annoying.

func TestFunction_stack(t *testing.T) {
	tt(t, func() {
		vm := New()

		s, err := vm.Compile("fake.js", `function X(fn1, fn2, fn3) { fn1(fn2, fn3); }`)
		require.NoError(t, err)
		_, err = vm.Run(s)
		require.NoError(t, err)

		expected := []frame{
			{native: true, nativeFile: "function_stack_test.go", nativeLine: 36, offset: 0, callee: "github.com/robertkrimen/otto.TestFunction_stack.func1.2"},
			{native: true, nativeFile: "function_stack_test.go", nativeLine: 29, offset: 0, callee: "github.com/robertkrimen/otto.TestFunction_stack.func1.1"},
			{native: false, nativeFile: "", nativeLine: 0, offset: 29, callee: "X", file: s.program.file},
			{native: false, nativeFile: "", nativeLine: 0, offset: 29, callee: "X", file: s.program.file},
		}

		err = vm.Set("A", func(c FunctionCall) Value {
			_, err := c.Argument(0).Call(UndefinedValue())
			require.NoError(t, err)
			return UndefinedValue()
		})
		require.NoError(t, err)

		err = vm.Set("B", func(c FunctionCall) Value {
			depth := 0
			for s := c.Otto.runtime.scope; s != nil; s = s.outer {
				// these properties are tested explicitly so that we don't test `.fn`,
				// which will differ from run to run
				is(s.frame.native, expected[depth].native)
				is(s.frame.nativeFile, expected[depth].nativeFile)
				is(s.frame.nativeLine, expected[depth].nativeLine)
				is(s.frame.offset, expected[depth].offset)
				is(s.frame.callee, expected[depth].callee)
				is(s.frame.file, expected[depth].file)
				depth++
			}

			is(depth, 4)

			return UndefinedValue()
		})
		require.NoError(t, err)

		x, err := vm.Get("X")
		require.NoError(t, err)
		a, err := vm.Get("A")
		require.NoError(t, err)
		b, err := vm.Get("B")
		require.NoError(t, err)

		_, err = x.Call(UndefinedValue(), x, a, b)
		require.NoError(t, err)
	})
}
