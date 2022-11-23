package otto

type Option func(*Otto)

func WithLogger(log Logger) func(o *Otto) {
	return func(o *Otto) {
		o.log = log
	}
}
