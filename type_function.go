package otto

type _functionObject struct {
	call      _callFunction
	construct _constructFunction
}

func (runtime *_runtime) newNativeFunctionObject(native _nativeFunction, length int) *_object {
	self := runtime.newClassObject("Function")
	self.value = _functionObject{
		call:      newNativeCallFunction(native),
		construct: defaultConstructFunction,
	}
	self.defineProperty("length", toValue_int(length), 0000, false)
	return self
}

func (runtime *_runtime) newNodeFunctionObject(node *_functionNode, scopeEnvironment _environment) *_object {
	self := runtime.newClassObject("Function")
	self.value = _functionObject{
		call:      newNodeCallFunction(node, scopeEnvironment),
		construct: defaultConstructFunction,
	}
	self.defineProperty("length", toValue_int(len(node.ParameterList)), 0000, false)
	return self
}

func (runtime *_runtime) newBoundFunctionObject(target *_object, this Value, argumentList []Value) *_object {
	self := runtime.newClassObject("Function")
	self.value = _functionObject{
		call:      newBoundCallFunction(target, this, argumentList),
		construct: defaultConstructFunction,
	}
	// FIXME
	self.defineProperty("length", toValue_int(0), 0000, false)
	return self
}

func (self *_object) functionValue() _functionObject {
	value, _ := self.value.(_functionObject)
	return value
}

func (self *_object) Call(this Value, argumentList ...interface{}) Value {
	if self.functionValue().call == nil {
		panic(newTypeError("%v is not a function", toValue_object(self)))
	}
	return self.runtime.Call(self, this, self.runtime.toValueArray(argumentList...), false)
	// ... -> runtime -> self.Function.Call.Dispatch -> ...
}

func (self *_object) Construct(this Value, argumentList ...interface{}) Value {
	function := self.functionValue()
	if function.call == nil {
		panic(newTypeError("%v is not a function", toValue_object(self)))
	}
	if function.construct == nil {
		panic(newTypeError("%v is not a constructor", toValue_object(self)))
	}
	return function.construct(self, this, self.runtime.toValueArray(argumentList...))
}

func defaultConstructFunction(self *_object, this Value, argumentList []Value) Value {
	newObject := self.runtime.newObject()
	newObject.class = "Object"
	prototypeValue := self.get("prototype")
	if !prototypeValue.IsObject() {
		prototypeValue = toValue_object(self.runtime.Global.ObjectPrototype)
	}
	newObject.prototype = prototypeValue._object()
	newObjectValue := toValue_object(newObject)
	result := self.Call(newObjectValue, argumentList)
	if result.IsObject() {
		return result
	}
	return newObjectValue
}

func (self *_object) CallGet(name string) Value {
	return self.runtime.Call(self, toValue_object(self), []Value{toValue_string(name)}, false)
}

func (self *_object) CallSet(name string, value Value) {
	self.runtime.Call(self, toValue_object(self), []Value{toValue_string(name), value}, false)
}

// 15.3.5.3
func (self *_object) HasInstance(of Value) bool {
	if self.functionValue().call == nil {
		// We should not have a HasInstance method
		panic(newTypeError())
	}
	if !of.IsObject() {
		return false
	}
	prototype := self.get("prototype")
	if !prototype.IsObject() {
		panic(newTypeError())
	}
	prototypeObject := prototype._object()

	value := of._object().prototype
	for value != nil {
		if value == prototypeObject {
			return true
		}
		value = value.prototype
	}
	return false
}

type _nativeFunction func(FunctionCall) Value

// _constructFunction
type _constructFunction func(*_object, Value, []Value) Value

// _callFunction
type _callFunction interface {
	Dispatch(*_object, *_functionEnvironment, *_runtime, Value, []Value, bool) Value
	Source() string
	ScopeEnvironment() _environment
}

// _nativeCallFunction
type _nativeCallFunction _nativeFunction

func newNativeCallFunction(native _nativeFunction) _nativeCallFunction {
	return _nativeCallFunction(native)
}

func (self _nativeCallFunction) Dispatch(_ *_object, _ *_functionEnvironment, runtime *_runtime, this Value, argumentList []Value, evalHint bool) Value {
	return self(FunctionCall{
		runtime:  runtime,
		evalHint: evalHint,

		This:         this,
		ArgumentList: argumentList,
		Otto:         runtime.Otto,
	})
}

func (self _nativeCallFunction) ScopeEnvironment() _environment {
	return nil
}

func (self _nativeCallFunction) Source() string {
	return ""
}

// _nodeCallFunction
type _nodeCallFunction struct {
	node             *_functionNode
	scopeEnvironment _environment // Can be either Lexical or Variable
}

func newNodeCallFunction(node *_functionNode, scopeEnvironment _environment) *_nodeCallFunction {
	self := &_nodeCallFunction{
		node: node,
	}
	self.scopeEnvironment = scopeEnvironment
	return self
}

func (self _nodeCallFunction) Dispatch(function *_object, environment *_functionEnvironment, runtime *_runtime, this Value, argumentList []Value, _ bool) Value {
	return runtime._callNode(function, environment, self.node, this, argumentList)
}

func (self _nodeCallFunction) ScopeEnvironment() _environment {
	return self.scopeEnvironment
}

func (self _nodeCallFunction) Source() string {
	return ""
}

type _boundCallFunction struct {
	target       *_object
	this         Value
	argumentList []Value
}

func newBoundCallFunction(target *_object, this Value, argumentList []Value) *_boundCallFunction {
	self := &_boundCallFunction{
		target:       target,
		this:         this,
		argumentList: argumentList,
	}
	return self
}

func (self _boundCallFunction) Dispatch(_ *_object, _ *_functionEnvironment, runtime *_runtime, this Value, argumentList []Value, _ bool) Value {
	argumentList = append(self.argumentList, argumentList...)
	return runtime.Call(self.target, self.this, argumentList, false)
}

func (self _boundCallFunction) ScopeEnvironment() _environment {
	return nil
}

func (self _boundCallFunction) Source() string {
	return ""
}

// FunctionCall{}

// FunctionCall is an enscapulation of a JavaScript function call.
type FunctionCall struct {
	runtime     *_runtime
	_thisObject *_object
	evalHint    bool

	This         Value
	ArgumentList []Value
	Otto         *Otto
}

// Argument will return the value of the argument at the given index.
//
// If no such argument exists, undefined is returned.
func (self FunctionCall) Argument(index int) Value {
	return valueOfArrayIndex(self.ArgumentList, index)
}

func (self FunctionCall) slice(index int) []Value {
	if index < len(self.ArgumentList) {
		return self.ArgumentList[index:]
	}
	return []Value{}
}

func (self *FunctionCall) thisObject() *_object {
	if self._thisObject == nil {
		this := self.runtime.GetValue(self.This) // FIXME Is this right?
		self._thisObject = self.runtime.toObject(this)
	}
	return self._thisObject
}

func (self *FunctionCall) thisClassObject(class string) *_object {
	thisObject := self.thisObject()
	if thisObject.class != class {
		panic(newTypeError())
	}
	return self._thisObject
}

func (self FunctionCall) toObject(value Value) *_object {
	return self.runtime.toObject(value)
}
