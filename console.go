package otto

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var ConsoleLogWriter io.Writer = os.Stdout
var ConsoleErrorWriter io.Writer = os.Stdout
var ConsoleDebugWriter io.Writer = os.Stdout
var ConsoleInfoWriter io.Writer = os.Stdout
var ConsoleWarnWriter io.Writer = os.Stdout

func formatForConsole(argumentList []Value) string {
	output := []string{}
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

func builtinConsole_log(call FunctionCall) Value {
	fmt.Fprintln(ConsoleLogWriter, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_debug(call FunctionCall) Value {
	fmt.Fprintln(ConsoleDebugWriter, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_info(call FunctionCall) Value {
	fmt.Fprintln(ConsoleInfoWriter, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_warn(call FunctionCall) Value {
	fmt.Fprintln(ConsoleWarnWriter, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_error(call FunctionCall) Value {
	fmt.Fprintln(ConsoleErrorWriter, formatForConsole(call.ArgumentList))
	return Value{}
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

func builtinConsole_trace(call FunctionCall) Value {
	return Value{}
}

func builtinConsole_assert(call FunctionCall) Value {
	return Value{}
}

func (runtime *_runtime) newConsole() *_object {

	return newConsoleObject(runtime)
}
