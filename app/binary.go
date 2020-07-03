package goassessment

import "strconv"

// write a function that returns the 0 or 1 value at the specified bit position
func valueAtbit(num int, bit int) int {
	return (num >> bit) & 0x01
}

// write a function that returns the base 10 value of the binary string
func base10(n string) int {
	i, _ := strconv.ParseInt(n, 2, 31)
	return int(i)
}

// write a function that returns the binary value of num as a string
func convertToBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
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
	return x ^ y&z
}
