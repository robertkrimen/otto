package otto

import (
	"strconv"
)

func (runtime *_runtime) newArrayObject(length uint32) *_object {
	self := runtime.newObject()
	self.class = "Array"
	self.defineProperty("length", toValue_uint32(length), 0100, false)
	self.objectClass = _classArray
	return self
}

func isArray(object *_object) bool {
	return object != nil && (object.class == "Array" || object.class == "GoArray")
}

func arrayDefineOwnProperty(self *_object, name string, descriptor _property, throw bool) bool {
	lengthProperty := self.getOwnProperty("length")
	lengthValue, valid := lengthProperty.value.(Value)
	if !valid {
		panic("Array.length != Value{}")
	}
	length := lengthValue.value.(uint32)
	if name == "length" {
		if descriptor.value == nil {
			return objectDefineOwnProperty(self, name, descriptor, throw)
		}
		newLength := uint32(0)
		{
			tmp := toInteger(descriptor.value.(Value))
			if !tmp.valid() || !isUint32(tmp.value) {
				panic(newRangeError())
			}
			newLength = uint32(tmp.value)
		}
		descriptor.value = toValue_uint32(newLength)
		if newLength > length {
			return objectDefineOwnProperty(self, name, descriptor, throw)
		}
		if !lengthProperty.writable() {
			goto Reject
		}
		newWritable := true
		if descriptor.mode&0700 == 0 {
			// If writable is off
			newWritable = false
			descriptor.mode |= 0100
		}
		if !objectDefineOwnProperty(self, name, descriptor, throw) {
			return false
		}
		for newLength < length {
			length -= 1
			if !self.delete(strconv.FormatInt(int64(length), 10), false) {
				descriptor.value = toValue_uint32(length + 1)
				if !newWritable {
					descriptor.mode &= 0077
				}
				objectDefineOwnProperty(self, name, descriptor, false)
				goto Reject
			}
		}
		if !newWritable {
			descriptor.mode &= 0077
			objectDefineOwnProperty(self, name, descriptor, false)
		}
	} else if index := stringToArrayIndex(name); index >= 0 {
		if index >= int64(length) && !lengthProperty.writable() {
			goto Reject
		}
		if !objectDefineOwnProperty(self, strconv.FormatInt(index, 10), descriptor, false) {
			goto Reject
		}
		if index >= int64(length) {
			lengthProperty.value = toValue_uint32(uint32(index + 1))
			objectDefineOwnProperty(self, "length", *lengthProperty, false)
			return true
		}
	}
	return objectDefineOwnProperty(self, name, descriptor, throw)
Reject:
	if throw {
		panic(newTypeError())
	}
	return false
}
