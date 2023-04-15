package otto

// Option is an Otto option.
type Option func(*Otto)

// WithConsole replaces the builtin Console used with the given object.
func WithConsole(console Console) func(o *Otto) {
	return func(o *Otto) {
		o.console = console
	}
}
