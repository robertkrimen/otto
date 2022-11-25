package otto

import (
	"strconv"
	"unicode/utf16"
	"unicode/utf8"
)

type _stringObject interface {
	Length() int
	At(int) rune
	String() string
}

type _stringASCII string

func (str _stringASCII) Length() int {
	return len(str)
}

func (str _stringASCII) At(at int) rune {
	return rune(str[at])
}

func (str _stringASCII) String() string {
	return string(str)
}

type _stringWide struct {
	string  string
	value16 []uint16
}

func (str _stringWide) Length() int {
	if str.value16 == nil {
		str.value16 = utf16.Encode([]rune(str.string))
	}
	return len(str.value16)
}

func (str _stringWide) At(at int) rune {
	if str.value16 == nil {
		str.value16 = utf16.Encode([]rune(str.string))
	}
	return rune(str.value16[at])
}

func (str _stringWide) String() string {
	return str.string
}

func _newStringObject(str string) _stringObject {
	for i := 0; i < len(str); i++ {
		if str[i] >= utf8.RuneSelf {
			goto wide
		}
	}

	return _stringASCII(str)

wide:
	return &_stringWide{
		string: str,
	}
}

func stringAt(str _stringObject, index int) rune {
	if 0 <= index && index < str.Length() {
		return str.At(index)
	}
	return utf8.RuneError
}

func (runtime *_runtime) newStringObject(value Value) *_object {
	str := _newStringObject(value.string())

	self := runtime.newClassObject(classString)
	self.defineProperty(propertyLength, toValue_int(str.Length()), 0, false)
	self.objectClass = _classString
	self.value = str
	return self
}

func (self *_object) stringValue() _stringObject {
	if str, ok := self.value.(_stringObject); ok {
		return str
	}
	return nil
}

func stringEnumerate(self *_object, all bool, each func(string) bool) {
	if str := self.stringValue(); str != nil {
		length := str.Length()
		for index := 0; index < length; index++ {
			if !each(strconv.FormatInt(int64(index), 10)) {
				return
			}
		}
	}
	objectEnumerate(self, all, each)
}

func stringGetOwnProperty(self *_object, name string) *_property {
	if property := objectGetOwnProperty(self, name); property != nil {
		return property
	}
	// TODO Test a string of length >= +int32 + 1?
	if index := stringToArrayIndex(name); index >= 0 {
		if chr := stringAt(self.stringValue(), int(index)); chr != utf8.RuneError {
			return &_property{toValue_string(string(chr)), 0}
		}
	}
	return nil
}
