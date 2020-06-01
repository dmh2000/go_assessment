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
