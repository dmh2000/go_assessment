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
	var app = "./app/" + test + ".go"
	var tst = "./app/" + test + "_test.go"
	var utl = "./app/test_util.go"
	var typ = "./app/app_types.go"
	cmd := exec.Command("go", "test", "-v", app, tst, utl, typ)
	b, _ = cmd.CombinedOutput()

	// save to results.txt
	// ignore errors
	if f != nil {
		_, err := f.Write(b)
		if err != nil {
			log.Printf("can't write results of %v to results.txt", title)
		}
	}

	// store the title
	s = make([]string, 0)
	s = append(s, "@start "+test)
	s = append(s, title+" : "+time.Now().Format("15:04:05"))

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
	{"@Operators", "operators"},
	{"@Flow Control", "flowControl"},
	{"@Strings", "strings"},
	{"@Slices", "slices"},
	{"@Binary", "binary"},
	{"@Functions", "functions"},
	{"@Methods", "methods"},
	{"@Collections", "collections"},
	{"@Recursion", "recursion"},
	{"@Regexp", "regex"},
	{"@Async", "async"},
}

// run tests, capture all output and return it as a string
func runAllTests(now time.Time) [][]string {
	var s [][]string
	var t []string

	f, err := os.OpenFile("./results.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Can't Open results.txt for writing")
		f = nil
	}

	if f != nil {
		// output date/time stamp
		f.WriteString(now.Local().Format("Mon Jan 2 15:04:05 -0700 MST 2006") + "\n")
	}

	s = make([][]string, 0)
	for _, p := range tests {
		t = runTest(p.title, p.prog, f)
		s = append(s, t)
	}

	f.Close()

	return s
}
