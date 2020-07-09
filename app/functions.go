package goassessment

import "strconv"

// write a function that returns a function
func fFunction(str string) func(string) string {
	return func(s string) string {
		return str + ", " + s
	}
}

// write a function that returns a slice of closures
func fMakeClosures(fn func(x int) int, arr []int) []func() int {
	var f []func() int

	f = make([]func() int, len(arr))
	for i := 0; i < len(arr); i++ {
		j := arr[i]
		f[i] = func() int {
			return fn(j)
		}
	}

	return f
}

// write a function that implements a partial application
func fPartial(
	fn func(a string, b string, c string) string,
	str1 string,
	str2 string,
) func(string) string {
	return func(s string) string {
		return fn(str1, str2, s)
	}
}

// write a function that creates a new array populated
// with the result of calling the provided function
func fMap(fn func(a int) int, arr []int) []int {
	var r []int
	r = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		r[i] = fn(arr[i])
	}
	return r
}

// write a function that applies the reducer to a slice of integers
func fReduce(fn func(acc int, val int) int, arr []int) int {
	var r int
	r = 0
	for i := 0; i < len(arr); i++ {
		r = fn(r, arr[i])
	}
	return r
}

// write a function that applies the filter to a slice of integers
func fFilter(fn func(a int) bool, arr []int) []int {
	var r []int

	r = make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) {
			r = append(r, arr[i])
		}
	}
	return r
}

// write a function that handles a variadic argument
func fVariadic(s ...int) string {

	var r string

	r = strconv.Itoa(s[0])
	for i := 1; i < len(s); i++ {
		r += "," + strconv.Itoa(s[i])
	}
	return r
}
