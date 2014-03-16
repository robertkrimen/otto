package otto

import (
	"errors"
	"fmt"
)

type _exception struct {
	value interface{}
}

func newException(value interface{}) *_exception {
	return &_exception{
		value: value,
	}
}

func (self *_exception) eject() interface{} {
	value := self.value
	self.value = nil // Prevent Go from holding on to the value, whatever it is
	return value
}

type _error struct {
	Name    string
	Message string

	Line int // Hackish -- line where the error/exception occurred
}

var messageDetail map[string]string = map[string]string{
	"notDefined": "%v is not defined",
}

func messageFromDescription(description string, argumentList ...interface{}) string {
	message := messageDetail[description]
	if message == "" {
		message = description
	}
	message = fmt.Sprintf(message, argumentList...)
	return message
}

func (self _error) MessageValue() Value {
	if self.Message == "" {
		return UndefinedValue()
	}
	return toValue_string(self.Message)
}

func (self _error) String() string {
	if len(self.Name) == 0 {
		return self.Message
	}
	if len(self.Message) == 0 {
		return self.Name
	}
	return fmt.Sprintf("%s: %s", self.Name, self.Message)
}

func newError(name string, argumentList ...interface{}) _error {
	description := ""
	var node _node = nil
	length := len(argumentList)
	if length > 0 {
		if node, _ = argumentList[length-1].(_node); node != nil || argumentList[length-1] == nil {
			argumentList = argumentList[0 : length-1]
			length -= 1
		}
		if length > 0 {
			description, argumentList = argumentList[0].(string), argumentList[1:]
		}
	}
	error := _error{
		Name:    name,
		Message: messageFromDescription(description, argumentList...),
		Line:    -1,
	}
	if node != nil {
		error.Line = node.position()
	}
	return error
}

func newReferenceError(argumentList ...interface{}) _error {
	return newError("ReferenceError", argumentList...)
}

func newTypeError(argumentList ...interface{}) _error {
	return newError("TypeError", argumentList...)
}

func newRangeError(argumentList ...interface{}) _error {
	return newError("RangeError", argumentList...)
}

func newSyntaxError(argumentList ...interface{}) _error {
	return newError("SyntaxError", argumentList...)
}

func newURIError(argumentList ...interface{}) _error {
	return newError("URIError", argumentList...)
}

func newRetValError(argumentList ...interface{}) _error {
	return newError("RetValError", argumentList...)
}

func typeErrorResult(throw bool) bool {
	if throw {
		panic(newTypeError())
	}
	return false
}

func catchPanic(function func()) (err error) {
	defer func() {
		if caught := recover(); caught != nil {
			if exception, ok := caught.(*_exception); ok {
				caught = exception.eject()
			}
			switch caught := caught.(type) {
			case *_syntaxError:
				err = errors.New(fmt.Sprintf("%s (line %d)", caught.String(), caught.Line+0))
				return
			case _error:
				if caught.Line == -1 {
					err = errors.New(caught.String())
				} else {
					// We're 0-based (for now), hence the + 1
					err = errors.New(fmt.Sprintf("%s (line %d)", caught.String(), caught.Line+1))
				}
				return
			case Value:
				err = errors.New(toString(caught))
				return
			}
			panic(caught)
		}
	}()
	function()
	return nil
}

// SyntaxError

type _syntaxError struct {
	Message   string
	Line      int
	Column    int
	Character int
}

func (self _syntaxError) String() string {
	name := "SyntaxError"
	if len(self.Message) == 0 {
		return name
	}
	return fmt.Sprintf("%s: %s", name, self.Message)
}

func (self _token) newSyntaxError(description string, argumentList ...interface{}) *_syntaxError {
	return &_syntaxError{
		Message:   messageFromDescription(description, argumentList...),
		Line:      self.Line,
		Column:    self.Column,
		Character: self.Character,
	}
}
