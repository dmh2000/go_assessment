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
			`<div>`,
			`<a id="`,
			s[1],
			`"></a>`,
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
		h := []string{
			`</div>`,
		}
		ti.html = append(ti.html, strings.Join(h, ""))
		return
	}

	// Test header
	r, err = regexp.Compile(`@(.+)`)
	s = r.FindStringSubmatch(line)
	if err != nil {
		panic(err)
	}
	if len(s) >= 2 {
		h := []string{
			`<hr/>`,
			`<div style="width:80%;overflow:hidden">`,
			`<span style="float:left;" class="h5 hdr">`,
			s[1],
			`</span>`,
			`<span style="float:right;"><a href="#top">&nbsp;▲&nbsp;</a></span>`,
			`</div>`,
		}
		ti.html = append(ti.html, strings.Join(h, ""))
		// ti.html = append(ti.html, fmt.Sprintf("<div style='width:'80%;'><span style='float:left;' class='h5 hdr'>%s</span><span style='float:right;'><a href='#top'>top</span></div>", s[1]))
		// ti.html = append(ti.html, fmt.Sprintf("<hr/><div class='h5 hdr'>%s</div>", s[1]))
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

	matched, err = regexp.MatchString(`.*GOAL*`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, fmt.Sprintf("<div class='spec'>%s</div>", line))
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

	matched, err = regexp.MatchString(`    test_util.go`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, fmt.Sprintf("<div class='item'>%s</div>", line))
		return
	}

	matched, err = regexp.MatchString(`^ok|^PASS`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		return
	}

	/* example output if tests did not compile properly
	# command-line-arguments [command-line-arguments.test]
	app/strings_test.go:4:2: imported and not used: "log"
	FAIL    command-line-arguments [build failed]
	FAIL
	*/
	matched, err = regexp.MatchString(`^app/.*\.go`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.pass = false
		ti.html = append(ti.html, fmt.Sprintf("<div class='fail'>compile failed : %s</div>", line))
	}

	return
}

// TestToHTML : convert array of test results to html content
func TestToHTML(timestamp string, results [][]string) string {
	var ti testResults

	// for each line in the test results
	for i := 0; i < len(results); i++ {
		// transform line to HTML
		for _, v := range results[i] {
			parseTestLine(v, &ti)
		}
	}

	// create the timestamp header
	timeheader := "<div>" + timestamp + "</div>"

	// return the body as a string
	return timeheader + strings.Join(ti.html, "")
}
