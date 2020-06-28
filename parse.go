package main

import (
	"fmt"
	"regexp"
	"strings"
)

type testResults struct {
	pass bool
	html []string
}

func parseTestLine(line string, ti *testResults) {
	var matched bool
	var err error
	var r *regexp.Regexp
	var s []string

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

	// === RUN
	matched, err = regexp.MatchString(`=== RUN`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, fmt.Sprintf("<div class='run'>%s</div>", line))
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
		ti.html = append(ti.html, "")
		return
	}

	ti.html = append(ti.html, fmt.Sprintf("<div class='other'>%s</div>", line))
	return
}

// TestToHTML : convert array of test results to html ouput
func TestToHTML(results [][]string) string {
	var ti testResults
	// html header
	// s = fmt.Sprintf("%s", header())

	for i := 0; i < len(results); i++ {
		// get test info
		// ...

		// get the test strings
		for _, v := range results[i] {
			parseTestLine(v, &ti)
		}
	}

	// html trailer
	// s += fmt.Sprintf("%s", trailer())

	return strings.Join(ti.html, "")
}
