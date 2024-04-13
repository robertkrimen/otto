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

		val, err := vm.Run(script)
		require.NoError(t, err)
		is(val, 2)

		// TODO(steve): Fix the underlying issues as to why this returns early.
		if true {
			return
		}

		tmp, err := script.marshalBinary()
		require.NoError(t, err)
		is(len(tmp), 1228)

		{
			script2 := &Script{}
			err = script2.unmarshalBinary(tmp)
			require.NoError(t, err)

			is(script2.String(), str)

			val, err = vm.Run(script2)
			require.NoError(t, err)
			is(val, 4)

			tmp, err = script2.marshalBinary()
			require.NoError(t, err)
			is(len(tmp), 1228)
		}

		{
			script2 := &Script{}
			err = script2.unmarshalBinary(tmp)
			require.NoError(t, err)

			is(script2.String(), str)

			val2, err2 := vm.Run(script2)
			require.NoError(t, err2)
			is(val2, 6)

			tmp, err2 = script2.marshalBinary()
			require.NoError(t, err2)
			is(len(tmp), 1228)
		}

		{
			version := scriptVersion
			scriptVersion = "bogus"

			script2 := &Script{}
			err = script2.unmarshalBinary(tmp)
			is(err, "version mismatch")

			is(script2.String(), "// \n")
			is(script2.version, "")
			is(script2.program == nil, true)
			is(script2.filename, "")
			is(script2.src, "")

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
