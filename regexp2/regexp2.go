// Package regexp2 provides ECMAScript compatible regexp.Finder and regexp.Creator.
//
// It only provides a subset of the standard library methods needed to satisfy regexp.Creator.
package regexp2

import (
	"github.com/dlclark/regexp2"
	"github.com/robertkrimen/otto/regexp"
)

// Validate we match the required interfaces.
var (
	_ regexp.Finder  = &Regexp{}
	_ regexp.Creator = &Creator{}
)

// DefaultOption is the default option used when creating new Regexp via Compile or MustCompile.
var DefaultOption regexp2.RegexOptions = regexp2.ECMAScript

// Regexp is regexp compatibility shim for regexp2 in ECMAScript mode.
type Regexp struct {
	re *regexp2.Regexp
}

// FindStringIndex returns a two-element slice of integers defining the location of the
// leftmost match in s of the regular expression. The match itself is at s[loc[0]:loc[1]].
// A return value of nil indicates no match.
func (re *Regexp) FindStringIndex(s string) []int {
	match, err := re.re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	return re.matchIndex(match)
}

// FindAllStringIndex is the 'All' version of FindStringIndex; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description in the package
// comment. A return value of nil indicates no match.
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int {
	if n == 0 {
		return nil
	}

	match, err := re.re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	var res [][]int
	for match != nil {
		res = append(res, re.matchIndex(match))
		if len(res) == n {
			return res
		}

		match, err = re.re.FindNextMatch(match)
		if err != nil {
			panic(err)
		}
	}

	return res
}

// FindAllSubmatchIndex is the 'All' version of FindSubmatchIndex;
// it returns a slice of all successive matches of the expression,
// as defined by the 'All' description in the package comment. A
// return value of nil indicates no match.
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int {
	return re.FindAllStringSubmatchIndex(string(b), n)
}

// FindAllStringSubmatchIndex is the 'All' version of FindStringSubmatchIndex;
// it returns a slice of all successive matches of the expression, as defined
// by the 'All' description in the package comment. A return value of nil
// indicates no match.
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int {
	if n == 0 {
		return nil
	}
	match, err := re.re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	var res [][]int
	for match != nil {
		res = append(res, re.submatchIndex(match))
		if len(res) == n {
			return res
		}

		match, err = re.re.FindNextMatch(match)
		if err != nil {
			panic(err)
		}
	}

	return res
}

// FindStringSubmatchIndex returns a slice holding the index pairs identifying the leftmost
// match of the regular expression in s and the matches, if any, of its sub expressions, as
// defined by the 'Submatch' and 'Index' descriptions in the package comment. A return value
// of nil indicates no match.
func (re *Regexp) FindStringSubmatchIndex(s string) []int {
	match, err := re.re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	return re.submatchIndex(match)
}

func (re *Regexp) matchIndex(match *regexp2.Match) []int {
	if match == nil || len(match.Captures) == 0 {
		return nil
	}

	res := make([]int, 0, len(match.Captures)*2)
	for _, c := range match.Captures {
		res = append(res, c.Index, c.Index+c.Length)
	}
	return res
}

func (re *Regexp) submatchIndex(match *regexp2.Match) []int {
	if match == nil || match.GroupCount() == 0 {
		return nil
	}

	var res []int
	for _, g := range match.Groups() {
		for _, c := range g.Captures {
			res = append(res, c.Index, c.Index+c.Length)
		}
	}

	return res
}

// Compile parses a regular expression and returns, if successful,
// a Regexp object that can be used to match against text.
func Compile(expr string) (*Regexp, error) {
	re, err := regexp2.Compile(expr, DefaultOption)
	if err != nil {
		return nil, err
	}
	return &Regexp{re: re}, nil
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular expressions.
func MustCompile(expr string) *Regexp {
	re, err := regexp2.Compile(expr, DefaultOption)
	if err != nil {
		panic(err)
	}
	return &Regexp{re: re}
}

// Creator implements regexp.Creator.
type Creator struct{}

// Compile implements regexp.Creator.
func (Creator) Compile(expr string) (regexp.Finder, error) {
	return Compile(expr)
}

// MustCompile implements regexp.Creator.
func (Creator) MustCompile(expr string) regexp.Finder {
	return MustCompile(expr)
}

// NoTransform implements regexp.Creator.
func (Creator) NoTransform() bool {
	return true
}
