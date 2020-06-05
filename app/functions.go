package goassessment

// return a function from a function
func fFunction(str string) func(string) string {
	return nil
}

// return an array of closures over the values in arr
func fMakeClosures(fn func(x int) int, arr []int) []func() int {
	return nil
}

// partial application
func fPartial(
	fn func(a string, b string, c string) string,
	str1 string,
	str2 string,
) func(string) string {
	return nil
}

// create a new array populated with tyhe result of calling the provided function
func fMap(fn func(a int), arr []int) []int {
	return nil
}

// return a value by applying the reducer fn to each element of the array
func fReduce(fn func(acc int, val int, index int, a []int), arr []int) int {
	return 0
}

// return a new array contains only the elements for which the function is true
func fFilter(fn func(a int) bool, arr []int) []int {
	return nil
}

// return a function that concatenates the arguments
func fVariadic(s ...string) string {
	return ""
}
