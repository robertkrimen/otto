package otto

import "testing"

type GoSliceTest []int

func (s GoSliceTest) Sum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func TestGoSlice(t *testing.T) {
	tt(t, func() {
		test, vm := test()
		vm.Set("TestSlice", GoSliceTest{1, 2, 3})
		is(test(`TestSlice.length`).export(), 3)
		is(test(`TestSlice[1]`).export(), 2)
		is(test(`TestSlice.Sum()`).export(), 6)
	})
}

func TestGoSliceSetLength(t *testing.T) {
	tt(t, func() {
		test, vm := test()
		vm.Set("s", GoSliceTest{1, 2, 3})
		test(`s.length = 2; s.length`, 2)
		test(`s.length = 4; [s.length, s[1], s[3]].join(",")`, "4,2,0")
		test(`raise: s.length = -1`, "RangeError: Invalid array length")
		test(`raise: s.length = 1.5`, "RangeError: Invalid array length")
		test(`raise: s.length = 4294967296`, "RangeError: Invalid array length")
		test(`raise: s.length = "abc"`, "RangeError: Invalid array length")
		test(`s.length`, 4)
	})
}
