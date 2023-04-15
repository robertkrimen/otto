package otto

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

}

// ExampleWithNoopConsole shows how to use WithNoopConsole while ensuring that the used no-op console
// does not actually print anything to stdout.
func ExampleWithNoopConsole() { //nolint: govet
	vm := New(WithNoopConsole())

	src := `console.log("Hello, World.");
console.dir(42);
console.trace(1, "abc", [1, 2, 3]);`
	value, err := vm.Run(src)
	fmt.Println(value)
	fmt.Println(err)

	// Output:
	// undefined
	// <nil>
}
