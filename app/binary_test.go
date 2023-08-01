package goassessment

import (
	"testing"
)

// write a function that returns the 0 or 1 value at the specified bit position
func TestValueAtBit(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to find the value of a given bit")
	var bit int

	// first bit is at bit position 0
	bit = valueAtbit(1, 0)
	if bit != 1 {
		t.Error(shouldBe(bit, 1))
	}

	bit = valueAtbit(128, 7)
	if bit != 1 {
		t.Error(shouldBe(bit, 1))
	}

	bit = valueAtbit(65, 0)
	if bit != 1 {
		t.Error(shouldBe(bit, 0))
	}

	bit = valueAtbit(65, 6)
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
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to return the base10 representation of a binary string")
	val := base10("11000000")
	if val != 192 {
		t.Error(shouldBe(val, 192))
	}
}

// write a function that converts the int value to a binary string
func TestConvertoBinary(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to convert an eight-bit number to a binary string")
	var bin string

	bin = convertToBinary(128)
	if bin != "10000000" {
		t.Error(shouldBe(bin, "10000000"))
	}

	bin = convertToBinary(65)
	if bin != "1000001" {
		t.Error(shouldBe(bin, "1000001"))
	}
}

// write a function that returns the bitwise OR of the input values
func TestBitwiseOr(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to use the bitwise OR operator")

	r := 770047
	x := bitwiseOr(0x12345, 0x33333, 0xabcde)
	if x != r {
		t.Error(shouldBe(x, r))
	}
}

// write a function that returns the bitwise AND of the input values
func TestBitwiseAnd(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to use the bitwise AND operator")

	r := 8192
	x := bitwiseAnd(0x12345, 0x33333, 0xabcde)
	if x != r {
		t.Error(shouldBe(x, r))
	}

}

// write a function that returns the bitwise XOR of the input values
func TestBitwiseXor(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to use the bitwise XOR operator")

	r := 568488
	x := bitwiseXor(0x12345, 0x33333, 0xabcde)
	if x != r {
		t.Error(shouldBe(x, r))
	}

}
