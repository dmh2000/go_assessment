package goassessment

import "fmt"

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
