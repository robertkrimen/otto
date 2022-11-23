package otto

import (
	"fmt"
	"io"
)

type Logger interface {
	Print(v ...interface{})
	Error(v ...interface{})
}

type console struct {
	out io.Writer
}

func (l *console) Print(v ...interface{}) {
	fmt.Fprintln(l.out, v...)
}

func (l *console) Error(v ...interface{}) {
	fmt.Fprintln(l.out, v...)
}
