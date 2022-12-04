package otto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNativeStackFrames(t *testing.T) {
	tt(t, func() {
		vm := New()

		s, err := vm.Compile("input.js", `
			function A() { ext1(); }
			function B() { ext2(); }
			A();
		`)
		require.NoError(t, err)

		err = vm.Set("ext1", func(c FunctionCall) Value {
			if _, err := c.Otto.Eval("B()"); err != nil {
				panic(err)
			}

			return UndefinedValue()
		})
		require.NoError(t, err)

		err = vm.Set("ext2", func(c FunctionCall) Value {
			{
				// no limit, include innermost native frames
				ctx := c.Otto.ContextSkip(-1, false)

				is(ctx.Stacktrace, []string{
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.2 (native_stack_test.go:29)",
					"B (input.js:3:19)",
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.1 (native_stack_test.go:20)",
					"A (input.js:2:19)", "input.js:4:4",
				})

				is(ctx.Callee, "github.com/robertkrimen/otto.TestNativeStackFrames.func1.2")
				is(ctx.Filename, "native_stack_test.go")
				is(ctx.Line, 29)
				is(ctx.Column, 0)
			}

			{
				// no limit, skip innermost native frames
				ctx := c.Otto.ContextSkip(-1, true)

				is(ctx.Stacktrace, []string{
					"B (input.js:3:19)",
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.1 (native_stack_test.go:20)",
					"A (input.js:2:19)", "input.js:4:4",
				})

				is(ctx.Callee, "B")
				is(ctx.Filename, "input.js")
				is(ctx.Line, 3)
				is(ctx.Column, 19)
			}

			if _, err := c.Otto.Eval("ext3()"); err != nil {
				panic(err)
			}

			return UndefinedValue()
		})
		require.NoError(t, err)

		err = vm.Set("ext3", func(c FunctionCall) Value {
			{
				// no limit, include innermost native frames
				ctx := c.Otto.ContextSkip(-1, false)

				is(ctx.Stacktrace, []string{
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.3 (native_stack_test.go:71)",
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.2 (native_stack_test.go:29)",
					"B (input.js:3:19)",
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.1 (native_stack_test.go:20)",
					"A (input.js:2:19)", "input.js:4:4",
				})

				is(ctx.Callee, "github.com/robertkrimen/otto.TestNativeStackFrames.func1.3")
				is(ctx.Filename, "native_stack_test.go")
				is(ctx.Line, 71)
				is(ctx.Column, 0)
			}

			{
				// no limit, skip innermost native frames
				ctx := c.Otto.ContextSkip(-1, true)

				is(ctx.Stacktrace, []string{
					"B (input.js:3:19)",
					"github.com/robertkrimen/otto.TestNativeStackFrames.func1.1 (native_stack_test.go:20)",
					"A (input.js:2:19)", "input.js:4:4",
				})

				is(ctx.Callee, "B")
				is(ctx.Filename, "input.js")
				is(ctx.Line, 3)
				is(ctx.Column, 19)
			}

			return UndefinedValue()
		})
		require.NoError(t, err)

		_, err = vm.Run(s)
		require.NoError(t, err)
	})
}
