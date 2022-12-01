package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
)

var flagUnderscore *bool = flag.Bool("underscore", true, "Load underscore into the runtime environment")

func readSource(filename string) ([]byte, error) {
	if filename == "" || filename == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(filename) //nolint: gosec
}

func main() {
	flag.Parse()

	if !*flagUnderscore {
		underscore.Disable()
	}

	err := func() error {
		src, err := readSource(flag.Arg(0))
		if err != nil {
			return err
		}

		vm := otto.New()
		_, err = vm.Run(src)
		return err
	}()
	if err != nil {
		var oerr *otto.Error
		if errors.As(err, &oerr) {
			fmt.Fprint(os.Stderr, oerr.String())
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(64)
	}
}
