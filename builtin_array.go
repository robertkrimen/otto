package otto

import (
	"strconv"
	"strings"
)

// Array

func builtinArray(call FunctionCall) Value {
	return toValue_object(builtinNewArrayNative(call.runtime, call.ArgumentList))
}

func builtinNewArray(self *_object, _ Value, argumentList []Value) Value {
	return toValue_object(builtinNewArrayNative(self.runtime, argumentList))
}

func builtinNewArrayNative(runtime *_runtime, argumentList []Value) *_object {
	if len(argumentList) == 1 {
		firstArgument := argumentList[0]
		if firstArgument.IsNumber() {
			return runtime.newArray(arrayUint32(firstArgument))
		}
	}
	return runtime.newArrayOf(argumentList)
}

func builtinArray_toString(call FunctionCall) Value {
	thisObject := call.thisObject()
	join := thisObject.get("join")
	if join.isCallable() {
		join := join._object()
		return join.Call(call.This, call.ArgumentList)
	}
	return builtinObject_toString(call)
}

func builtinArray_toLocaleString(call FunctionCall) Value {
	separator := ","
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))
	if length == 0 {
		return toValue_string("")
	}
	stringList := make([]string, 0, length)
	for index := int64(0); index < length; index += 1 {
		value := thisObject.get(arrayIndexToString(index))
		stringValue := ""
		switch value._valueType {
		case valueEmpty, valueUndefined, valueNull:
		default:
			object := call.runtime.toObject(value)
			toLocaleString := object.get("toLocaleString")
			if !toLocaleString.isCallable() {
				panic(newTypeError())
			}
			stringValue = toLocaleString.call(toValue_object(object)).toString()
		}
		stringList = append(stringList, stringValue)
	}
	return toValue_string(strings.Join(stringList, separator))
}

func builtinArray_concat(call FunctionCall) Value {
	thisObject := call.thisObject()
	valueArray := []Value{}
	source := append([]Value{toValue_object(thisObject)}, call.ArgumentList...)
	for _, item := range source {
		switch item._valueType {
		case valueObject:
			object := item._object()
			if isArray(object) {
				length := toInteger(object.get("length")).value
				for index := int64(0); index < length; index += 1 {
					name := strconv.FormatInt(index, 10)
					if object.hasProperty(name) {
						valueArray = append(valueArray, object.get(name))
					} else {
						valueArray = append(valueArray, Value{})
					}
				}
				continue
			}
			fallthrough
		default:
			valueArray = append(valueArray, item)
		}
	}
	return toValue_object(call.runtime.newArrayOf(valueArray))
}

func builtinArray_shift(call FunctionCall) Value {
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))
	if 0 == length {
		thisObject.put("length", toValue_int64(0), true)
		return UndefinedValue()
	}
	first := thisObject.get("0")
	for index := int64(1); index < length; index++ {
		from := arrayIndexToString(index)
		to := arrayIndexToString(index - 1)
		if thisObject.hasProperty(from) {
			thisObject.put(to, thisObject.get(from), true)
		} else {
			thisObject.delete(to, true)
		}
	}
	thisObject.delete(arrayIndexToString(length-1), true)
	thisObject.put("length", toValue_int64(length-1), true)
	return first
}

func builtinArray_push(call FunctionCall) Value {
	thisObject := call.thisObject()
	itemList := call.ArgumentList
	index := int64(toUint32(thisObject.get("length")))
	for len(itemList) > 0 {
		thisObject.put(arrayIndexToString(index), itemList[0], true)
		itemList = itemList[1:]
		index += 1
	}
	length := toValue_int64(index)
	thisObject.put("length", length, true)
	return length
}

func builtinArray_pop(call FunctionCall) Value {
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))
	if 0 == length {
		thisObject.put("length", toValue_uint32(0), true)
		return UndefinedValue()
	}
	last := thisObject.get(arrayIndexToString(length - 1))
	thisObject.delete(arrayIndexToString(length-1), true)
	thisObject.put("length", toValue_int64(length-1), true)
	return last
}

func builtinArray_join(call FunctionCall) Value {
	separator := ","
	{
		argument := call.Argument(0)
		if argument.IsDefined() {
			separator = toString(argument)
		}
	}
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))
	if length == 0 {
		return toValue_string("")
	}
	stringList := make([]string, 0, length)
	for index := int64(0); index < length; index += 1 {
		value := thisObject.get(arrayIndexToString(index))
		stringValue := ""
		switch value._valueType {
		case valueEmpty, valueUndefined, valueNull:
		default:
			stringValue = toString(value)
		}
		stringList = append(stringList, stringValue)
	}
	return toValue_string(strings.Join(stringList, separator))
}

