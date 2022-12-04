package otto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCloneGetterSetter(t *testing.T) {
	vm := New()

	_, err := vm.Run(`var x = Object.create(null, {
    x: {
      get: function() {},
      set: function() {},
    },
  })`)
	require.NoError(t, err)
	require.NotPanics(t, func() {
		vm.Copy()
	})
}
