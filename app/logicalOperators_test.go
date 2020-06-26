package goassessment

import "testing"

// write a function that returns true if either argument is true
func TestLogicalOr(t *testing.T) {
	t.Log("you should be able to work with logical or")
	var b bool

	b = either(false, true)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = either(true, false)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = either(true, true)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = either(false, false)
	if b {
		t.Error(shouldBe(b, false))
	}
}

// write a function that returns true only if both arguments are true
func TestLogicalAnd(t *testing.T) {
	t.Log("you should be able to work with logical and")
	var b bool

	b = both(false, true)
	if b {
		t.Error(shouldBe(b, false))
	}

	b = both(true, false)
	if b {
		t.Error(shouldBe(b, false))
	}

	b = both(true, true)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = both(false, false)
	if b {
		t.Error(shouldBe(b, false))
	}
}

// write a function that returns true only if both arguments are false
func TestLogicalNot(t *testing.T) {
	t.Log("you should be able to work with logical not")
	var b bool

	b = none(false, true)
	if !b {
		t.Error(shouldBe(b, false))
	}

	b = none(true, false)
	if b {
		t.Error(shouldBe(b, false))
	}

	b = none(true, true)
	if b {
		t.Error(shouldBe(b, false))
	}

	b = none(false, false)
	if b {
		t.Error(shouldBe(b, true))
	}
}
