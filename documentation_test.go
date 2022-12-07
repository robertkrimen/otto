package otto

import (
	"fmt"
	"strings"
)

func ExampleSynopsis() { //nolint: govet
	vm := New()
	vm.Run(`
        abc = 2 + 2;
        console.log("The value of abc is " + abc); // 4
    `)

	value, _ := vm.Get("abc")
	{
		value, _ := value.ToInteger()
		fmt.Println(value)
	}

	vm.Set("def", 11)
	vm.Run(`
        console.log("The value of def is " + def);
    `)

	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
        console.log(xyzzy.length);
    `)

	value, _ = vm.Run("xyzzy.length")
	{
		value, _ := value.ToInteger()
		fmt.Println(value)
	}

	value, err := vm.Run("abcdefghijlmnopqrstuvwxyz.length")
	fmt.Println(value)
	fmt.Println(err)

	vm.Set("sayHello", func(call FunctionCall) Value {
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return UndefinedValue()
	})

	vm.Set("twoPlus", func(call FunctionCall) Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})

	value, _ = vm.Run(`
        sayHello("Xyzzy");
        sayHello();

        result = twoPlus(2.0);
    `)
	fmt.Println(value)

	// custom object
	console := map[string]interface{}{
		"log": func(call FunctionCall) Value {
			args := make([]string, 0, len(call.ArgumentList))
			for _, arg := range call.ArgumentList {
				args = append(args, fmt.Sprintf("%v", arg))
			}
			fmt.Println("console.log:", strings.Join(args, " "))
			return UndefinedValue()
		},
	}

	if err := vm.Set("console", console); err != nil {
		panic(fmt.Errorf("console error: %w", err))
	}

	value, err = vm.Run(`console.log("Hello, World.");`)
	fmt.Println(value)
	fmt.Println(err)

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
	// console.log: Hello, World.
	// undefined
	// <nil>
}

type printer struct{}

func (c *printer) Log(v ...interface{})   { fmt.Print("console.log: "); fmt.Println(v...) }
func (c *printer) Trace(v ...interface{}) { fmt.Print("console.trace: "); fmt.Println(v...) }
func (c *printer) Debug(v ...interface{}) { fmt.Print("console.debug: "); fmt.Println(v...) }
func (c *printer) Info(v ...interface{})  { fmt.Print("console.info: "); fmt.Println(v...) }
func (c *printer) Warn(v ...interface{})  { fmt.Print("console.warn: "); fmt.Println(v...) }
func (c *printer) Error(v ...interface{}) { fmt.Print("console.error: "); fmt.Println(v...) }

func ExampleConsole() {
	vm := New(WithConsole(new(printer)))

	value, err := vm.Run(`console.log("Hello, World.");`)
	fmt.Println(value)
	fmt.Println(err)

	// Output:
	// console.log: Hello, World.
	// undefined
	// <nil>
}
