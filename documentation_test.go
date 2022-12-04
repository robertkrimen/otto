package otto

import (
	"fmt"
	"os"
)

func ExampleSynopsis() { //nolint: govet
	vm := New()
	_, err := vm.Run(`
        abc = 2 + 2;
        console.log("The value of abc is " + abc); // 4
    `)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	value, err := vm.Get("abc")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	iv, err := value.ToInteger()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(iv)

	err = vm.Set("def", 11)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	_, err = vm.Run(`
        console.log("The value of def is " + def);
    `)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = vm.Set("xyzzy", "Nothing happens.")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	_, err = vm.Run(`
        console.log(xyzzy.length);
    `)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	value, err = vm.Run("xyzzy.length")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	iv, err = value.ToInteger()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(iv)

	value, err = vm.Run("abcdefghijlmnopqrstuvwxyz.length")
	fmt.Println(value)
	fmt.Println(err) // Expected error.

	err = vm.Set("sayHello", func(call FunctionCall) Value {
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return UndefinedValue()
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = vm.Set("twoPlus", func(call FunctionCall) Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	value, err = vm.Run(`
        sayHello("Xyzzy");
        sayHello();

        result = twoPlus(2.0);
    `)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(value)

	// Output:
	// The value of abc is 4
	// 4
	// The value of def is 11
	// 16
	// 16
	// undefined
	// ReferenceError: 'abcdefghijlmnopqrstuvwxyz' is not defined
	// Hello, Xyzzy.
	// Hello, undefined.
	// 4
}

func ExampleConsole() { //nolint: govet
	vm := New()
	console := map[string]interface{}{
		"log": func(call FunctionCall) Value {
			fmt.Println("console.log:", formatForConsole(call.ArgumentList))
			return UndefinedValue()
		},
	}

	err := vm.Set("console", console)
	if err != nil {
		panic(fmt.Errorf("console error: %w", err))
	}

	value, err := vm.Run(`console.log("Hello, World.");`)
	fmt.Println(value)
	fmt.Println(err)

	// Output:
	// console.log: Hello, World.
	// undefined
	// <nil>
}
