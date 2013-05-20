package otto

import (
	"math"
	"strings"
)

func (self *_runtime) evaluateConditional(node *_conditionalNode) Value {
	test := self.evaluate(node.Test)
	testValue := self.GetValue(test)
	if toBoolean(testValue) {
		return self.evaluate(node.Consequent)
	}
	return self.evaluate(node.Alternate)
}

func (self *_runtime) evaluateNew(node *_newNode) Value {
	callee := self.evaluate(node.Callee)
	calleeValue := self.GetValue(callee)
	argumentList := []Value{}
	for _, argumentNode := range node.ArgumentList {
		argumentList = append(argumentList, self.GetValue(self.evaluate(argumentNode)))
	}
	this := UndefinedValue()
	if !calleeValue.IsFunction() {
		panic(newTypeError("%v is not a function", calleeValue))
	}
	return calleeValue._object().Construct(this, argumentList)
}

func (self *_runtime) evaluateArray(node *_arrayNode) Value {

	valueArray := []Value{}

	for _, node := range node.nodeList {
		valueArray = append(valueArray, self.GetValue(self.evaluate(node)))
	}

	result := self.newArrayOf(valueArray)

	return toValue(result)
}

func (self *_runtime) evaluateObject(node *_objectNode) Value {

	result := self.newObject()
	descriptors := map[string]_property{}
	descrOrder := []string{}

	for _, property := range node.propertyList {
		value := self.GetValue(self.evaluate(property.Value))
		if property.isdescriptor {
			key, value := property.Key, toPropertyDescriptor(value)
			if descriptor, exists := descriptors[key]; exists {
				if current, isgetset := descriptor.value.(_propertyGetSet); isgetset {
					if next, isgetset := value.value.(_propertyGetSet); isgetset {
						if nil != next[0] {
							current[0] = next[0]
						}
						if nil != next[1] {
							current[1] = next[1]
						}
						descriptor.value = current
						descriptors[key] = descriptor
					}
				}
			} else {
				descrOrder = append(descrOrder, key)
				descriptors[key] = value
			}
		} else {
			result.defineProperty(property.Key, self.GetValue(self.evaluate(property.Value)), 0111, false)
		}
	}

	for _, key := range descrOrder {
		if descriptor, exists := descriptors[key]; exists {
			result.defineOwnProperty(key, descriptor, false)
		}
	}

	return toValue(result)
}

func (self *_runtime) evaluateRegExp(node *_regExpNode) Value {
	return toValue(self._newRegExp(node.Pattern, node.Flags))
}

func (self *_runtime) evaluateUnaryOperation(node *_unaryOperationNode) Value {

	target := self.evaluate(node.Target)
	switch node.Operator {
	case "typeof", "delete":
		if target._valueType == valueReference && target.reference().IsUnresolvable() {
			if node.Operator == "typeof" {
				return toValue("undefined")
			}
			return TrueValue()
		}
	}

	targetValue := self.GetValue(target)

	switch node.Operator {
	case "!":
		if targetValue.toBoolean() {
			return FalseValue()
		}
		return TrueValue()
	case "~":
		integerValue := toInt32(targetValue)
		return toValue(^integerValue)
	case "+":
		return toValue(targetValue.toFloat())
	case "-":
		value := targetValue.toFloat()
		// TODO Test this
		sign := float64(-1)
		if math.Signbit(value) {
			sign = 1
		}
		return toValue(math.Copysign(value, sign))
	case "++=": // Prefix ++
		newValue := toValue(+1 + targetValue.toFloat())
		self.PutValue(target.reference(), newValue)
		return newValue
	case "--=": // Prefix --
		newValue := toValue(-1 + targetValue.toFloat())
		self.PutValue(target.reference(), newValue)
		return newValue
	case "=++": // Postfix ++
		oldValue := targetValue.toFloat()
		newValue := toValue(+1 + oldValue)
		self.PutValue(target.reference(), newValue)
		return toValue(oldValue)
	case "=--": // Postfix --
		oldValue := targetValue.toFloat()
		newValue := toValue(-1 + oldValue)
		self.PutValue(target.reference(), newValue)
		return toValue(oldValue)
	case "void":
		return UndefinedValue()
	case "delete":
		reference := target.reference()
		if reference == nil {
			return TrueValue()
		}
		return toValue(target.reference().Delete())
	case "typeof":
		switch targetValue._valueType {
		case valueUndefined:
			return toValue("undefined")
		case valueNull:
			return toValue("object")
		case valueBoolean:
			return toValue("boolean")
		case valueNumber:
			return toValue("number")
		case valueString:
			return toValue("string")
		case valueObject:
			if targetValue._object().functionValue() != nil {
				return toValue("function")
			}
			return toValue("object")
		default:
			// ?
		}
	}

	panic(hereBeDragons())
}

