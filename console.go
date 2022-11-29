package otto

import (
	"fmt"
	"os"
)

// Console is implemented by type which can be used to process console output.
type Console interface {
	Log(v ...interface{})
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

type console struct {
}

func (c *console) log(v ...interface{}) {
	fmt.Fprintln(os.Stdout, v...)
}

func (c *console) Log(v ...interface{}) {
	c.log(v...)
}

func (c *console) Trace(v ...interface{}) {
	c.log(v...)
}

func (c *console) Debug(v ...interface{}) {
	c.log(v...)
}

func (c *console) Info(v ...interface{}) {
	c.log(v...)
}

func (c *console) Warn(v ...interface{}) {
	c.log(v...)
}

func (c *console) Error(v ...interface{}) {
	c.log(v...)
}

// TODO use x/exp/slices
func argsAsAny(argumentList []Value) []interface{} {
	output := make([]interface{}, 0, len(argumentList))
	for _, argument := range argumentList {
		output = append(output, argument.String())
	}
	return output
}

func callWithArgs(callee func(v ...interface{})) func(call FunctionCall) Value {
	return func(call FunctionCall) Value {
		callee(argsAsAny(call.ArgumentList)...)
		return Value{}
	}
}

// Nothing happens.
func builtinConsole_dir(call FunctionCall) Value {
	return Value{}
}

func builtinConsole_time(call FunctionCall) Value {
	return Value{}
}

func builtinConsole_timeEnd(call FunctionCall) Value {
	return Value{}
}

func builtinConsole_assert(call FunctionCall) Value {
	return Value{}
}

func (runtime *_runtime) newConsole() *_object {
	return newConsoleObject(runtime)
}
