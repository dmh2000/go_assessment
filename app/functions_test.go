package goassessment

import (
	"fmt"
	"math/rand"
	"testing"
)

// return a function from a function
func TestFunction(t *testing.T) {
	var s string
	var r string
	t.Log("you should be able to return a function from a function")

	r = "Hello, world"

	// ===============================
	// your function
	// ===============================
	f := fFunction("Hello")

	// check return is valid
	if f == nil {
		t.Error("fFunction returned nil")
		return
	}

	// check type of function
	q := fmt.Sprintf("%T", f)
	if q != "func(string) string" {
		t.Error(shouldBe(q, "func(string) string"))
	}

	s = f("world")
	if s != r {
		t.Error(shouldBe(s, r))
	}

	r = "Hai, can i haz funxtion?"
	f = fFunction("Hai")
	if f == nil {
		t.Error("fFunction returnd nil")
	}

	s = f("can i haz funxtion?")
	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// return an array of closures over the values in arr
func TestClosures(t *testing.T) {
	t.Log("you should be able to use closures")

	arr := []int{rand.Int(), rand.Int(), rand.Int()}
	square := func(x int) int { return x * x }

	// ===============================
	// your function
	// ===============================
	funcs := fMakeClosures(square, arr)
	// ===============================

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

// partial application
func TestPartial(t *testing.T) {
	t.Log("you should be able to create a 'partial' function")

	f := func(a string, b string, c string) string {
		return a + ", " + b + c
	}

	// ===============================
	// your function
	// ===============================
	g := fPartial(f, "Hello", "Ellie")
	// ===============================

	// valid return
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

	r := "Hello, Ellie!!!"
	if s != r {
		t.Error(shouldBe(s, r))
	}

}

// create a new array populated with tyhe result of calling the provided function
func TestMap(t *testing.T) {
	t.Log("you should be able to implement a function MAP function")

	square := func(x int) int { return x * x }

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 4, 9, 16, 25}

	// ===============================
	// your function
	// ===============================
	c := fMap(square, a)
	// ===============================

	if len(c) != len(a) {
		t.Error(shouldBe(len(c), len(a)))
	}

	for i := 0; i < len(b); i++ {
		if c[i] != b[i] {
			t.Error(shouldBe(c[i], b[i]))
		}
	}
}

// return a value by applying the reducer fn to each element of the array
func TestReduce(t *testing.T) {
	t.Log("you should be able to implement a function REDUCE function")

	sum := func(acc int, val int) int {
		acc += val
		return acc
	}

	a := []int{1, 2, 3, 4, 5}

	// ===============================
	// your function
	// ===============================
	b := fReduce(sum, a)
	// ===============================

	if b != 15 {
		t.Error(shouldBe(b, 15))
	}
}

// return a new array contains only the elements for which the function is true
func TestFilter(t *testing.T) {

	odd := func(v int) bool { return (v % 2) == 1 }

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 3, 5}

	// ===============================
	// your function
	// ===============================
	c := fFilter(odd, a)
	// ===============================

	if len(c) != len(b) {
		t.Error(shouldBe(len(c), len(b)))
	}

	for i := 0; i < len(b); i++ {
		if c[i] != b[i] {
			t.Error(shouldBe(c[i], b[i]))
		}
	}
}
