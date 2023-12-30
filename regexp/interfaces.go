// Package regexp defines Regular Expression interfaces used for otto.
package regexp

// Finder is implemented by types which can act as a regular expression finder.
type Finder interface {
	// FindStringIndex returns a two-element slice of integers defining the location of the
	// leftmost match in s of the regular expression. The match itself is at s[loc[0]:loc[1]].
	// A return value of nil indicates no match.
	FindStringIndex(s string) []int

	// FindAllStringIndex is the 'All' version of FindStringIndex; it returns a slice of all
	// successive matches of the expression, as defined by the 'All' description in the package
	// comment. A return value of nil indicates no match.
	FindAllStringIndex(s string, n int) [][]int

	// FindStringSubmatchIndex returns a slice holding the index pairs identifying the leftmost
	// match of the regular expression in s and the matches, if any, of its sub expressions, as
	// defined by the 'Submatch' and 'Index' descriptions in the package comment. A return value
	// of nil indicates no match.
	FindStringSubmatchIndex(s string) []int

	// FindAllStringSubmatchIndex is the 'All' version of FindStringSubmatchIndex; it returns a
	// slice of all successive matches of the expression, as defined by the 'All' description in
	// the package comment. A return value of nil indicates no match.
	FindAllStringSubmatchIndex(s string, n int) [][]int

	// FindAllSubmatchIndex is the 'All' version of FindSubmatchIndex; it returns a slice of all
	// successive matches of the expression, as defined by the 'All' description in the package
	// comment. A return value of nil indicates no match.
	FindAllSubmatchIndex(b []byte, n int) [][]int
}

// Creator is implemented by types which can operate a RegExp provider for otto.
type Creator interface {
	// Compile parses a regular expression and returns, if successful, a Finder that
	// can be used to match against text.
	Compile(str string) (Finder, error)

	// MustCompile is like Compile but panics if the expression cannot be parsed. It simplifies safe
	// initialization of global variables holding compiled regular expressions.
	MustCompile(str string) Finder

	// NoTransform returns true if no RegExp transformation should be done.
	NoTransform() bool
}
