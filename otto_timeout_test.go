package otto

import (
	"testing"
	"time"
)

func TestRunWithTimeout(t *testing.T) {
	tt(t, func() {
		vm := New()
		_, err := vm.RunWithTimeout(time.Nanosecond, `while (true) {}`)

		is(err, TimeoutError)
	})
}

func TestRunWithTimeoutSuccessfully(t *testing.T) {
	tt(t, func() {
		vm := New()
		val, err := vm.RunWithTimeout(time.Millisecond, `40 + 2`)

		is(err, nil)
		intVal, err := val.ToInteger()

		is(err, nil)
		is(intVal, 42)
	})
}

func TestEvalWithTimeout(t *testing.T) {
	tt(t, func() {
		vm := New()
		_, err := vm.EvalWithTimeout(time.Nanosecond, `while (true) {}`)

		is(err, TimeoutError)
	})
}

func TestEvalWithTimeoutSuccessfully(t *testing.T) {
	tt(t, func() {
		vm := New()
		val, err := vm.EvalWithTimeout(time.Millisecond, `40 + 2`)

		is(err, nil)
		intVal, err := val.ToInteger()

		is(err, nil)
		is(intVal, 42)
	})
}

func TestCallWithTimeout(t *testing.T) {
	tt(t, func() {
		vm := New()
		_, err := vm.Run(`function endless() {while (true) {} }`)

		is(err, nil)

		_, err = vm.CallWithTimeout(time.Nanosecond, "endless", nil)

		is(err, TimeoutError)
	})
}

func TestCallWithTimeoutSuccessfully(t *testing.T) {
	tt(t, func() {
		vm := New()
		_, err := vm.Run(`function finite() { return true; }`)

		is(err, nil)

		//val, err := vm.Call("finite", nil)
		val, err := vm.CallWithTimeout(time.Millisecond, "finite", nil)

		is(err, nil)

		boolVal, err := val.ToBoolean()

		is(err, nil)
		is(boolVal, true)
	})
}
