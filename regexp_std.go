package otto

import (
	restd "regexp"

	"github.com/robertkrimen/otto/regexp"
)

// regexpStd implements regexp.Creator using the standard regexp library.
type regexpStd struct{}

// Compile implements regexp.Creator.
func (regexpStd) Compile(expr string) (regexp.Finder, error) {
	return restd.Compile(expr)
}

// MustCompile implements regexp.Creator.
func (regexpStd) MustCompile(expr string) regexp.Finder {
	return restd.MustCompile(expr)
}

// NoTransform implements regexp.Creator.
func (regexpStd) NoTransform() bool {
	return false
}
