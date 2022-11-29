package ast_test

import (
	"testing"

	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"github.com/robertkrimen/otto/parser"
	"github.com/stretchr/testify/require"
)

type walker struct {
	stack     []ast.Node
	source    string
	shift     file.Idx
	seen      map[ast.Node]struct{}
	duplicate int
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

	toPop := w.stack[size-1]
	if toPop != n {
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

	if id, ok := n.(*ast.Identifier); ok && id != nil {
		idx := n.Idx0() + w.shift - 1
		s := w.source[:idx] + "IDENT_" + w.source[idx:]
		w.source = s
		w.shift += 6
	}
	if v, ok := n.(*ast.VariableExpression); ok && v != nil {
		idx := n.Idx0() + w.shift - 1
		s := w.source[:idx] + "VAR_" + w.source[idx:]
		w.source = s
		w.shift += 4
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
