package otto

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func formatForConsole(argumentList []Value) string {
	output := []string{}
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

func builtinConsoleLog(call FunctionCall) Value {
	var w io.Writer
	if call.runtime.stdoutWriter != nil {
		w = call.runtime.stdoutWriter
	} else {
		w = os.Stdout
	}
	fmt.Fprintln(w, formatForConsole(call.ArgumentList)) //nolint:errcheck // Nothing we can do if this fails.
	return Value{}
}

func builtinConsoleError(call FunctionCall) Value {
	var w io.Writer
	if call.runtime.stdoutWriter != nil {
		w = call.runtime.stdoutWriter
	} else {
		w = os.Stdout
	}
	fmt.Fprintln(w, formatForConsole(call.ArgumentList)) //nolint:errcheck // Nothing we can do if this fails.
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
