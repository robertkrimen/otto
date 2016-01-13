package walk

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"reflect"
)

// Metadata contains information about a node
type Metadata struct {
	Parent ast.Node
}

// String displays information about the metadata
func (md Metadata) String() string {
	return fmt.Sprintf("{parent:%v}", reflect.TypeOf(md.Parent))
}
