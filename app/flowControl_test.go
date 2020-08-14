package goassessment

import (
	"testing"
)

// write a function that receives a number as its argument;
// if the number is divisible by 3, the function should return 'fizz';
// if the number is divisible by 5, the function should return 'buzz';
// if the number is divisible by 3 and 5, the function should return
// 'fizzbuzz';
//
// otherwise the function should return the number as a string
func TestFlowControl(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to conditionally branch your code")

	var num int
	var s string

	// prime number
	num = 97

	s = fizzBuzz(2)
	if s != "2" {
		t.Error(shouldBe(s, "2"))
	}

	s = fizzBuzz(101)
	if s != "101" {
		t.Error(shouldBe(s, "101"))
	}

	s = fizzBuzz(3)
	if s != "fizz" {
		t.Error(shouldBe(s, "fizz"))
	}

	s = fizzBuzz(6)
	if s != "fizz" {
		t.Error(shouldBe(s, "fizz"))
	}

	s = fizzBuzz(num * 3)
	if s != "fizz" {
		t.Error(shouldBe(s, "fizz"))
	}

	s = fizzBuzz(5)
	if s != "buzz" {
		t.Error(shouldBe(s, "buzz"))
	}

	s = fizzBuzz(10)
	if s != "buzz" {
		t.Error(shouldBe(s, "buzz"))
	}

	s = fizzBuzz(num * 5)
	if s != "buzz" {
		t.Error(shouldBe(s, "buzz"))
	}

	s = fizzBuzz(15)
	if s != "fizzbuzz" {
		t.Error(shouldBe(s, "fizzbuzz"))
	}

	s = fizzBuzz(num * 3 * 5)
	if s != "fizzbuzz" {
		t.Error(shouldBe(s, "fizzbuzz"))
	}
}
