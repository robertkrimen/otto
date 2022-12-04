package otto

import (
	"testing"

	"github.com/robertkrimen/otto/parser"
)

func Test_cmpl(t *testing.T) {
	tt(t, func() {
		vm := New()

		test := func(src string, expect ...interface{}) {
			program, err := parser.ParseFile(nil, "", src, 0)
			is(err, nil)
			{
				program := cmplParse(program)
				value := vm.runtime.cmplEvaluateNodeProgram(program, false)
				if len(expect) > 0 {
					is(value, expect[0])
				}
			}
		}

		test(``, Value{})

		test(`var abc = 1; abc;`, 1)

		test(`var abc = 1 + 1; abc;`, 2)

		test(`1 + 2;`, 3)
	})
}

func TestParse_cmpl(t *testing.T) {
	tt(t, func() {
		test := func(src string) {
			program, err := parser.ParseFile(nil, "", src, 0)
			is(err, nil)
			is(cmplParse(program), "!=", nil)
		}

		test(``)

		test(`var abc = 1; abc;`)

		test(`
            function abc() {
                return;
            }
        `)
	})
}