func builtinArray_splice(call FunctionCall) Value {
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))

	start := valueToRangeIndex(call.Argument(0), length, false)
	deleteCount := valueToRangeIndex(call.Argument(1), int64(length)-start, true)
	valueArray := make([]Value, deleteCount)

	for index := int64(0); index < deleteCount; index++ {
		indexString := arrayIndexToString(int64(start + index))
		if thisObject.hasProperty(indexString) {
			valueArray[index] = thisObject.get(indexString)
		}
	}

	// 0, <1, 2, 3, 4>, 5, 6, 7
	// a, b
	// length 8 - delete 4 @ start 1

	itemList := []Value{}
	itemCount := int64(len(call.ArgumentList))
	if itemCount > 2 {
		itemCount -= 2 // Less the first two arguments
		itemList = call.ArgumentList[2:]
	} else {
		itemCount = 0
	}
	if itemCount < deleteCount {
		// The Object/Array is shrinking
		stop := int64(length) - deleteCount
		// The new length of the Object/Array before
		// appending the itemList remainder
		// Stopping at the lower bound of the insertion:
		// Move an item from the after the deleted portion
		// to a position after the inserted portion
		for index := start; index < stop; index++ {
			from := arrayIndexToString(index + deleteCount) // Position just after deletion
			to := arrayIndexToString(index + itemCount)     // Position just after splice (insertion)
			if thisObject.hasProperty(from) {
				thisObject.put(to, thisObject.get(from), true)
			} else {
				thisObject.delete(to, true)
			}
		}
		// Delete off the end
		// We don't bother to delete below <stop + itemCount> (if any) since those
		// will be overwritten anyway
		for index := int64(length); index > (stop + itemCount); index-- {
			thisObject.delete(arrayIndexToString(index-1), true)
		}
	} else if itemCount > deleteCount {
		// The Object/Array is growing
		// The itemCount is greater than the deleteCount, so we do
		// not have to worry about overwriting what we should be moving
		// ---
		// Starting from the upper bound of the deletion:
		// Move an item from the after the deleted portion
		// to a position after the inserted portion
		for index := int64(length) - deleteCount; index > start; index-- {
			from := arrayIndexToString(index + deleteCount - 1)
			to := arrayIndexToString(index + itemCount - 1)
			if thisObject.hasProperty(from) {
				thisObject.put(to, thisObject.get(from), true)
			} else {
				thisObject.delete(to, true)
			}
		}
	}

	for index := int64(0); index < itemCount; index++ {
		thisObject.put(arrayIndexToString(index+start), itemList[index], true)
	}
	thisObject.put("length", toValue_int64(int64(length)+itemCount-deleteCount), true)

	return toValue_object(call.runtime.newArrayOf(valueArray))
}

func builtinArray_slice(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("length")))
	start, end := rangeStartEnd(call.ArgumentList, length, false)

	if start >= end {
		// Always an empty array
		return toValue_object(call.runtime.newArray(0))
	}
	sliceLength := end - start
	sliceValueArray := make([]Value, sliceLength)

	for index := int64(0); index < sliceLength; index++ {
		from := arrayIndexToString(index + start)
		if thisObject.hasProperty(from) {
			sliceValueArray[index] = thisObject.get(from)
		}
	}

	return toValue_object(call.runtime.newArrayOf(sliceValueArray))
}

func builtinArray_unshift(call FunctionCall) Value {
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))
	itemList := call.ArgumentList
	itemCount := int64(len(itemList))

	for index := length; index > 0; index-- {
		from := arrayIndexToString(index - 1)
		to := arrayIndexToString(index + itemCount - 1)
		if thisObject.hasProperty(from) {
			thisObject.put(to, thisObject.get(from), true)
		} else {
			thisObject.delete(to, true)
		}
	}

	for index := int64(0); index < itemCount; index++ {
		thisObject.put(arrayIndexToString(index), itemList[index], true)
	}

	newLength := toValue_int64(length + itemCount)
	thisObject.put("length", newLength, true)
	return newLength
}

