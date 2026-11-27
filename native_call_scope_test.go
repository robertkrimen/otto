package otto_test

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/require"
)

func TestNativeCallEvalReceiver(t *testing.T) {
	for _, entry := range []string{"object", "value", "bound", "otto", "script"} {
		t.Run(entry, func(t *testing.T) {
			vm := otto.New()
			require.NoError(t, vm.Set("globalValue", 10))
			object, err := vm.Object(`({value: 7})`)
			require.NoError(t, err)
			require.NoError(t, object.Set("read", func(call otto.FunctionCall) otto.Value {
				result, evalErr := call.Otto.Eval(`var localOnly = 1; this.value + globalValue`)
				require.NoError(t, evalErr)
				return result
			}))
			require.NoError(t, vm.Set("receiver", object))
			var result otto.Value
			switch entry {
			case "object":
				result, err = object.Call("read")
			case "value":
				fn, getErr := object.Get("read")
				require.NoError(t, getErr)
				result, err = fn.Call(object.Value())
			case "bound":
				fn, runErr := vm.Run(`receiver.read.bind(receiver)`)
				require.NoError(t, runErr)
				result, err = fn.Call(otto.UndefinedValue())
			case "otto":
				result, err = vm.Call("receiver.read", nil)
			case "script":
				result, err = vm.Run(`receiver.read()`)
			}
			require.NoError(t, err)
			require.Equal(t, "17", result.String())
			result, err = vm.Eval(`typeof localOnly`)
			require.NoError(t, err)
			require.Equal(t, "undefined", result.String())
		})
	}
}

func TestNativeCallRestoresScopeAfterError(t *testing.T) {
	vm := otto.New()
	outer, err := vm.Object(`({value: "outer"})`)
	require.NoError(t, err)
	inner, err := vm.Object(`({value: "inner"})`)
	require.NoError(t, err)
	require.NoError(t, inner.Set("fail", func(call otto.FunctionCall) otto.Value {
		value, evalErr := call.Otto.Eval(`this.value`)
		require.NoError(t, evalErr)
		require.Equal(t, "inner", value.String())
		panic(call.Otto.MakeCustomError("Error", "expected"))
	}))
	require.NoError(t, outer.Set("read", func(call otto.FunctionCall) otto.Value {
		_, callErr := inner.Call("fail")
		require.ErrorContains(t, callErr, "expected")
		value, evalErr := call.Otto.Eval(`this.value`)
		require.NoError(t, evalErr)
		return value
	}))
	value, err := outer.Call("read")
	require.NoError(t, err)
	require.Equal(t, "outer", value.String())
	value, err = vm.Eval(`typeof this.value`)
	require.NoError(t, err)
	require.Equal(t, "undefined", value.String())
}
