package goassessment

import (
	"testing"
)

// Dir recursive definition of a file structure
type Dir struct {
	name  string
	files []string
	dirs  []Dir
}

var fileData = Dir{
	name:  "app",
	files: []string{"index.html"},
	dirs: []Dir{
		{
			name:  "js",
			files: []string{"main.js", "app.js", "misc.js"},
			dirs: []Dir{
				{
					name:  "vendor",
					files: []string{"jquery.js", "underscore.js"},
				},
				{
					name:  "css",
					files: []string{"reset.css", "main.css"},
					dirs:  []Dir{},
				},
			},
		},
	},
}

// write a function that returns a list of files starting from the named directory
func TestListFiles(t *testing.T) {
	t.Log("you should be able to return a list of files from the data")
	var index int
	// passing in 'nil' means list all files

	var result = listFiles(fileData, "")
	if len(result) != 8 {
		t.Error(shouldBe(len(result), 8))
	}

	index = testIndexOfStrings(result, "index.html")
	if index < 0 {
		t.Error(index, "should be >= 0")
	}

	index = testIndexOfStrings(result, "main.js")
	if index < 0 {
		t.Error(index, "should be >= 0")
	}

	index = testIndexOfStrings(result, "notfound.js")
	if index < 0 {
		t.Error(shouldBe(index, -1))
	}
}

// wrote a function that returns all permutations of the input array
func TestListDir(t *testing.T) {
	t.Log("you should be able to return a list of files in a subdir")
	var index int
	// passing in 'nil' means list all files

	var result = listFiles(fileData, "js")
	if len(result) != 5 {
		t.Error(shouldBe(len(result), 5))
	}

	index = testIndexOfStrings(result, "main.js")
	if index < 0 {
		t.Error(index, "should be >= 0")
	}

	index = testIndexOfStrings(result, "underscore.js")
	if index < 0 {
		t.Error(index, "should be >= 0")
	}
}

var permData = []int{1, 2, 3, 4}

var permutations = [][]int{
	{1, 2, 3, 4},
	{1, 2, 4, 3},
	{1, 3, 2, 4},
	{1, 3, 4, 2},
	{1, 4, 2, 3},
	{1, 4, 3, 2},
	{2, 1, 3, 4},
	{2, 1, 4, 3},
	{2, 3, 1, 4},
	{2, 3, 4, 1},
	{2, 4, 1, 3},
	{2, 4, 3, 1},
	{3, 1, 2, 4},
	{3, 1, 4, 2},
	{3, 2, 1, 4},
	{3, 2, 4, 1},
	{3, 4, 1, 2},
	{3, 4, 2, 1},
	{4, 1, 2, 3},
	{4, 1, 3, 2},
	{4, 2, 1, 3},
	{4, 2, 3, 1},
	{4, 3, 1, 2},
	{4, 3, 2, 1},
}

// write a function that returns the fibonnaci number of n
// Note : order of return values is arbitrary
func TestPermute(t *testing.T) {

	t.Log("you should be able to return the permutations of an array")
	var result [][]int

	result = permute(permData)

	// length of answers should be the same
	if len(result) != len(permutations) {
		t.Error(shouldBe(len(result), len(permutations)))
	}

	// check for all permutations, order not required
	for _, target := range result {
		index := testIndexOfIntSlice(permutations, target)
		if index < 0 {
			t.Error("index of ", target, " should be >=0 ")
		}
	}
}

// write a function that returns the fibonnaci number of n
func TestFibonacci(t *testing.T) {
	t.Log("you should be able to return the nth number in a fibonacci sequence")
	var fib int

	fib = fibonacci(2)
	if fib != 1 {
		t.Error(shouldBe(fib, 1))
	}
	fib = fibonacci(6)
	if fib != 8 {
		t.Error(shouldBe(fib, 8))
	}
}

// write a function that returns an array of strings with all valid sets of n parens
func TestValidParens(t *testing.T) {
	t.Log("you should be able to return the set of all valid combinations of n pairs of parentheses.")

	var s []string
	var r []string

	r = []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	s = validParentheses(3)
	if len(s) != len(r) {
		t.Error(shouldBe(len(s), len(r)))
	}

	for i := 0; i < len(r); i++ {
		index := testIndexOfStrings(s, r[i])
		if index < 0 {
			t.Error(s, "should contain", r[i])
		}
	}
}
