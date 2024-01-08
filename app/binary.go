package goassessment

import (
	"fmt"
)

// write a function that returns the 0 or 1 value at the specified bit position
func valueAtbit(num int, bit int) int {
	v := num >> bit
	return v & 0x0001
}

// write a function that returns the base 10 value of the binary string
func base10(n string) int {
	// bit shift version
	r := 0
	shift := 0
	for _, v := range n {
		r <<= shift
		if v == '1' {
			r |= 1
		}
		shift = 1
	}
	// fmt version
	// fmt.Sscanf(n, "%b", &r)
	return r
}

// write a function that returns the binary value of num as a string
func convertToBinary(num int) string {
	return fmt.Sprintf("%b", num)
}

// write a function that returns the bitwise OR of the input values
func bitwiseOr(x int, y int, z int) int {
	return x | y | z
}

// write a function that returns the bitwise AND of the input values
func bitwiseAnd(x int, y int, z int) int {
	return x & y & z
}

// write a function that returns the bitwise XOR of the input values
func bitwiseXor(x int, y int, z int) int {
	return x ^ (y ^ z)
}
