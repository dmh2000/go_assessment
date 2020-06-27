package goassessment

import (
	"testing"
)

// write a function that returns the 0 or 1 value at the specified bit position
func TestValueAtBit(t *testing.T) {
	t.Log("you should be able to find the value of a given bit")
	var bit int

	bit = valueAtbit(128, 8)
	if bit != 1 {
		t.Error(shouldBe(bit, 1))
	}

	bit = valueAtbit(65, 1)
	if bit != 1 {
		t.Error(shouldBe(bit, 1))
	}

	bit = valueAtbit(65, 7)
	if bit != 1 {
		t.Error(shouldBe(bit, 1))
	}

	bit = valueAtbit(128, 1)
	if bit != 0 {
		t.Error(shouldBe(bit, 0))
	}

}

// write a function that returns the base 10 integer value of the binary string
func TestBase10(t *testing.T) {
	t.Log("you should be able to return the base10 representation of a binary string")
	var val int

	val = base10("11000000")
	if val != 192 {
		t.Error(shouldBe(val, 192))
	}
}

// write a function taht converts the int value to a binary string
func TestConvertoBinary(t *testing.T) {
	t.Log("you should be able to convert an eight-bit number to a binary string")
	var bin string

	bin = convertToBinary(128)
	if bin != "10000000" {
		t.Error(shouldBe(bin, "10000000"))
	}

	bin = convertToBinary(65)
	if bin != "01000001" {
		t.Error(shouldBe(bin, "01000001"))
	}
}

// write a function that returns the bitwise OR of the input values
func TestBitwiseOr(t *testing.T) {
	var x int
	var r int

	r = 770047
	x = bitwiseOr(0x12345, 0x33333, 0xabcde)
	if x != r {
		t.Error(shouldBe(x, r))
	}
}

// write a function that returns the bitwise AND of the input values
func TestBitwiseAnd(t *testing.T) {
	var x int
	var r int

	r = 8192
	x = bitwiseOr(0x12345, 0x33333, 0xabcde)
	if x != r {
		t.Error(shouldBe(x, r))
	}

}

// write a function that returns the bitwise XOR of the input values
func TestBitwiseXor(t *testing.T) {
	var x int
	var r int

	r = 568488
	x = bitwiseOr(0x12345, 0x33333, 0xabcde)
	if x != r {
		t.Error(shouldBe(x, r))
	}

}
