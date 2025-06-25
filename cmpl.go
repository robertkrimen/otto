package otto

import (
	"github.com/kubeshark/otto/ast"
	"github.com/kubeshark/otto/file"
)

type compiler struct {
	file    *file.File
	program *ast.Program
}
