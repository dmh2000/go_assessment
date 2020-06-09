package goassessment

import "testing"

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

func TestBase10(t *testing.T) {
	t.Log("you should be able to return the base10 representation of a binary string'")

	var val int
	val = base10("11000000")
	if val != 192 {
		t.Error(shouldBe(val, 192))
	}
}

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

func TestMultiplyWithPrecision(t *testing.T) {
	t.Log("you should be able to multiply with precision")

	var f float64

	f = multiplyWithPrecision(3.0, 0.1)
	if f != 0.3 {
		t.Error(shouldBe(f, 0.3))
	}

	f = multiplyWithPrecision(3.0, 0.2)
	if f != 0.6 {
		t.Error(shouldBe(f, 0.6))
	}

	f = multiplyWithPrecision(3.0, 0.0001)
	if f != 0.3 {
		t.Error(shouldBe(f, 0.0003))
	}

}
