package goassessment

// write a function that returns a function
func fFunction(str string) func(string) string {
	return nil
}

// write a function that returns a slice of closures
func fMakeClosures(fn func(x int) int, arr []int) []func() int {
	return nil
}

// write a function that implements a partial application
func fPartial(
	fn func(a string, b string, c string) string,
	str1 string,
	str2 string,
) func(string) string {
	return nil
}

// write a function that creates a new array populated
// with the result of calling the provided function
func fMap(fn func(a int) int, arr []int) []int {
	return nil
}

// write a function that applies the reducer to a slice of integers
func fReduce(fn func(acc int, val int) int, arr []int) int {
	return -1
}

// write a function that applies the filter to a slice of integers
func fFilter(fn func(a int) bool, arr []int) []int {
	return nil
}

// write a function that handles a variadic argument
func fVariadic(s ...int) string {
	return ""
}
