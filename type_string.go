package otto

import (
	"strconv"

	"code.google.com/p/go.exp/utf8string"
)

type _stringObject struct {
	value Value
	utf8string.String
}

func (runtime *_runtime) newStringObject(value Value) *_object {
	s := toString(value)
	obj := _stringObject{value: toValue_string(s)}
	obj.Init(s)

	self := runtime.newClassObject("String")
	self.defineProperty("length", toValue_int(obj.RuneCount()), 0, false)
	self.objectClass = _classString
	self.value = obj
	return self
}

func (self *_object) stringValue() (string, _stringObject) {
	value, valid := self.value.(_stringObject)
	if valid {
		return value.String.String(), value
	}
	return "", _stringObject{}
}

func stringEnumerate(self *_object, all bool, each func(string) bool) {
	_, so := self.stringValue()
	for index := 0; index < so.RuneCount(); index++ {
		if !each(strconv.FormatInt(int64(index), 10)) {
			return
		}
	}
	objectEnumerate(self, all, each)
}

func stringGetOwnProperty(self *_object, name string) *_property {
	if property := objectGetOwnProperty(self, name); property != nil {
		return property
	}
	index := stringToArrayIndex(name)
	if i := int(index); i >= 0 {
		_, so := self.stringValue()
		if i < so.RuneCount() {
			return &_property{toValue_string(string(so.At(i))), 0}
		}
	}
	return nil
}
