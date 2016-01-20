package walk

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"reflect"
)

const (
	Vars string = "vars"
)

// Metadata contains information about a node.
// It is a map of values, by default the parent of the current node is inserted.
type Metadata map[string]interface{}

// NewMetadata returns a new instance
func NewMetadata(parent ast.Node) Metadata {
	md := Metadata{"parent": parent}
	return md
}

// Parent retrieves the parent of the node
func (md Metadata) Parent() ast.Node {
	parent, ok := md["parent"].(ast.Node)
	if !ok {
		return nil
	}

	return parent
}

// AddParent inserts the given node as the parent
func (md Metadata) AddParent(parent ast.Node) {
	md["parent"] = parent
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
	return fmt.Sprintf("{parent:%v}", reflect.TypeOf(md["parent"]))
}

type Variables map[string]file.Idx

func NewVariables() Variables {
	return Variables{}
}
