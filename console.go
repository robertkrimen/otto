package otto

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type ConsoleWriter interface {
	Write(call FunctionCall, str string)
}

type ConsoleIOWriter struct {
	writer io.Writer
}

func (w ConsoleIOWriter) Write(call FunctionCall, str string) {
	w.writer.Write([]byte(str))
}

var ConsoleLogWriter ConsoleWriter = ConsoleIOWriter{os.Stdout}
var ConsoleErrorWriter ConsoleWriter = ConsoleIOWriter{os.Stdout}
var ConsoleDebugWriter ConsoleWriter = ConsoleIOWriter{os.Stdout}
var ConsoleInfoWriter ConsoleWriter = ConsoleIOWriter{os.Stdout}
var ConsoleWarnWriter ConsoleWriter = ConsoleIOWriter{os.Stdout}

func formatForConsole(argumentList []Value) string {
	output := []string{}
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

func builtinConsole_log(call FunctionCall) Value {
	ConsoleLogWriter.Write(call, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_debug(call FunctionCall) Value {
	ConsoleDebugWriter.Write(call, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_info(call FunctionCall) Value {
	ConsoleInfoWriter.Write(call, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_warn(call FunctionCall) Value {
	ConsoleWarnWriter.Write(call, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_error(call FunctionCall) Value {
	ConsoleErrorWriter.Write(call, formatForConsole(call.ArgumentList))
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
