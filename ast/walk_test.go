package ast_test

import (
	"testing"

	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"github.com/robertkrimen/otto/parser"
	"github.com/stretchr/testify/require"
)

type walker struct {
	stack             []ast.Node
	source            string
	shift             file.Idx
	seen              map[ast.Node]struct{}
	duplicate         int
	newExpressionIdx1 file.Idx
}

// push and pop below are to prove the symmetry of Enter/Exit calls

func (w *walker) push(n ast.Node) {
	w.stack = append(w.stack, n)
}

func (w *walker) pop(n ast.Node) {
	size := len(w.stack)
	if size <= 0 {
		panic("pop of empty stack")
	}

	if toPop := w.stack[size-1]; toPop != n {
		panic("pop: nodes do not equal")
	}

	w.stack[size-1] = nil
	w.stack = w.stack[:size-1]
}

func (w *walker) Enter(n ast.Node) ast.Visitor {
	w.push(n)
	if _, ok := w.seen[n]; ok {
		// Skip items we've already seen which occurs due to declarations.
		w.duplicate++
		return w
	}

	w.seen[n] = struct{}{}

	switch t := n.(type) {
	case *ast.Identifier:
		if t != nil {
			idx := n.Idx0() + w.shift - 1
			s := w.source[:idx] + "IDENT_" + w.source[idx:]
			w.source = s
			w.shift += 6
		}
	case *ast.VariableExpression:
		if t != nil {
			idx := n.Idx0() + w.shift - 1
			s := w.source[:idx] + "VAR_" + w.source[idx:]
			w.source = s
			w.shift += 4
		}
	case *ast.NewExpression:
		w.newExpressionIdx1 = n.Idx1()
	}

	return w
}

func (w *walker) Exit(n ast.Node) {
	w.pop(n)
}

func TestVisitorRewrite(t *testing.T) {
	source := `var b = function() {
		test();
		try {} catch(e) {}
		var test = "test(); var test = 1"
	} // test`
	program, err := parser.ParseFile(nil, "", source, 0)
	require.NoError(t, err)

	w := &walker{
		source: source,
		seen:   make(map[ast.Node]struct{}),
	}
	ast.Walk(w, program)

	xformed := `var VAR_b = function() {
		IDENT_test();
		try {} catch(IDENT_e) {}
		var VAR_test = "test(); var test = 1"
	} // test`

	require.Equal(t, xformed, w.source)
	require.Len(t, w.stack, 0)
	require.Equal(t, w.duplicate, 0)
}

func Test_issue261(t *testing.T) {
	tests := map[string]struct {
		code string
		want file.Idx
	}{
		"no-parenthesis": {
			code: `var i = new Image;`,
			want: 18,
		},
		"no-args": {
			code: `var i = new Image();`,
			want: 20,
		},
		"two-args": {
			code: `var i = new Image(1, 2);`,
			want: 24,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			prog, err := parser.ParseFile(nil, "", tt.code, 0)
			require.NoError(t, err)

			w := &walker{
				source: tt.code,
				seen:   make(map[ast.Node]struct{}),
			}
			ast.Walk(w, prog)

			require.Equal(t, tt.want, w.newExpressionIdx1)
			require.Len(t, w.stack, 0)
			require.Equal(t, w.duplicate, 0)
		})
	}
}

func TestBadStatement(t *testing.T) {
	source := `
	var abc;
	break; do {
	} while(true);
`
	program, err := parser.ParseFile(nil, "", source, 0)
	require.ErrorContains(t, err, "Illegal break statement")

	w := &walker{
		source: source,
		seen:   make(map[ast.Node]struct{}),
	}

	require.NotPanics(t, func() {
		ast.Walk(w, program)
	})
}
