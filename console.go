package otto

import (
	"fmt"
	"os"
	"strings"
)

// Console describes the well known console object present in browsers and other runtimes.
type Console interface {
	Log(call FunctionCall) Value
	Trace(call FunctionCall) Value
	Debug(call FunctionCall) Value
	Info(call FunctionCall) Value
	Warn(call FunctionCall) Value
	Error(call FunctionCall) Value

	Dir(call FunctionCall) Value
	Time(call FunctionCall) Value
	TimeEnd(call FunctionCall) Value
	Assert(call FunctionCall) Value
}

func formatForConsole(argumentList []Value) string {
	output := []string{}
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

type builtinConsole struct{}

func (b *builtinConsole) Log(call FunctionCall) Value {
	return builtinConsoleLog(call)
}

func (b *builtinConsole) Trace(call FunctionCall) Value {
	return builtinConsoleLog(call)
}

func (b *builtinConsole) Debug(call FunctionCall) Value {
	return builtinConsoleLog(call)
}

func (b *builtinConsole) Info(call FunctionCall) Value {
	return builtinConsoleLog(call)
}

func (b *builtinConsole) Warn(call FunctionCall) Value {
	return builtinConsoleLog(call)
}

func (b *builtinConsole) Error(call FunctionCall) Value {
	return builtinConsoleLog(call)
}

func (b *builtinConsole) Dir(call FunctionCall) Value {
	return builtinConsoleDir(call)
}

func (b *builtinConsole) Time(call FunctionCall) Value {
	return builtinConsoleTime(call)
}

func (b *builtinConsole) TimeEnd(call FunctionCall) Value {
	return builtinConsoleTimeEnd(call)
}

func (b *builtinConsole) Assert(call FunctionCall) Value {
	return builtinConsoleAssert(call)
}

func builtinConsoleLog(call FunctionCall) Value {
	fmt.Fprintln(os.Stdout, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsoleError(call FunctionCall) Value {
	fmt.Fprintln(os.Stdout, formatForConsole(call.ArgumentList))
	return Value{}
}

type noopConsole struct {
}

func (n *noopConsole) Log(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Trace(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Debug(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Info(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Warn(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Error(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Dir(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Time(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) TimeEnd(_ FunctionCall) Value { return Value{} }

func (n *noopConsole) Assert(_ FunctionCall) Value { return Value{} }

// Nothing happens.
func builtinConsoleDir(call FunctionCall) Value {
	return Value{}
}

func builtinConsoleTime(call FunctionCall) Value {
	return Value{}
}

func builtinConsoleTimeEnd(call FunctionCall) Value {
	return Value{}
}

func builtinConsoleTrace(call FunctionCall) Value {
	return Value{}
}

func builtinConsoleAssert(call FunctionCall) Value {
	return Value{}
}
