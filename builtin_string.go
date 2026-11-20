package otto

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// String

func stringValueFromStringArgumentList(argumentList []Value) Value {
	if len(argumentList) > 0 {
		return stringValue(argumentList[0].string())
	}
	return stringValue("")
}

func builtinString(call FunctionCall) Value {
	return stringValueFromStringArgumentList(call.ArgumentList)
}

func builtinNewString(obj *object, argumentList []Value) Value {
	return objectValue(obj.runtime.newString(stringValueFromStringArgumentList(argumentList)))
}

func builtinStringToString(call FunctionCall) Value {
	return call.thisClassObject(classStringName).primitiveValue()
}

func builtinStringValueOf(call FunctionCall) Value {
	return call.thisClassObject(classStringName).primitiveValue()
}

func builtinStringFromCharCode(call FunctionCall) Value {
	chrList := make([]uint16, len(call.ArgumentList))
	for index, value := range call.ArgumentList {
		chrList[index] = toUint16(value)
	}
	return string16Value(chrList)
}

func builtinStringCharAt(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	idx := int(call.Argument(0).number().int64)
	chr := stringAt(call.This.object().stringValue(), idx)
	if chr == utf8.RuneError {
		return stringValue("")
	}
	return stringValue(string(chr))
}

func builtinStringCharCodeAt(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	idx := int(call.Argument(0).number().int64)
	chr := stringAt(call.This.object().stringValue(), idx)
	if chr == utf8.RuneError {
		return NaNValue()
	}
	return uint16Value(uint16(chr))
}

func builtinStringConcat(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	var value bytes.Buffer
	value.WriteString(call.This.string())
	for _, item := range call.ArgumentList {
		value.WriteString(item.string())
	}
	return stringValue(value.String())
}

func lastIndexRune(s, substr string) int {
	if i := strings.LastIndex(s, substr); i >= 0 {
		return utf16Length(s[:i])
	}
	return -1
}

func indexRune(s, substr string) int {
	if i := strings.Index(s, substr); i >= 0 {
		return utf16Length(s[:i])
	}
	return -1
}

func utf16Length(s string) int {
	return len(utf16.Encode([]rune(s)))
}

// clampStringPosition converts an integer position to an index in [0, length].
func clampStringPosition(position float64, length int) int {
	switch {
	case position <= 0:
		return 0
	case position >= float64(length):
		return length
	default:
		return int(position)
	}
}

