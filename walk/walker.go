package walk

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
)

// Walker can walk a given AST with a visitor
type Walker struct {
	Visitor Visitor
}

func NewWalker(visitor Visitor) *Walker {
	return &Walker{
		Visitor: visitor,
	}
}

// Visitor interface for the walker.
type Visitor interface {
	VisitArray(walker *Walker, node *ast.ArrayLiteral, metadata []Metadata) Metadata
	VisitAssign(walker *Walker, node *ast.AssignExpression, metadata []Metadata) Metadata
	VisitBad(walker *Walker, node *ast.BadExpression, metadata []Metadata) Metadata
	VisitBadStatement(walker *Walker, node *ast.BadStatement, metadata []Metadata) Metadata
	VisitBinary(walker *Walker, node *ast.BinaryExpression, metadata []Metadata) Metadata
	VisitBlock(walker *Walker, node *ast.BlockStatement, metadata []Metadata) Metadata
	VisitBoolean(walker *Walker, node *ast.BooleanLiteral, metadata []Metadata) Metadata
	VisitBracket(walker *Walker, node *ast.BracketExpression, metadata []Metadata) Metadata
	VisitBranch(walker *Walker, node *ast.BranchStatement, metadata []Metadata) Metadata
	VisitCall(walker *Walker, node *ast.CallExpression, metadata []Metadata) Metadata
	VisitCase(walker *Walker, node *ast.CaseStatement, metadata []Metadata) Metadata
	VisitCatch(walker *Walker, node *ast.CatchStatement, metadata []Metadata) Metadata
	VisitConditional(walker *Walker, node *ast.ConditionalExpression, metadata []Metadata) Metadata
	VisitDebugger(walker *Walker, node *ast.DebuggerStatement, metadata []Metadata) Metadata
	VisitDot(walker *Walker, node *ast.DotExpression, metadata []Metadata) Metadata
	VisitDoWhile(walker *Walker, node *ast.DoWhileStatement, metadata []Metadata) Metadata
	VisitEmpty(walker *Walker, node *ast.EmptyExpression, metadata []Metadata) Metadata
	VisitEmptyStatement(walker *Walker, node *ast.EmptyStatement, metadata []Metadata) Metadata
	VisitExpression(walker *Walker, node *ast.ExpressionStatement, metadata []Metadata) Metadata
	VisitForIn(walker *Walker, node *ast.ForInStatement, metadata []Metadata) Metadata
	VisitFor(walker *Walker, node *ast.ForStatement, metadata []Metadata) Metadata
	VisitFunction(walker *Walker, node *ast.FunctionLiteral, metadata []Metadata) Metadata
	VisitFunctionStatement(walker *Walker, node *ast.FunctionStatement, metadata []Metadata) Metadata
	VisitIdentifier(walker *Walker, node *ast.Identifier, metadata []Metadata) Metadata
	VisitIf(walker *Walker, node *ast.IfStatement, metadata []Metadata) Metadata
	VisitLabelled(walker *Walker, node *ast.LabelledStatement, metadata []Metadata) Metadata
	VisitNew(walker *Walker, node *ast.NewExpression, metadata []Metadata) Metadata
	VisitNull(walker *Walker, node *ast.NullLiteral, metadata []Metadata) Metadata
	VisitNumber(walker *Walker, node *ast.NumberLiteral, metadata []Metadata) Metadata
	VisitObject(walker *Walker, node *ast.ObjectLiteral, metadata []Metadata) Metadata
	VisitProgram(walker *Walker, node *ast.Program, metadata []Metadata) Metadata
	VisitReturn(walker *Walker, node *ast.ReturnStatement, metadata []Metadata) Metadata
	VisitRegex(walker *Walker, node *ast.RegExpLiteral, metadata []Metadata) Metadata
	VisitSequence(walker *Walker, node *ast.SequenceExpression, metadata []Metadata) Metadata
	VisitString(walker *Walker, node *ast.StringLiteral, metadata []Metadata) Metadata
	VisitSwitch(walker *Walker, node *ast.SwitchStatement, metadata []Metadata) Metadata
	VisitThis(walker *Walker, node *ast.ThisExpression, metadata []Metadata) Metadata
	VisitThrow(walker *Walker, node *ast.ThrowStatement, metadata []Metadata) Metadata
	VisitTry(walker *Walker, node *ast.TryStatement, metadata []Metadata) Metadata
	VisitUnary(walker *Walker, node *ast.UnaryExpression, metadata []Metadata) Metadata
	VisitVariable(walker *Walker, node *ast.VariableExpression, metadata []Metadata) Metadata
	VisitVariableStatement(walker *Walker, node *ast.VariableStatement, metadata []Metadata) Metadata
	VisitWhile(walker *Walker, node *ast.WhileStatement, metadata []Metadata) Metadata
	VisitWith(walker *Walker, node *ast.WithStatement, metadata []Metadata) Metadata
}