func (self *_runtime) evaluateMultiply(left float64, right float64) Value {
	// TODO 11.5.1
	return UndefinedValue()
}

func (self *_runtime) evaluateDivide(left float64, right float64) Value {
	if math.IsNaN(left) || math.IsNaN(right) {
		return NaNValue()
	}
	if math.IsInf(left, 0) && math.IsInf(right, 0) {
		return NaNValue()
	}
	if left == 0 && right == 0 {
		return NaNValue()
	}
	if math.IsInf(left, 0) {
		if math.Signbit(left) == math.Signbit(right) {
			return positiveInfinityValue()
		} else {
			return negativeInfinityValue()
		}
	}
	if math.IsInf(right, 0) {
		if math.Signbit(left) == math.Signbit(right) {
			return positiveZeroValue()
		} else {
			return negativeZeroValue()
		}
	}
	if right == 0 {
		if math.Signbit(left) == math.Signbit(right) {
			return positiveInfinityValue()
		} else {
			return negativeInfinityValue()
		}
	}
	return toValue(left / right)
}

func (self *_runtime) evaluateModulo(left float64, right float64) Value {
	// TODO 11.5.3
	return UndefinedValue()
}

func (self *_runtime) calculateBinaryOperation(operator string, left Value, right Value) Value {

	leftValue := self.GetValue(left)

	switch operator {

	// Additive
	case "+":
		leftValue = toPrimitive(leftValue)
		rightValue := self.GetValue(right)
		rightValue = toPrimitive(rightValue)

		if leftValue.IsString() || rightValue.IsString() {
			return toValue(strings.Join([]string{leftValue.toString(), rightValue.toString()}, ""))
		} else {
			return toValue(leftValue.toFloat() + rightValue.toFloat())
		}
	case "-":
		rightValue := self.GetValue(right)
		return toValue(leftValue.toFloat() - rightValue.toFloat())

	// Multiplicative
	case "*":
		rightValue := self.GetValue(right)
		return toValue(leftValue.toFloat() * rightValue.toFloat())
	case "/":
		rightValue := self.GetValue(right)
		return self.evaluateDivide(leftValue.toFloat(), rightValue.toFloat())
	case "%":
		rightValue := self.GetValue(right)
		return toValue(math.Mod(leftValue.toFloat(), rightValue.toFloat()))

	// Logical
	case "&&":
		left := toBoolean(leftValue)
		if !left {
			return FalseValue()
		}
		return toValue(toBoolean(self.GetValue(right)))
	case "||":
		left := toBoolean(leftValue)
		if left {
			return TrueValue()
		}
		return toValue(toBoolean(self.GetValue(right)))

	// Bitwise
	case "&":
		rightValue := self.GetValue(right)
		return toValue(toInt32(leftValue) & toInt32(rightValue))
	case "|":
		rightValue := self.GetValue(right)
		return toValue(toInt32(leftValue) | toInt32(rightValue))
	case "^":
		rightValue := self.GetValue(right)
		return toValue(toInt32(leftValue) ^ toInt32(rightValue))

	// Shift
	// (Masking of 0x1f is to restrict the shift to a maximum of 31 places)
	case "<<":
		rightValue := self.GetValue(right)
		return toValue(toInt32(leftValue) << (toUint32(rightValue) & 0x1f))
	case ">>":
		rightValue := self.GetValue(right)
		return toValue(toInt32(leftValue) >> (toUint32(rightValue) & 0x1f))
	case ">>>":
		rightValue := self.GetValue(right)
		// Shifting an unsigned integer is a logical shift
		return toValue(toUint32(leftValue) >> (toUint32(rightValue) & 0x1f))

	case "instanceof":
		rightValue := self.GetValue(right)
		if !rightValue.IsObject() {
			panic(newTypeError("Expecting a function in instanceof check, but got: %v", rightValue))
		}
		return toValue(rightValue._object().HasInstance(leftValue))

	case "in":
		rightValue := self.GetValue(right)
		if !rightValue.IsObject() {
			panic(newTypeError())
		}
		return toValue(rightValue._object().hasProperty(toString(leftValue)))
	}

	panic(hereBeDragons(operator))
}

