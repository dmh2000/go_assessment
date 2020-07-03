package goassessment

import (
	"fmt"
	"runtime"
	"testing"
)

// test output helper
func shouldBe(a interface{}, b interface{}) string {
	return fmt.Sprintf("%v should be %v\n", a, b)
}

// check if two slices are equal
func testIntSliceEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true

}

// return index of target or -1 if not found
func testIndexOfStrings(a []string, target string) int {
	for i, v := range a {
		if v == target {
			return i
		}
	}
	return -1
}

// find index of slice of integers in 2d array
func testIndexOfIntSlice(a [][]int, target []int) int {
	for i, v := range a {
		// different lengths
		if len(v) != len(target) {
			return -1
		}
		// equal values
		for j := 0; j < len(v); j++ {
			if target[j] == v[j] {
				return i
			}
		}
	}
	return -1
}

// catch a panic in a test
// must be called with defer
func testPanic(t *testing.T) {
	r := recover()
	if r != nil {
		_, _, line, ok := runtime.Caller(3)
		if ok {
			t.Error("panic in solution at line ", line)
		} else {
			t.Error("panic in solution (no file/line available) ")
		}
	}
}