// Begin the walk of the given AST node
func (w *Walker) Begin(node ast.Node) {
	md := []Metadata{NewMetadata(nil)}
	w.Walk(node, md)
}

// CollectScope collects information about the given scope
func CollectScope(metadata Metadata, declarations []ast.Declaration) {
	// Initialize the scope variables field in the metadata
	vars, ok := metadata[Vars].(Variables)
	if !ok {
		vars = NewVariables()
		metadata[Vars] = vars
	}

	for _, vd := range declarations {
		switch d := vd.(type) {
		case *ast.VariableDeclaration:
			for _, v := range d.List {
				vars[v.Name] = v.Idx
			}
		}
	}
}

func FindVariable(metadata []Metadata, name string) file.Idx {
	md := metadata[len(metadata)-1]

	vars, ok := md[Vars].(Variables)
	if ok {
		for v, i := range vars {
			if v == name {
				return i
			}
		}
	}

	if len(metadata) > 1 {
		metadata = metadata[:len(metadata)-1]
		return FindVariable(metadata, name)
	}

	return -1
}

// Walk the AST, including metadata
func (w *Walker) Walk(node ast.Node, metadata []Metadata) Metadata {

	// Create metadata for current node
	md := NewMetadata(node)

	// Scope things
	switch n := node.(type) {
	case *ast.Program:
		CollectScope(md, n.DeclarationList)
	case *ast.FunctionLiteral:
		CollectScope(md, n.DeclarationList)
	}

	// Append the node
	metadata = append(metadata, md)

	switch n := node.(type) {
	case *ast.ArrayLiteral:
		return w.Visitor.VisitArray(w, n, metadata)
	case *ast.AssignExpression:
		return w.Visitor.VisitAssign(w, n, metadata)
	case *ast.BadExpression:
		return w.Visitor.VisitBad(w, n, metadata)
	case *ast.BadStatement:
		return w.Visitor.VisitBadStatement(w, n, metadata)
	case *ast.BinaryExpression:
		return w.Visitor.VisitBinary(w, n, metadata)
	case *ast.BlockStatement:
		return w.Visitor.VisitBlock(w, n, metadata)
	case *ast.BooleanLiteral:
		return w.Visitor.VisitBoolean(w, n, metadata)
	case *ast.BracketExpression:
		return w.Visitor.VisitBracket(w, n, metadata)
	case *ast.BranchStatement:
		return w.Visitor.VisitBranch(w, n, metadata)
	case *ast.CallExpression:
		return w.Visitor.VisitCall(w, n, metadata)
	case *ast.CaseStatement:
		return w.Visitor.VisitCase(w, n, metadata)
	case *ast.CatchStatement:
		return w.Visitor.VisitCatch(w, n, metadata)
	case *ast.ConditionalExpression:
		return w.Visitor.VisitConditional(w, n, metadata)
	case *ast.DebuggerStatement:
		return w.Visitor.VisitDebugger(w, n, metadata)
	case *ast.DotExpression:
		return w.Visitor.VisitDot(w, n, metadata)
	case *ast.DoWhileStatement:
		return w.Visitor.VisitDoWhile(w, n, metadata)
	case *ast.EmptyExpression:
		return w.Visitor.VisitEmpty(w, n, metadata)
	case *ast.EmptyStatement:
		return w.Visitor.VisitEmptyStatement(w, n, metadata)
	case *ast.ExpressionStatement:
		return w.Visitor.VisitExpression(w, n, metadata)
	case *ast.ForInStatement:
		return w.Visitor.VisitForIn(w, n, metadata)
	case *ast.ForStatement:
		return w.Visitor.VisitFor(w, n, metadata)
	case *ast.FunctionLiteral:
		return w.Visitor.VisitFunction(w, n, metadata)
	case *ast.FunctionStatement:
		return w.Visitor.VisitFunctionStatement(w, n, metadata)
	case *ast.Identifier:
		return w.Visitor.VisitIdentifier(w, n, metadata)
	case *ast.IfStatement:
		return w.Visitor.VisitIf(w, n, metadata)
	case *ast.LabelledStatement:
		return w.Visitor.VisitLabelled(w, n, metadata)
	case *ast.NewExpression:
		return w.Visitor.VisitNew(w, n, metadata)
	case *ast.NullLiteral:
		return w.Visitor.VisitNull(w, n, metadata)
	case *ast.NumberLiteral:
		return w.Visitor.VisitNumber(w, n, metadata)
	case *ast.ObjectLiteral:
		return w.Visitor.VisitObject(w, n, metadata)
	case *ast.Program:
		return w.Visitor.VisitProgram(w, n, metadata)
	case *ast.ReturnStatement:
		return w.Visitor.VisitReturn(w, n, metadata)
	case *ast.RegExpLiteral:
		return w.Visitor.VisitRegex(w, n, metadata)
	case *ast.SequenceExpression:
		return w.Visitor.VisitSequence(w, n, metadata)
	case *ast.StringLiteral:
		return w.Visitor.VisitString(w, n, metadata)
	case *ast.SwitchStatement:
		return w.Visitor.VisitSwitch(w, n, metadata)
	case *ast.ThisExpression:
		return w.Visitor.VisitThis(w, n, metadata)
	case *ast.ThrowStatement:
		return w.Visitor.VisitThrow(w, n, metadata)
	case *ast.TryStatement:
		return w.Visitor.VisitTry(w, n, metadata)
	case *ast.UnaryExpression:
		return w.Visitor.VisitUnary(w, n, metadata)
	case *ast.VariableExpression:
		return w.Visitor.VisitVariable(w, n, metadata)
	case *ast.VariableStatement:
		return w.Visitor.VisitVariableStatement(w, n, metadata)
	case *ast.WhileStatement:
		return w.Visitor.VisitWhile(w, n, metadata)
	case *ast.WithStatement:
		return w.Visitor.VisitWith(w, n, metadata)
	}

	return nil
}

