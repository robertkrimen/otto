package walk

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"reflect"
)

// Metadata contains information about a node
type Metadata map[string]interface{}

func NewMetadata(parent ast.Node) Metadata {
	md := Metadata{"parent": parent}
	return md
}

func (md Metadata) Parent() ast.Node {
	parent, ok := md["parent"].(ast.Node)
	if !ok {
		return nil
	}

	return parent
}

func (md Metadata) AddParent(parent ast.Node) {
	md["parent"] = parent
}

// String displays information about the metadata
func (md Metadata) String() string {
	return fmt.Sprintf("{parent:%v}", reflect.TypeOf(md["parent"]))
}
