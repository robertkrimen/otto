package walk

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"reflect"
)

const (
	Vars      string = "vars"
	NodeField        = "node"
)

// Metadata contains information about a node.
// It is a map of values, by default the parent of the current node is inserted.
type Metadata map[string]interface{}

// NewMetadata returns a new instance
func NewMetadata(node ast.Node) Metadata {
	md := Metadata{NodeField: node}
	return md
}

// Parent retrieves the parent of the node
func (md Metadata) Node() ast.Node {
	parent, ok := md[NodeField].(ast.Node)
	if !ok {
		return nil
	}

	return parent
}

// AddParent inserts the given node as the parent
func (md Metadata) AddParent(parent ast.Node) {
	md[NodeField] = parent
}

// CurrentMetadata returns the last added element as the current metadata
func CurrentMetadata(metadata []Metadata) Metadata {
	l := len(metadata)
	if l == 0 {
		return nil
	}

	return metadata[l-1]
}

// ParentMetadata returns the second last added element as the parent metadata
func ParentMetadata(metadata []Metadata) Metadata {
	l := len(metadata)
	if l < 2 {
		return nil
	}

	return metadata[l-2]
}

// String displays information about the metadata
func (md Metadata) String() string {
	return fmt.Sprintf("{node:%v}", reflect.TypeOf(md[NodeField]))
}

type Variables map[string]file.Idx

func NewVariables() Variables {
	return Variables{}
}
