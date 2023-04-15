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
