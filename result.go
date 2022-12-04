package otto

type resultKind int

const (
	_ resultKind = iota
	resultReturn
	resultBreak
	resultContinue
)

type result struct {
	kind   resultKind
	value  Value
	target string
}

func newReturnResult(value Value) result {
	return result{resultReturn, value, ""}
}

func newContinueResult(target string) result {
	return result{resultContinue, emptyValue, target}
}

func newBreakResult(target string) result {
	return result{resultBreak, emptyValue, target}
}