func (self *_runtime) evaluateAssignment(node *_assignmentNode) Value {

	left := self.evaluate(node.Left)
	right := self.evaluate(node.Right)
	rightValue := self.GetValue(right)

	result := rightValue
	if node.Operator != "" {
		result = self.calculateBinaryOperation(node.Operator, left, rightValue)
	}

	self.PutValue(left.reference(), result)

	return result
}

func valueKindDispatchKey(left _valueType, right _valueType) int {
	return (int(left) << 2) + int(right)
}

var equalDispatch map[int](func(Value, Value) bool) = makeEqualDispatch()

func makeEqualDispatch() map[int](func(Value, Value) bool) {
	key := valueKindDispatchKey
	return map[int](func(Value, Value) bool){

		key(valueNumber, valueObject): func(x Value, y Value) bool { return x.toFloat() == y.toFloat() },
		key(valueString, valueObject): func(x Value, y Value) bool { return x.toFloat() == y.toFloat() },
		key(valueObject, valueNumber): func(x Value, y Value) bool { return x.toFloat() == y.toFloat() },
		key(valueObject, valueString): func(x Value, y Value) bool { return x.toFloat() == y.toFloat() },
	}
}

type _lessThanResult int

const (
	lessThanFalse _lessThanResult = iota
	lessThanTrue
	lessThanUndefined
)

func calculateLessThan(left Value, right Value, leftFirst bool) _lessThanResult {

	x := UndefinedValue()
	y := x

	if leftFirst {
		x = toNumberPrimitive(left)
		y = toNumberPrimitive(right)
	} else {
		y = toNumberPrimitive(right)
		x = toNumberPrimitive(left)
	}

	result := false
	if x._valueType != valueString || y._valueType != valueString {
		x, y := x.toFloat(), y.toFloat()
		if math.IsNaN(x) || math.IsNaN(y) {
			return lessThanUndefined
		}
		result = x < y
	} else {
		x, y := x.toString(), y.toString()
		result = x < y
	}

	if result {
		return lessThanTrue
	}

	return lessThanFalse
}

var lessThanTable [4](map[_lessThanResult]bool) = [4](map[_lessThanResult]bool){
	// <
	map[_lessThanResult]bool{
		lessThanFalse:     false,
		lessThanTrue:      true,
		lessThanUndefined: false,
	},

	// >
	map[_lessThanResult]bool{
		lessThanFalse:     false,
		lessThanTrue:      true,
		lessThanUndefined: false,
	},

	// <=
	map[_lessThanResult]bool{
		lessThanFalse:     true,
		lessThanTrue:      false,
		lessThanUndefined: false,
	},

	// >=
	map[_lessThanResult]bool{
		lessThanFalse:     true,
		lessThanTrue:      false,
		lessThanUndefined: false,
	},
}

func (self *_runtime) calculateComparison(comparator string, left Value, right Value) bool {

	// TODO This might be redundant now (with regards to evaluateComparison)
	x := self.GetValue(left)
	y := self.GetValue(right)

	kindEqualKind := false
	result := true
	negate := false

	switch comparator {
	case "<":
		result = lessThanTable[0][calculateLessThan(x, y, true)]
	case ">":
		result = lessThanTable[1][calculateLessThan(y, x, false)]
	case "<=":
		result = lessThanTable[2][calculateLessThan(y, x, false)]
	case ">=":
		result = lessThanTable[3][calculateLessThan(x, y, true)]
	case "!==":
		negate = true
		fallthrough
	case "===":
		if x._valueType != y._valueType {
			result = false
		} else {
			kindEqualKind = true
		}
	case "!=":
		negate = true
		fallthrough
	case "==":
		if x._valueType == y._valueType {
			kindEqualKind = true
		} else if x._valueType <= valueUndefined && y._valueType <= valueUndefined {
			result = true
		} else if x._valueType <= valueUndefined || y._valueType <= valueUndefined {
			result = false
		} else if x._valueType <= valueString && y._valueType <= valueString {
			result = x.toFloat() == y.toFloat()
		} else if x._valueType == valueBoolean {
			result = self.calculateComparison("==", toValue(x.toFloat()), y)
		} else if y._valueType == valueBoolean {
			result = self.calculateComparison("==", x, toValue(y.toFloat()))
		} else if x._valueType == valueObject {
			result = self.calculateComparison("==", toPrimitive(x), y)
		} else if y._valueType == valueObject {
			result = self.calculateComparison("==", x, toPrimitive(y))
		} else {
			panic(hereBeDragons("Unable to test for equality: %v ==? %v", x, y))
		}
	}

	if kindEqualKind {
		switch x._valueType {
		case valueUndefined, valueNull:
			result = true
		case valueNumber:
			x := x.toFloat()
			y := y.toFloat()
			if math.IsNaN(x) || math.IsNaN(y) {
				result = false
			} else {
				result = x == y
			}
		case valueString:
			result = x.toString() == y.toString()
		case valueBoolean:
			result = x.toBoolean() == y.toBoolean()
		case valueObject:
			result = x._object() == y._object()
		default:
			goto ERROR
		}
	}

	if negate {
		result = !result
	}

	return result

ERROR:
	panic(hereBeDragons("%v (%v) %s %v (%v)", x, x._valueType, comparator, y, y._valueType))
}