// VisitorImpl is a default implementation of the Visitor interface
type VisitorImpl struct {
}

func (v *VisitorImpl) VisitProgram(w *Walker, node *ast.Program, metadata []Metadata) Metadata {
	for _, e := range node.Body {
		w.Walk(e, metadata)
	}

	// Walking function and variable declarations
	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			//w.Walk(value.Function, metadata)
		case *ast.VariableDeclaration:
			// Not needed, variable declarations are found in the AST
		default:
			panic(fmt.Errorf("Here be dragons: visit Program.DeclarationList(%T)", value))
		}
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitArray(w *Walker, node *ast.ArrayLiteral, metadata []Metadata) Metadata {
	for _, e := range node.Value {
		w.Walk(e, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitAssign(w *Walker, node *ast.AssignExpression, metadata []Metadata) Metadata {
	w.Walk(node.Left, metadata)
	w.Walk(node.Right, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBad(w *Walker, node *ast.BadExpression, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBadStatement(w *Walker, node *ast.BadStatement, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBinary(w *Walker, node *ast.BinaryExpression, metadata []Metadata) Metadata {
	w.Walk(node.Left, metadata)
	w.Walk(node.Right, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBlock(w *Walker, node *ast.BlockStatement, metadata []Metadata) Metadata {
	for _, value := range node.List {
		w.Walk(value, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBoolean(w *Walker, node *ast.BooleanLiteral, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBracket(w *Walker, node *ast.BracketExpression, metadata []Metadata) Metadata {
	w.Walk(node.Left, metadata)
	w.Walk(node.Member, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitBranch(w *Walker, node *ast.BranchStatement, metadata []Metadata) Metadata {
	if node.Label != nil {
		w.Walk(node.Label, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitCall(w *Walker, node *ast.CallExpression, metadata []Metadata) Metadata {
	w.Walk(node.Callee, metadata)
	for _, value := range node.ArgumentList {
		w.Walk(value, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitCase(w *Walker, node *ast.CaseStatement, metadata []Metadata) Metadata {
	w.Walk(node.Test, metadata)
	for _, e := range node.Consequent {
		w.Walk(e, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitCatch(w *Walker, node *ast.CatchStatement, metadata []Metadata) Metadata {
	w.Walk(node.Parameter, metadata)
	w.Walk(node.Body, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitConditional(w *Walker, node *ast.ConditionalExpression, metadata []Metadata) Metadata {
	w.Walk(node.Test, metadata)
	w.Walk(node.Consequent, metadata)
	w.Walk(node.Alternate, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitDebugger(w *Walker, node *ast.DebuggerStatement, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitDot(w *Walker, node *ast.DotExpression, metadata []Metadata) Metadata {
	w.Walk(node.Left, metadata)
	w.Walk(node.Identifier, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitDoWhile(w *Walker, node *ast.DoWhileStatement, metadata []Metadata) Metadata {
	w.Walk(node.Test, metadata)
	w.Walk(node.Body, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitEmpty(w *Walker, node *ast.EmptyExpression, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitEmptyStatement(w *Walker, node *ast.EmptyStatement, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitExpression(w *Walker, node *ast.ExpressionStatement, metadata []Metadata) Metadata {
	w.Walk(node.Expression, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitForIn(w *Walker, node *ast.ForInStatement, metadata []Metadata) Metadata {
	w.Walk(node.Into, metadata)
	w.Walk(node.Source, metadata)
	w.Walk(node.Body, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitFor(w *Walker, node *ast.ForStatement, metadata []Metadata) Metadata {
	w.Walk(node.Initializer, metadata)
	w.Walk(node.Test, metadata)
	w.Walk(node.Update, metadata)
	w.Walk(node.Body, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitFunction(w *Walker, node *ast.FunctionLiteral, metadata []Metadata) Metadata {
	w.Walk(node.Name, metadata)
	for _, value := range node.ParameterList.List {
		w.Walk(value, metadata)
	}
	w.Walk(node.Body, metadata)

	for _, value := range node.DeclarationList {
		switch value := value.(type) {
		case *ast.FunctionDeclaration:
			//w.Walk(value.Function, metadata)
		case *ast.VariableDeclaration:
			for _, value := range value.List {
				w.Walk(value, metadata)
			}
		default:
			panic(fmt.Errorf("Here be dragons: visit Function.declaration(%T)", value))
		}
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitFunctionStatement(w *Walker, node *ast.FunctionStatement, metadata []Metadata) Metadata {
	if node.Function != nil {
		w.Walk(node.Function, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitIdentifier(w *Walker, node *ast.Identifier, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitIf(w *Walker, node *ast.IfStatement, metadata []Metadata) Metadata {
	w.Walk(node.Test, metadata)
	w.Walk(node.Consequent, metadata)
	w.Walk(node.Alternate, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitLabelled(w *Walker, node *ast.LabelledStatement, metadata []Metadata) Metadata {
	w.Walk(node.Label, metadata)
	w.Walk(node.Statement, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitNew(w *Walker, node *ast.NewExpression, metadata []Metadata) Metadata {
	w.Walk(node.Callee, metadata)
	for _, e := range node.ArgumentList {
		w.Walk(e, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitNull(w *Walker, node *ast.NullLiteral, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitNumber(w *Walker, node *ast.NumberLiteral, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitObject(w *Walker, node *ast.ObjectLiteral, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitReturn(w *Walker, node *ast.ReturnStatement, metadata []Metadata) Metadata {
	w.Walk(node.Argument, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitRegex(w *Walker, node *ast.RegExpLiteral, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitSequence(w *Walker, node *ast.SequenceExpression, metadata []Metadata) Metadata {
	for _, e := range node.Sequence {
		w.Walk(e, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitString(w *Walker, node *ast.StringLiteral, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitSwitch(w *Walker, node *ast.SwitchStatement, metadata []Metadata) Metadata {
	w.Walk(node.Discriminant, metadata)
	for _, e := range node.Body {
		w.Walk(e, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitThis(w *Walker, node *ast.ThisExpression, metadata []Metadata) Metadata {
	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitThrow(w *Walker, node *ast.ThrowStatement, metadata []Metadata) Metadata {
	w.Walk(node.Argument, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitTry(w *Walker, node *ast.TryStatement, metadata []Metadata) Metadata {
	w.Walk(node.Body, metadata)
	w.Walk(node.Catch, metadata)
	w.Walk(node.Finally, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitUnary(w *Walker, node *ast.UnaryExpression, metadata []Metadata) Metadata {
	w.Walk(node.Operand, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitVariable(w *Walker, node *ast.VariableExpression, metadata []Metadata) Metadata {
	w.Walk(node.Initializer, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitVariableStatement(w *Walker, node *ast.VariableStatement, metadata []Metadata) Metadata {
	for _, e := range node.List {
		w.Walk(e, metadata)
	}

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitWhile(w *Walker, node *ast.WhileStatement, metadata []Metadata) Metadata {
	w.Walk(node.Test, metadata)
	w.Walk(node.Body, metadata)

	return CurrentMetadata(metadata)
}

func (v *VisitorImpl) VisitWith(w *Walker, node *ast.WithStatement, metadata []Metadata) Metadata {
	w.Walk(node.Object, metadata)
	w.Walk(node.Body, metadata)

	return CurrentMetadata(metadata)
}