func builtinArray_reverse(call FunctionCall) Value {
	thisObject := call.thisObject()
	length := int64(toUint32(thisObject.get("length")))

	lower := struct {
		name   string
		index  int64
		exists bool
	}{}
	upper := lower

	lower.index = 0
	middle := length / 2 // Division will floor

	for lower.index != middle {
		lower.name = arrayIndexToString(lower.index)
		upper.index = length - lower.index - 1
		upper.name = arrayIndexToString(upper.index)

		lower.exists = thisObject.hasProperty(lower.name)
		upper.exists = thisObject.hasProperty(upper.name)

		if lower.exists && upper.exists {
			lowerValue := thisObject.get(lower.name)
			upperValue := thisObject.get(upper.name)
			thisObject.put(lower.name, upperValue, true)
			thisObject.put(upper.name, lowerValue, true)
		} else if !lower.exists && upper.exists {
			value := thisObject.get(upper.name)
			thisObject.delete(upper.name, true)
			thisObject.put(lower.name, value, true)
		} else if lower.exists && !upper.exists {
			value := thisObject.get(lower.name)
			thisObject.delete(lower.name, true)
			thisObject.put(upper.name, value, true)
		} else {
			// Nothing happens.
		}

		lower.index += 1
	}

	return call.This
}

func sortCompare(thisObject *_object, index0, index1 uint, compare *_object) int {
	j := struct {
		name    string
		exists  bool
		defined bool
		value   string
	}{}
	k := j
	j.name = arrayIndexToString(int64(index0))
	j.exists = thisObject.hasProperty(j.name)
	k.name = arrayIndexToString(int64(index1))
	k.exists = thisObject.hasProperty(k.name)

	if !j.exists && !k.exists {
		return 0
	} else if !j.exists {
		return 1
	} else if !k.exists {
		return -1
	}

	x := thisObject.get(j.name)
	y := thisObject.get(k.name)
	j.defined = x.IsDefined()
	k.defined = y.IsDefined()

	if !j.defined && !k.defined {
		return 0
	} else if !j.defined {
		return 1
	} else if !k.defined {
		return -1
	}

	if compare == nil {
		j.value = toString(x)
		k.value = toString(y)

		if j.value == k.value {
			return 0
		} else if j.value < k.value {
			return -1
		}

		return 1
	}

	return int(toInt32(compare.Call(UndefinedValue(), []Value{x, y})))
}

func arraySortSwap(thisObject *_object, index0, index1 uint) {

	j := struct {
		name   string
		exists bool
	}{}
	k := j

	j.name = arrayIndexToString(int64(index0))
	j.exists = thisObject.hasProperty(j.name)
	k.name = arrayIndexToString(int64(index1))
	k.exists = thisObject.hasProperty(k.name)

	if j.exists && k.exists {
		jValue := thisObject.get(j.name)
		kValue := thisObject.get(k.name)
		thisObject.put(j.name, kValue, true)
		thisObject.put(k.name, jValue, true)
	} else if !j.exists && k.exists {
		value := thisObject.get(k.name)
		thisObject.delete(k.name, true)
		thisObject.put(j.name, value, true)
	} else if j.exists && !k.exists {
		value := thisObject.get(j.name)
		thisObject.delete(j.name, true)
		thisObject.put(k.name, value, true)
	} else {
		// Nothing happens.
	}
}

func arraySortQuickPartition(thisObject *_object, left, right, pivot uint, compare *_object) uint {
	arraySortSwap(thisObject, pivot, right) // Right is now the pivot value
	cursor := left
	for index := left; index < right; index++ {
		if sortCompare(thisObject, index, right, compare) == -1 { // Compare to the pivot value
			arraySortSwap(thisObject, index, cursor)
			cursor += 1
		}
	}
	arraySortSwap(thisObject, cursor, right)
	return cursor
}

func arraySortQuickSort(thisObject *_object, left, right uint, compare *_object) {
	if left < right {
		pivot := left + (right-left)/2
		pivot = arraySortQuickPartition(thisObject, left, right, pivot, compare)
		if pivot > 0 {
			arraySortQuickSort(thisObject, left, pivot-1, compare)
		}
		arraySortQuickSort(thisObject, pivot+1, right, compare)
	}
}

func builtinArray_sort(call FunctionCall) Value {
	thisObject := call.thisObject()
	length := uint(toUint32(thisObject.get("length")))
	compareValue := call.Argument(0)
	compare := compareValue._object()
	if compareValue.IsUndefined() {
	} else if !compareValue.isCallable() {
		panic(newTypeError())
	}
	if length > 1 {
		arraySortQuickSort(thisObject, 0, length-1, compare)
	}
	return call.This
}

func builtinArray_isArray(call FunctionCall) Value {
	return toValue_bool(isArray(call.Argument(0)._object()))
}

