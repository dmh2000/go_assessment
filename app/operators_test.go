package goassessment

import "testing"

// ARITHMETIC OPERATORS ON INTEGERS

// write a function that returns the sum of two values
func TestIntegerSum(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with addition operator")
	a := 6
	b := 2
	r := add(a, b)
	if r != 8 {
		t.Error(shouldBe(r, 8))
	}
}

// write a function that returns the subtract between two values
func TestIntegerDifference(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	var r int

	t.Log("GOAL: you should be able to work with the subtraction operator")
	a := 6
	b := 2
	r = subtract(a, b)
	if r != 4 {
		t.Error(shouldBe(r, 4))
	}
}

// write a function that returns a value equal to the multiply of two values
func TestIntegerProduct(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with the multiplication operator")
	a := 6
	b := 2
	r := multiply(a, b)
	if r != 12 {
		t.Error(shouldBe(r, 12))
	}
}

// write a function that returns a value equal to the divide of two values
func TestIntegerQuotient(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with the division operator")
	a := 6
	b := 2
	r := divide(a, b)
	if r != 3 {
		t.Error(shouldBe(r, 3))
	}

	a = 7
	b = 2
	r = divide(a, b)
	if r != 3 {
		t.Error(shouldBe(r, 3))
	}
}

// write a function that returns a value equal to the remainder of two values
func TestIntegerRemainder(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with the remainder (modulo) operator")
	var a int
	var b int
	var r int

	a = 7
	b = 2
	r = remainder(a, b)
	if r != 1 {
		t.Error(shouldBe(r, 1))
	}

	a = 6
	b = 2
	r = remainder(a, b)
	if r != 0 {
		t.Error(shouldBe(r, 0))
	}
}

// COMPARISON OPERATORS
// write a function that returns true or false if two values are equal
func TestEqual(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging
	var a int
	var b int
	var r bool

	t.Log("GOAL: you should be able to work with the equal operator")
	a = 7
	b = 2
	r = equal(a, b)
	if r {
		t.Error(shouldBe(r, false))
	}
	a = 2
	b = 2
	r = equal(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
}

// write a function that returns true or false if two values are not equal
func TestNotEqual(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging
	var a int
	var b int
	var r bool

	t.Log("GOAL: you should be able to work with the not equal operator")
	a = 7
	b = 2
	r = notEqual(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
	a = 2
	b = 2
	r = notEqual(a, b)
	if r {
		t.Error(shouldBe(r, false))
	}
}

// write a function that returns true or false if one value is less than the other one
func TestLessThan(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging
	var a int
	var b int
	var r bool

	t.Log("GOAL: you should be able to work with the less than operator")
	a = 7
	b = 2
	r = lessThan(a, b)
	if r {
		t.Error(shouldBe(r, false))
	}
	a = 2
	b = 7
	r = lessThan(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
}

// write a function that returns true or false if one value is less than or equal to the other one
func TestLessThanEqual(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging
	var a int
	var b int
	var r bool

	t.Log("GOAL: you should be able to work with the less than or equal than operator")
	a = 7
	b = 2
	r = lessThanEqual(a, b)
	if r {
		t.Error(shouldBe(r, false))
	}
	a = 2
	b = 7
	r = lessThanEqual(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}

	a = 2
	b = 2
	r = lessThanEqual(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
}

// write a function that returns true or false if one value is greater than the other one
func TestGreaterThan(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging
	var a int
	var b int
	var r bool

	t.Log("GOAL: you should be able to work with the greater than operator")
	a = 7
	b = 2
	r = greaterThan(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
	a = 2
	b = 7
	r = greaterThan(a, b)
	if r {
		t.Error(shouldBe(r, false))
	}
}

// write a function that returns true or false if one value is greater than or equal to the other one
func TestGreaterThanEqual(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging
	var a int
	var b int
	var r bool

	t.Log("GOAL: you should be able to work with the greater than or equal operator")
	a = 7
	b = 2
	r = greaterThanEqual(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
	a = 2
	b = 7
	r = greaterThanEqual(a, b)
	if r {
		t.Error(shouldBe(r, false))
	}

	a = 2
	b = 2
	r = greaterThanEqual(a, b)
	if !r {
		t.Error(shouldBe(r, true))
	}
}

// LOGICAL OPERATORS

// write a function that returns true if either argument is true
func TestLogicalOr(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with logical or")
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
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with logical and")
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
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to work with logical not")
	var b bool

	b = none(false, true)
	if b {
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
	if !b {
		t.Error(shouldBe(b, true))
	}
}
