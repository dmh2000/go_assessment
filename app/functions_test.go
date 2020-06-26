package goassessment

import (
	"fmt"
	"math/rand"
	"testing"
)

// write a function that returns a function
func TestFunction(t *testing.T) {
	t.Log("you should be able to return a function from a function")
	var s string
	var r string

	f := fFunction("Hello")
	if f == nil {
		t.Error("fFunction returned nil")
		return
	}

	// check type of function
	q := fmt.Sprintf("%T", f)
	if q != "func(string) string" {
		t.Error(shouldBe(q, "func(string) string"))
	}

	//
	r = "Hello, world"
	s = f("world")
	if s != r {
		t.Error(shouldBe(s, r))
	}

	r = "Hai, can i haz funxtion?"
	f = fFunction("Hai")
	if f == nil {
		t.Error("fFunction returned nil")
	}

	r = "Hai, can i haz funxtion?"
	s = f("can i haz funxtion?")
	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// write a function that returns a slice of closures
func TestClosures(t *testing.T) {
	t.Log("you should be able to use closures")

	arr := []int{rand.Int(), rand.Int(), rand.Int()}
	square := func(x int) int { return x * x }

	funcs := fMakeClosures(square, arr)
	if len(funcs) != len(arr) {
		t.Error(shouldBe(len(funcs), len(arr)))
	}

	for i, f := range funcs {
		u := f()
		v := square(arr[i])
		if u != v {
			t.Error(shouldBe(u, v))
		}
	}
}

// write a function that implements a partial application
func TestPartial(t *testing.T) {
	t.Log("you should be able to create a 'partial' function")

	f := func(a string, b string, c string) string {
		return a + ", " + b + c
	}

	g := fPartial(f, "Hello", "Ellie")
	// check valid return
	if g == nil {
		t.Error("fPartial returned nil")
		return
	}

	// check type of function
	q := fmt.Sprintf("%T", g)
	if q != "func(string) string" {
		t.Error(shouldBe(q, "func(string) string"))
	}

	// invoke it
	s := g("!!!")

	// check function output
	r := "Hello, Ellie!!!"
	if s != r {
		t.Error(shouldBe(s, r))
	}

}

// write a function that creates a new array populated
// with the result of calling the provided function
func TestMap(t *testing.T) {
	t.Log("you should be able to implement a function MAP function")

	square := func(x int) int { return x * x }

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 4, 9, 16, 25}

	c := fMap(square, a)
	if len(c) != len(a) {
		t.Error(shouldBe(len(c), len(a)))
	}

	for i := 0; i < len(b); i++ {
		if c[i] != b[i] {
			t.Error(shouldBe(c[i], b[i]))
		}
	}
}

// write a function that applies the reducer to a slice of integers
func TestReduce(t *testing.T) {
	t.Log("you should be able to implement a function REDUCE function")

	sum := func(acc int, val int) int {
		acc += val
		return acc
	}

	a := []int{1, 2, 3, 4, 5}

	b := fReduce(sum, a)
	if b != 15 {
		t.Error(shouldBe(b, 15))
	}
}

// write a function that applies the filter to a slice of integers
func TestFilter(t *testing.T) {

	// condition
	odd := func(v int) bool { return (v % 2) == 1 }

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 3, 5}

	c := fFilter(odd, a)
	if len(c) != len(b) {
		t.Error(shouldBe(len(c), len(b)))
	}

	for i := 0; i < len(b); i++ {
		if c[i] != b[i] {
			t.Error(shouldBe(c[i], b[i]))
		}
	}
}

// write a function that handles a variadic argument
func TestVariadic(t *testing.T) {

	a := []int{1, 2, 3, 4, 5}

	s := fVariadic(a...)
	r := "1,2,3,4,5"
	if s != r {
		t.Error(shouldBe(s, r))
	}
}
