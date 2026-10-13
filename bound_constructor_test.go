package otto_test

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/require"
)

func TestBoundConstructorsForwardArguments(t *testing.T) {
	tests := []struct{ name, source, want string }{
		{"array", `var Bound = Array.bind(null, "first"); JSON.stringify(new Bound("second"));`, `["first","second"]`},
		{"nested array", `var Bound = Array.bind(null, "first").bind(null, "second"); JSON.stringify(new Bound("third"));`, `["first","second","third"]`},
		{"date", `var Bound = Date.bind(null, 2020); var d = new Bound(1, 2); [d.getFullYear(), d.getMonth(), d.getDate()].join(",");`, "2020,1,2"},
		{"regexp", `var Bound = RegExp.bind(null, "hello"); var re = new Bound("i"); String(re.test("HELLO"));`, "true"},
		{"nested function", `function Constructor(a,b,c) { this.values=[a,b,c]; } var Bound = Constructor.bind(null,"first").bind(null,"second"); JSON.stringify(new Bound("third").values);`, `["first","second","third"]`},
		{"prototype and receiver", `function Constructor(a,b) { this.values=[a,b]; } Constructor.prototype.marker="prototype"; var receiver={}; var Bound=Constructor.bind(receiver,"first"); var a=new Bound("second"); var b=new Bound("third"); JSON.stringify([a.values,b.values,a.marker,a instanceof Constructor,typeof receiver.values]);`, `[["first","second"],["first","third"],"prototype",true,"undefined"]`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vm := otto.New()
			value, err := vm.Run(test.source)
			require.NoError(t, err)
			require.Equal(t, test.want, value.String())
		})
	}
}

func TestBoundNonConstructorReturnsError(t *testing.T) {
	vm := otto.New()
	require.NotPanics(t, func() {
		_, err := vm.Run(`new (Math.max.bind(null, 1))(2)`)
		require.ErrorContains(t, err, "not a constructor")
	})
}
