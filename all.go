package main

import (
	"os/exec"
)

func runTest(title string, test string) []string {
	var b []byte
	var s []string
	var app string = "./app/" + test + ".go"
	var tst string = "./app/" + test + "_test.go"
	var utl string = "./app/test_util.go"
	cmd := exec.Command("go", "test", "-v", app, tst, utl)
	b, _ = cmd.CombinedOutput()

	// store the title
	s = make([]string, 0)
	s = append(s, "@start "+test)
	s = append(s, title)

	// parse out lines in the output
	t := ""
	for _, c := range b {
		if c == 10 {
			s = append(s, t)
			t = ""
		} else {
			t += string(c)
		}
	}
	s = append(s, t)

	s = append(s, "@end")

	return s
}

// test command arguments
type testPair struct {
	title string
	prog  string
}

// list of all tests to execute
var tests []testPair = []testPair{
	{"@Flow Control", "flowControl"},
	{"@Slices", "slices"},
	{"@Numbers", "numbers"},
	{"@Logical Operators", "logicalOperators"},
	{"@Strings", "strings"},
	{"@Functions", "functions"},
	{"@Recursion", "recursion"},
	{"@Methods", "methods"},
	{"@Regex", "regex"},
	{"@Async", "async"},
}

// run tests, capture all output and return it as a string
func runAllTests() [][]string {
	var s [][]string
	var t []string

	s = make([][]string, 0)
	for _, p := range tests {
		t = runTest(p.title, p.prog)
		s = append(s, t)
	}

	return s
}