func builtinArray_indexOf(call FunctionCall) Value {
	thisObject, matchValue := call.thisObject(), call.Argument(0)
	if length := int64(toUint32(thisObject.get("length"))); length > 0 {
		index := int64(0)
		if len(call.ArgumentList) > 1 {
			index = toInteger(call.Argument(1)).value
		}
		if index < 0 {
			if index += length; index < 0 {
				index = 0
			}
		} else if index >= length {
			index = -1
		}
		for ; index >= 0 && index < length; index++ {
			name := arrayIndexToString(int64(index))
			if !thisObject.hasProperty(name) {
				continue
			}
			value := thisObject.get(name)
			if strictEqualityComparison(matchValue, value) {
				return toValue_uint32(uint32(index))
			}
		}
	}
	return toValue_int(-1)
}

func builtinArray_lastIndexOf(call FunctionCall) Value {
	thisObject, matchValue := call.thisObject(), call.Argument(0)
	length := int64(toUint32(thisObject.get("length")))
	index := length - 1
	if len(call.ArgumentList) > 1 {
		index = toInteger(call.Argument(1)).value
	}
	if 0 > index {
		index += length
	}
	if index > length {
		index = length - 1
	} else if 0 > index {
		return toValue_int(-1)
	}
	for ; index >= 0; index-- {
		name := arrayIndexToString(int64(index))
		if !thisObject.hasProperty(name) {
			continue
		}
		value := thisObject.get(name)
		if strictEqualityComparison(matchValue, value) {
			return toValue_uint32(uint32(index))
		}
	}
	return toValue_int(-1)
}

func builtinArray_every(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		callThis := call.Argument(1)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				if value := thisObject.get(key); iterator.call(callThis, value, toValue_int64(index), this).isTrue() {
					continue
				}
				return FalseValue()
			}
		}
		return TrueValue()
	}
	panic(newTypeError())
}

func builtinArray_some(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		callThis := call.Argument(1)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				if value := thisObject.get(key); iterator.call(callThis, value, toValue_int64(index), this).isTrue() {
					return TrueValue()
				}
			}
		}
		return FalseValue()
	}
	panic(newTypeError())
}

func builtinArray_forEach(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		callThis := call.Argument(1)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				iterator.call(callThis, thisObject.get(key), toValue_int64(index), this)
			}
		}
		return UndefinedValue()
	}
	panic(newTypeError())
}

func builtinArray_map(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		callThis := call.Argument(1)
		values := make([]Value, length)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				values[index] = iterator.call(callThis, thisObject.get(key), index, this)
			} else {
				values[index] = UndefinedValue()
			}
		}
		return toValue_object(call.runtime.newArrayOf(values))
	}
	panic(newTypeError())
}

func builtinArray_filter(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		callThis := call.Argument(1)
		values := make([]Value, 0)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				value := thisObject.get(key)
				if iterator.call(callThis, value, index, this).isTrue() {
					values = append(values, value)
				}
			}
		}
		return toValue_object(call.runtime.newArrayOf(values))
	}
	panic(newTypeError())
}

func builtinArray_reduce(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		start := call.Argument(1)
		length := int64(toUint32(thisObject.get("length")))
		index := int64(0)
		if length > 0 || start.IsDefined() {
			var accumulator Value
			if !start.IsDefined() {
				for ; index < length; index++ {
					if key := arrayIndexToString(index); thisObject.hasProperty(key) {
						accumulator = thisObject.get(key)
						index++
						break
					}
				}
			} else {
				accumulator = start
			}
			for ; index < length; index++ {
				if key := arrayIndexToString(index); thisObject.hasProperty(key) {
					accumulator = iterator.call(UndefinedValue(), accumulator, thisObject.get(key), key, this)
				}
			}
			return accumulator
		}
	}
	panic(newTypeError())
}

func builtinArray_reduceRight(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if iterator := call.Argument(0); iterator.isCallable() {
		start := call.Argument(1)
		length := int64(toUint32(thisObject.get("length")))
		if length > 0 || start.IsDefined() {
			index := length - 1
			var accumulator Value
			if !start.IsDefined() {
				for ; index >= 0; index-- {
					if key := arrayIndexToString(index); thisObject.hasProperty(key) {
						accumulator = thisObject.get(key)
						index -= 1
						break
					}
				}
			} else {
				accumulator = start
			}
			for ; index >= 0; index-- {
				if key := arrayIndexToString(index); thisObject.hasProperty(key) {
					accumulator = iterator.call(UndefinedValue(), accumulator, thisObject.get(key), key, this)
				}
			}
			return accumulator
		}
	}
	panic(newTypeError())
}
