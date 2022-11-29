package otto

// Option is an Otto option.
type Option func(*Otto)

// WithConsole adds a console option to the otto instance.
func WithConsole(console Console) func(o *Otto) {
	return func(o *Otto) {
		o.console = console
	}
}
