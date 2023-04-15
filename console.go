package otto

import (
	"fmt"
	"os"
	"strings"
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

func formatForConsole(argumentList []Value) string {
	output := []string{}
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

func builtinConsoleLog(call FunctionCall) Value {
	fmt.Fprintln(os.Stdout, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsoleError(call FunctionCall) Value {
	fmt.Fprintln(os.Stdout, formatForConsole(call.ArgumentList))
	return Value{}
}

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