// isASCII reports whether s has only ASCII bytes, so byte and UTF-16 offsets agree.
func isASCII(s string) bool {
	for i := range len(s) {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// indexUTF16 returns the first index of target in value, or -1, in O(len(value)+len(target)).
func indexUTF16(value, target []uint16) int {
	if len(target) == 0 {
		return 0
	}
	// Knuth-Morris-Pratt: fail[i] is the length of the longest proper border of target[:i+1].
	fail := make([]int, len(target))
	for i, k := 1, 0; i < len(target); i++ {
		for k > 0 && target[i] != target[k] {
			k = fail[k-1]
		}
		if target[i] == target[k] {
			k++
		}
		fail[i] = k
	}
	for i, k := 0, 0; i < len(value); i++ {
		for k > 0 && value[i] != target[k] {
			k = fail[k-1]
		}
		if value[i] == target[k] {
			k++
		}
		if k == len(target) {
			return i - k + 1
		}
	}
	return -1
}

// reverseUTF16 returns a reversed copy of units.
func reverseUTF16(units []uint16) []uint16 {
	reversed := make([]uint16, len(units))
	for i, unit := range units {
		reversed[len(units)-1-i] = unit
	}
	return reversed
}

func builtinStringIndexOf(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	value := call.This.string()
	target := call.Argument(0).string()
	if 2 > len(call.ArgumentList) {
		return intValue(indexRune(value, target))
	}
	position := toIntegerFloat(call.Argument(1))
	// Positions count UTF-16 code units, not bytes.
	if isASCII(value) {
		start := clampStringPosition(position, len(value))
		if index := strings.Index(value[start:], target); index >= 0 {
			return intValue(start + index)
		}
		return intValue(-1)
	}
	source := utf16.Encode([]rune(value))
	start := clampStringPosition(position, len(source))
	if index := indexUTF16(source[start:], utf16.Encode([]rune(target))); index >= 0 {
		return intValue(start + index)
	}
	return intValue(-1)
}

func builtinStringLastIndexOf(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	value := call.This.string()
	target := call.Argument(0).string()
	if 2 > len(call.ArgumentList) || call.ArgumentList[1].IsUndefined() {
		return intValue(lastIndexRune(value, target))
	}
	// Convert the position once: float64 may call valueOf or toString.
	position := call.ArgumentList[1].float64()
	if math.IsNaN(position) {
		// A NaN position means +Infinity, so the whole string is searched.
		position = math.Inf(1)
	}
	position = math.Trunc(position)
	// Positions count UTF-16 code units, not bytes.
	if isASCII(value) {
		end := min(clampStringPosition(position, len(value))+len(target), len(value))
		return intValue(strings.LastIndex(value[:end], target))
	}
	source := utf16.Encode([]rune(value))
	search := utf16.Encode([]rune(target))
	end := min(clampStringPosition(position, len(source))+len(search), len(source))
	// The last match ending by end is the first match of the reversed target in the reversed prefix.
	index := indexUTF16(reverseUTF16(source[:end]), reverseUTF16(search))
	if index < 0 {
		return intValue(-1)
	}
	return intValue(end - index - len(search))
}

func builtinStringMatch(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := call.This.string()
	matcherValue := call.Argument(0)
	matcher := matcherValue.object()
	if !matcherValue.IsObject() || matcher.class != classRegExpName {
		matcher = call.runtime.newRegExp(matcherValue, Value{})
	}
	global := matcher.get("global").bool()
	if !global {
		match, result := execRegExp(matcher, target)
		if !match {
			return nullValue
		}
		return objectValue(execResultToArray(call.runtime, target, result))
	}

	result := matcher.regExpValue().regularExpression.FindAllStringIndex(target, -1)
	if result == nil {
		matcher.put("lastIndex", intValue(0), true)
		return Value{} // !match
	}
	matchCount := len(result)
	valueArray := make([]Value, matchCount)
	for index := range matchCount {
		valueArray[index] = stringValue(target[result[index][0]:result[index][1]])
	}
	matcher.put("lastIndex", intValue(result[matchCount-1][1]), true)
	return objectValue(call.runtime.newArrayOf(valueArray))
}

var builtinStringReplaceRegexp = regexp.MustCompile("\\$(?:[\\$\\&\\'\\`1-9]|0[1-9]|[1-9][0-9])")

func builtinStringFindAndReplaceString(input []byte, lastIndex int, match []int, target []byte, replaceValue []byte) []byte {
	matchCount := len(match) / 2
	output := input
	if match[0] != lastIndex {
		output = append(output, target[lastIndex:match[0]]...)
	}
	replacement := builtinStringReplaceRegexp.ReplaceAllFunc(replaceValue, func(part []byte) []byte {
		// TODO Check if match[0] or match[1] can be -1 in this scenario
		switch part[1] {
		case '$':
			return []byte{'$'}
		case '&':
			return target[match[0]:match[1]]
		case '`':
			return target[:match[0]]
		case '\'':
			return target[match[1]:]
		}
		matchNumberParse, err := strconv.ParseInt(string(part[1:]), 10, 64)
		if err != nil {
			return nil
		}
		matchNumber := int(matchNumberParse)
		if matchNumber >= matchCount {
			return nil
		}
		offset := 2 * matchNumber
		if match[offset] != -1 {
			return target[match[offset]:match[offset+1]]
		}
		return nil // The empty string
	})

	return append(output, replacement...)
}

func builtinStringReplace(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := []byte(call.This.string())
	searchValue := call.Argument(0)
	searchObject := searchValue.object()

	// TODO If a capture is -1?
	var search *regexp.Regexp
	global := false
	find := 1
	if searchValue.IsObject() && searchObject.class == classRegExpName {
		regExp := searchObject.regExpValue()
		search = regExp.regularExpression
		if regExp.global {
			find = -1
			global = true
		}
	} else {
		search = regexp.MustCompile(regexp.QuoteMeta(searchValue.string()))
	}

	found := search.FindAllSubmatchIndex(target, find)
	if found == nil {
		return stringValue(string(target)) // !match
	}

	lastIndex := 0
	result := []byte{}
	replaceValue := call.Argument(1)
	if replaceValue.isCallable() {
		target := string(target)
		replace := replaceValue.object()
		for _, match := range found {
			if match[0] != lastIndex {
				result = append(result, target[lastIndex:match[0]]...)
			}
			matchCount := len(match) / 2
			argumentList := make([]Value, matchCount+2)
			for index := range matchCount {
				offset := 2 * index
				if match[offset] != -1 {
					argumentList[index] = stringValue(target[match[offset]:match[offset+1]])
				} else {
					argumentList[index] = Value{}
				}
			}
			// Replace expects rune offsets not byte offsets.
			startIndex := utf8.RuneCountInString(target[0:match[0]])
			argumentList[matchCount+0] = intValue(startIndex)
			argumentList[matchCount+1] = stringValue(target)
			replacement := replace.call(Value{}, argumentList, false, nativeFrame).string()
			result = append(result, []byte(replacement)...)
			lastIndex = match[1]
		}
	} else {
		replace := []byte(replaceValue.string())
		for _, match := range found {
			result = builtinStringFindAndReplaceString(result, lastIndex, match, target, replace)
			lastIndex = match[1]
		}
	}

	if lastIndex != len(target) {
		result = append(result, target[lastIndex:]...)
	}

	if global && searchObject != nil {
		searchObject.put("lastIndex", intValue(lastIndex), true)
	}

	return stringValue(string(result))
}

func builtinStringSearch(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := call.This.string()
	searchValue := call.Argument(0)
	search := searchValue.object()
	if !searchValue.IsObject() || search.class != classRegExpName {
		search = call.runtime.newRegExp(searchValue, Value{})
	}
	result := search.regExpValue().regularExpression.FindStringIndex(target)
	if result == nil {
		return intValue(-1)
	}
	return intValue(result[0])
}

func builtinStringSplit(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := call.This.string()

	separatorValue := call.Argument(0)
	limitValue := call.Argument(1)
	limit := -1
	if limitValue.IsDefined() {
		limit = int(toUint32(limitValue))
	}

	if limit == 0 {
		return objectValue(call.runtime.newArray(0))
	}

	if separatorValue.IsUndefined() {
		return objectValue(call.runtime.newArrayOf([]Value{stringValue(target)}))
	}

	if separatorValue.isRegExp() {
		targetLength := len(target)
		search := separatorValue.object().regExpValue().regularExpression
		valueArray := []Value{}
		result := search.FindAllStringSubmatchIndex(target, -1)
		lastIndex := 0
		found := 0

		for _, match := range result {
			if match[0] == match[1] {
				// FIXME Ugh, this is a hack
				if match[0] == 0 || match[0] == targetLength {
					continue
				}
			}

			if lastIndex != match[0] {
				valueArray = append(valueArray, stringValue(target[lastIndex:match[0]]))
				found++
			} else if lastIndex == match[0] {
				if lastIndex != -1 {
					valueArray = append(valueArray, stringValue(""))
					found++
				}
			}

			lastIndex = match[1]
			if found == limit {
				goto RETURN
			}

			captureCount := len(match) / 2
			for index := 1; index < captureCount; index++ {
				offset := index * 2
				value := Value{}
				if match[offset] != -1 {
					value = stringValue(target[match[offset]:match[offset+1]])
				}
				valueArray = append(valueArray, value)
				found++
				if found == limit {
					goto RETURN
				}
			}
		}

		if found != limit {
			if lastIndex != targetLength {
				valueArray = append(valueArray, stringValue(target[lastIndex:targetLength]))
			} else {
				valueArray = append(valueArray, stringValue(""))
			}
		}

	RETURN:
		return objectValue(call.runtime.newArrayOf(valueArray))
	} else {
		separator := separatorValue.string()

		splitLimit := limit
		excess := false
		if limit > 0 {
			splitLimit = limit + 1
			excess = true
		}

		split := strings.SplitN(target, separator, splitLimit)

		if excess && len(split) > limit {
			split = split[:limit]
		}

		valueArray := make([]Value, len(split))
		for index, value := range split {
			valueArray[index] = stringValue(value)
		}

		return objectValue(call.runtime.newArrayOf(valueArray))
	}
}

// builtinStringSlice returns the string sliced by the given values
// which are rune not byte offsets, as per String.prototype.slice.
func builtinStringSlice(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := []rune(call.This.string())

	length := int64(len(target))
	start, end := rangeStartEnd(call.ArgumentList, length, false)
	if end-start <= 0 {
		return stringValue("")
	}
	return stringValue(string(target[start:end]))
}

func builtinStringSubstring(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := []rune(call.This.string())

	length := int64(len(target))
	start, end := rangeStartEnd(call.ArgumentList, length, true)
	if start > end {
		start, end = end, start
	}
	return stringValue(string(target[start:end]))
}

func builtinStringSubstr(call FunctionCall) Value {
	target := []rune(call.This.string())

	size := int64(len(target))
	start, length := rangeStartLength(call.ArgumentList, size)

	if start >= size {
		return stringValue("")
	}

	if length <= 0 {
		return stringValue("")
	}

	if length >= size-start {
		// Cap length to be to the end of the string
		// start = 3, length = 5, size = 4 [0, 1, 2, 3]
		// 4 - 3 = 1
		// target[3:4]
		length = size - start
	}

	return stringValue(string(target[start : start+length]))
}

func builtinStringStartsWith(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	target := call.This.string()
	search := call.Argument(0).string()
	length := len(search)
	if length > len(target) {
		return boolValue(false)
	}
	return boolValue(target[:length] == search)
}

func builtinStringToLowerCase(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	return stringValue(strings.ToLower(call.This.string()))
}

func builtinStringToUpperCase(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	return stringValue(strings.ToUpper(call.This.string()))
}

// 7.2 Table 2 — Whitespace Characters & 7.3 Table 3 - Line Terminator Characters.
const builtinStringTrimWhitespace = "\u0009\u000A\u000B\u000C\u000D\u0020\u00A0\u1680\u180E\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u2028\u2029\u202F\u205F\u3000\uFEFF"

func builtinStringTrim(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	return toValue(strings.Trim(call.This.string(),
		builtinStringTrimWhitespace))
}

func builtinStringTrimStart(call FunctionCall) Value {
	return builtinStringTrimLeft(call)
}

func builtinStringTrimEnd(call FunctionCall) Value {
	return builtinStringTrimRight(call)
}

// Mozilla extension, not ECMAScript 5.
func builtinStringTrimLeft(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	return toValue(strings.TrimLeft(call.This.string(),
		builtinStringTrimWhitespace))
}

// Mozilla extension, not ECMAScript 5.
func builtinStringTrimRight(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	return toValue(strings.TrimRight(call.This.string(),
		builtinStringTrimWhitespace))
}

func builtinStringLocaleCompare(call FunctionCall) Value {
	checkObjectCoercible(call.runtime, call.This)
	this := call.This.string() //nolint:ifshort
	that := call.Argument(0).string()
	if this < that {
		return intValue(-1)
	} else if this == that {
		return intValue(0)
	}
	return intValue(1)
}

func builtinStringToLocaleLowerCase(call FunctionCall) Value {
	return builtinStringToLowerCase(call)
}

func builtinStringToLocaleUpperCase(call FunctionCall) Value {
	return builtinStringToUpperCase(call)
}
