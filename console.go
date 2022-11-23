package otto

// TODO use x/exp/slices
func argsAsAny(argumentList []Value) []interface{} {
	output := make([]interface{}, 0, len(argumentList))
	for _, argument := range argumentList {
		output = append(output, argument.String())
	}
	return output
}

func builtinConsole_log(otto *Otto) func(call FunctionCall) Value {
	return func(call FunctionCall) Value {
		// println("---------")
		// args := argsAsAny(call.ArgumentList)
		// fmt.Printf("args %v\n", args)
		// otto.log.Print(args...)
		otto.log.Print(argsAsAny(call.ArgumentList)...)
		return Value{}
	}
}

func builtinConsole_error(otto *Otto) func(call FunctionCall) Value {
	return func(call FunctionCall) Value {
		otto.log.Error(argsAsAny(call.ArgumentList)...)
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

func builtinConsole_trace(call FunctionCall) Value {
	return Value{}
}

func builtinConsole_assert(call FunctionCall) Value {
	return Value{}
}

func (runtime *_runtime) newConsole() *_object {
	return newConsoleObject(runtime)
}
