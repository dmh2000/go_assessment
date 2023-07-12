package goassessment

import (
	"strconv"
)

// write a function that returns a function
func fFunction(str string) func(string) string {

	f := func(s string) string {
		return str + ", " + s
	}

	return f
}

// write a function that returns a slice of closures
func fMakeClosures(fn func(x int) int, arr []int) []func() int {
	a := make([](func() int), 0)

	for i := range arr {
		u := arr[i]
		f := func() int {
			r := fn(u)
			return r
		}
		a = append(a, f)
	}
	return a
}

// write a function that implements a partial application
func fPartial(
	fn func(a string, b string, c string) string,
	str1 string,
	str2 string,
) func(string) string {
	f := func(s string) string {
		return fn(str1, str2, s)
	}
	return f
}

// write a function that creates a new array populated
// with the result of calling the provided function
func fMap(fn func(a int) int, arr []int) []int {
	r := []int{}
	for _, v := range arr {
		r = append(r, fn(v))
	}
	return r
}

// write a function that applies the reducer to a slice of integers
func fReduce(fn func(acc int, val int) int, arr []int) int {
	a := 0
	for _, v := range arr {
		a = fn(a, v)
	}
	return a
}

// write a function that applies the filter to a slice of integers
func fFilter(fn func(a int) bool, arr []int) []int {
	r := []int{}

	for _, v := range arr {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}

// write a function that handles a variadic argument
func fVariadic(s ...int) string {
	r := ""
	comma := ""
	for _, v := range s {
		r += comma + strconv.Itoa(v)
		comma = ","
	}
	return r
}
