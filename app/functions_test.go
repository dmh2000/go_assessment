package goassessment

import (
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
	s = fFunction("Hello")("world")

	if s != r {
		t.Error(shouldBe(s, r))
	}

	r = "Hai, can i haz funxtion?"
	s = fFunction("Hai")("can i haz funxtion?")
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
	s := fPartial(f, "Hello", "Ellie")("!!!")

	r := "Hello, Ellie!!!"
	if s != r {
		t.Error(shouldBe(s, r))
	}

}

// create a new array populated with tyhe result of calling the provided function
func TestMap(t *testing.T) {

}

// return a value by applying the reducer fn to each element of the array
func TestReduce(t *testing.T) {

}

// return a new array contains only the elements for which the function is true
func TestFilter(t *testing.T) {

}
