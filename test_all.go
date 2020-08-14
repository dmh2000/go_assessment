package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func runTest(title string, test string, f *os.File) []string {
	var b []byte
	var s []string
	var app string = "./app/" + test + ".go"
	var tst string = "./app/" + test + "_test.go"
	var utl string = "./app/test_util.go"
	var typ string = "./app/app_types.go"
	cmd := exec.Command("go", "test", "-v", app, tst, utl, typ)
	b, _ = cmd.CombinedOutput()

	// save to results.txt
	// ignore errors
	_, err := f.Write(b)
	if err != nil {
		log.Printf("can't write results of %v to results.txt", title)
	}

	// store the title
	s = make([]string, 0)
	s = append(s, "@start "+test)
	s = append(s, title+" : "+time.Now().Format("04:05"))

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
	{"@Logical Operators", "logicalOperators"},
	{"@Strings", "strings"},
	{"@Slices", "slices"},
	{"@Binary", "binary"},
	{"@Functions", "functions"},
	{"@Recursion", "recursion"},
	{"@Methods", "methods"},
	{"@Regexp", "regex"},
	{"@Async", "async"},
}

// run tests, capture all output and return it as a string
func runAllTests() [][]string {
	var s [][]string
	var t []string

	f, err := os.OpenFile("./results.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Can't Open results.txt for writing")
	}

	s = make([][]string, 0)
	for _, p := range tests {
		t = runTest(p.title, p.prog, f)
		s = append(s, t)
	}

	f.Close()

	return s
}
