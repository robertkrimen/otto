package otto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScript(t *testing.T) {
	tt(t, func() {
		vm := New()

		script, err := vm.Compile("xyzzy", `var abc; if (!abc) abc = 0; abc += 2; abc;`)
		require.NoError(t, err)

		str := script.String()
		is(str, "// xyzzy\nvar abc; if (!abc) abc = 0; abc += 2; abc;")

		value, err := vm.Run(script)
		require.NoError(t, err)
		is(value, 2)

		// TODO(steve): Fix the underlying issues as to why this returns early.
		if true {
			return
		}

		tmp, err := script.marshalBinary()
		require.NoError(t, err)
		is(len(tmp), 1228)

		{
			script := &Script{}
			err = script.unmarshalBinary(tmp)
			require.NoError(t, err)

			is(script.String(), str)

			value, err = vm.Run(script)
			require.NoError(t, err)
			is(value, 4)

			tmp, err = script.marshalBinary()
			require.NoError(t, err)
			is(len(tmp), 1228)
		}

		{
			script := &Script{}
			err = script.unmarshalBinary(tmp)
			require.NoError(t, err)

			is(script.String(), str)

			value, err := vm.Run(script)
			require.NoError(t, err)
			is(value, 6)

			tmp, err = script.marshalBinary()
			require.NoError(t, err)
			is(len(tmp), 1228)
		}

		{
			version := scriptVersion
			scriptVersion = "bogus"

			script := &Script{}
			err = script.unmarshalBinary(tmp)
			is(err, "version mismatch")

			is(script.String(), "// \n")
			is(script.version, "")
			is(script.program == nil, true)
			is(script.filename, "")
			is(script.src, "")

			scriptVersion = version
		}
	})
}

func TestFunctionCall_CallerLocation(t *testing.T) {
	tt(t, func() {
		vm := New()
		err := vm.Set("loc", func(call FunctionCall) Value {
			return toValue(call.CallerLocation())
		})
		require.NoError(t, err)
		script, err := vm.Compile("somefile.js", `var where = loc();`)
		require.NoError(t, err)
		_, err = vm.Run(script)
		require.NoError(t, err)
		where, err := vm.Get("where")
		require.NoError(t, err)
		is(where, "somefile.js:1:13")
	})
}
