// Package underscore contains the source for the JavaScript utility-belt library.
//
//	import (
//		_ "github.com/robertkrimen/otto/underscore"
//	)
//
// Every Otto runtime will now include [underscore] for more information see the [underscore docs]
//
// By importing this package, you'll automatically load underscore every time you create a new Otto runtime.
//
// To prevent this behavior, you can do the following:
//
//	import (
//		"github.com/robertkrimen/otto/underscore"
//	)
//
//	func init() {
//		underscore.Disable()
//	}
//
// [underscore]: http://underscorejs.org
// [underscore docs]: https://github.com/documentcloud/underscore
package underscore

import (
	_ "embed"

	"github.com/robertkrimen/otto/registry"
)

//go:embed underscore-min.js
var underscore string
var entry *registry.Entry = registry.Register(Source)

// Enable underscore runtime inclusion.
func Enable() {
	entry.Enable()
}

// Disable underscore runtime inclusion.
func Disable() {
	entry.Disable()
}

// Source returns the underscore source.
func Source() string {
	return underscore
}