func (self *_runtime) evaluateComparison(node *_comparisonNode) Value {

	left := self.GetValue(self.evaluate(node.Left))
	right := self.GetValue(self.evaluate(node.Right))

	return toValue(self.calculateComparison(node.Comparator, left, right))
}

func (self *_runtime) evaluateBinaryOperation(node *_binaryOperationNode) Value {

	left := self.evaluate(node.Left)
	leftValue := self.GetValue(left)

	switch node.Operator {
	// Logical
	case "&&":
		if !toBoolean(leftValue) {
			return leftValue
		}
		right := self.evaluate(node.Right)
		return self.GetValue(right)
	case "||":
		if toBoolean(leftValue) {
			return leftValue
		}
		right := self.evaluate(node.Right)
		return self.GetValue(right)
	}

	return self.calculateBinaryOperation(node.Operator, leftValue, self.evaluate(node.Right))
}

func (self *_runtime) evaluateCall(node *_callNode, withArgumentList []interface{}) Value {
	callee := self.evaluate(node.Callee)
	calleeValue := self.GetValue(callee)
	argumentList := []Value{}
	if withArgumentList != nil {
		argumentList = self.toValueArray(withArgumentList...)
	} else {
		for _, argumentNode := range node.ArgumentList {
			argumentList = append(argumentList, self.GetValue(self.evaluate(argumentNode)))
		}
	}
	this := UndefinedValue()
	calleeReference := callee.reference()
	evalHint := false
	if calleeReference != nil {
		if calleeReference.IsPropertyReference() {
			calleeObject := calleeReference.GetBase().(*_object)
			this = toValue(calleeObject)
		} else {
			// TODO ImplictThisValue
		}
		if calleeReference.GetName() == "eval" {
			evalHint = true // Possible direct eval
		}
	}
	if !calleeValue.IsFunction() {
		panic(newTypeError("%v is not a function", calleeValue))
	}
	return self.Call(calleeValue._object(), this, argumentList, evalHint)
}

func (self *_runtime) evaluateFunction(node *_functionNode) Value {
	return toValue(self.newNodeFunction(node, self.LexicalEnvironment()))
}

func (self *_runtime) evaluateDotMember(node *_dotMemberNode) Value {
	target := self.evaluate(node.Target)
	targetValue := self.GetValue(target)
	// TODO Pass in base value as-is, and defer toObject till later?
	return toValue(newPropertyReference(self.toObject(targetValue), node.Member, false, node))
}

func (self *_runtime) evaluateBracketMember(node *_bracketMemberNode) Value {
	target := self.evaluate(node.Target)
	targetValue := self.GetValue(target)
	member := self.evaluate(node.Member)
	memberValue := self.GetValue(member)

	// TODO Pass in base value as-is, and defer toObject till later?
	return toValue(newPropertyReference(self.toObject(targetValue), toString(memberValue), false, node))
}

func (self *_runtime) evaluateIdentifier(node *_identifierNode) Value {
	name := node.Value
	// TODO Should be true or false (strictness) depending on context
	// getIdentifierReference should not return nil, but we check anyway and panic
	// so as not to propagate the nil into something else
	reference := getIdentifierReference(self.LexicalEnvironment(), name, false, node)
	if reference == nil {
		// Should never get here!
		panic(hereBeDragons("referenceError == nil: " + name))
	}
	return toValue(reference)
}

func (self *_runtime) evaluateValue(node *_valueNode) Value {
	return node.Value
}

func (self *_runtime) evaluateComma(node *_commaNode) Value {
	var result Value
	for _, node := range node.Sequence {
		result = self.evaluate(node)
		result = self.GetValue(result)
	}
	return result
}
