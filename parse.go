package main

import (
	"fmt"
	"regexp"
	"strings"
)

// results of a single test
type testResults struct {
	pass bool
	html []string
}

func parseTestLine(line string, ti *testResults) {
	var matched bool
	var err error
	var r *regexp.Regexp
	var s []string

	// ===================================================================
	// find matches to the line format and output the corresponding HTML
	// ===================================================================

	// @start
	r, err = regexp.Compile(`@start (.+)`)
	s = r.FindStringSubmatch(line)
	if err != nil {
		panic(err)
	}
	if len(s) >= 2 {
		h := []string{
			`<a id="`,
			s[1],
			`"`,
			`</a>`,
			`<div>`,
		}
		ti.html = append(ti.html, strings.Join(h, ""))
		return
	}

	// @end
	matched, err = regexp.MatchString(`@end`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, "</div>")
		return
	}

	// Test header
	r, err = regexp.Compile(`@(.+)`)
	s = r.FindStringSubmatch(line)
	if err != nil {
		panic(err)
	}
	if len(s) >= 2 {
		ti.html = append(ti.html, fmt.Sprintf("<hr/><div class='h5 hdr'>%s</div>", s[1]))
		return
	}

	matched, err = regexp.MatchString(`=== RUN`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		// don't print the === RUN output
		return
	}

	matched, err = regexp.MatchString(`--- PASS`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.pass = true
		ti.html = append(ti.html, fmt.Sprintf("<div class='pass'>%s</div>", line))
		return
	}

	matched, err = regexp.MatchString(`--- FAIL`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.pass = false
		ti.html = append(ti.html, fmt.Sprintf("<div class='fail'>%s</div>", line))
		return
	}

	matched, err = regexp.MatchString(`    .+_test`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, fmt.Sprintf("<div class='item'>%s</div>", line))
		return
	}

	matched, err = regexp.MatchString(`^FAIL|^ok|^PASS`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		return
	}

	return
}

// TestToHTML : convert array of test results to html ouput
func TestToHTML(results [][]string) string {
	var ti testResults

	// for each line in the test results
	for i := 0; i < len(results); i++ {
		// transform line to HTML
		for _, v := range results[i] {
			parseTestLine(v, &ti)
		}
	}

	return strings.Join(ti.html, "")
}
