package ast_test

import (
	"fmt"
	"log"

	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"github.com/robertkrimen/otto/parser"
)

func ExampleVisitorRewrite() {
	source := `var b = function() {test(); try {} catch(e) {} var test = "test(); var test = 1"} // test`
	program, err := parser.ParseFile(nil, "", source, 0)
	if err != nil {
		log.Fatal(err)
	}

	var shift file.Idx

	ast.Walk(ast.VisitorFunc(func(v ast.Visitor, n ast.Node) ast.Visitor {
		if n == nil {
			return v
		}
		if id, ok := n.(*ast.Identifier); ok && id != nil {
			idx := n.Idx0() + shift - 1
			s := source[:idx] + "new_" + source[idx:]
			source = s
			shift += 4
		}
		if v, ok := n.(*ast.VariableExpression); ok && v != nil {
			idx := n.Idx0() + shift - 1
			s := source[:idx] + "varnew_" + source[idx:]
			source = s
			shift += 7
		}

		return v
	}), program)

	fmt.Println(source)
	// Output: var varnew_b = function() {new_test(); try {} catch(new_e) {} var varnew_test = "test(); var test = 1"} // test
}
