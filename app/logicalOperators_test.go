package goassessment

import "testing"

func TestLogicalOr(t *testing.T) {
	t.Log("you should be able to work with logical or")
	var b bool

	b = or(false, true)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = or(true, false)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = or(true, true)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = or(false, false)
	if b {
		t.Error(shouldBe(b, false))
	}
}

func TestLogicalAnd(t *testing.T) {
	t.Log("you should be able to work with logical and")
	var b bool

	b = and(false, true)
	if b {
		t.Error(shouldBe(b, false))
	}

	b = and(true, false)
	if b {
		t.Error(shouldBe(b, false))
	}

	b = and(true, true)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = and(false, false)
	if b {
		t.Error(shouldBe(b, false))
	}
}

func TestLogicalNot(t *testing.T) {
	t.Log("you should be able to work with logical not")
	var b bool

	b = not(false)
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = not(true)
	if b {
		t.Error(shouldBe(b, false))
	}
}
