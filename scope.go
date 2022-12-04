package otto

// An ECMA-262 ExecutionContext.
type scope struct {
	lexical  stasher
	variable stasher
	this     *object
	eval     bool // Replace this with kind?
	outer    *scope
	depth    int

	frame frame
}

func newScope(lexical stasher, variable stasher, this *object) *scope {
	return &scope{
		lexical:  lexical,
		variable: variable,
		this:     this,
	}
}
