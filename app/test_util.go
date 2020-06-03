package goassessment

import "fmt"

func shouldBe(a interface{}, b interface{}) string {
	return fmt.Sprintf("%v should be %v\n", a, b)
}

func intSliceEqual(a []int, b []int) bool {
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
func indexOfStrings(a []string, target string) int {
	for i, v := range a {
		if v == target {
			return i
		}
	}
	return -1
}

func indexOfIntSlice(a [][]int, target []int) int {
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
