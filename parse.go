package main

import (
	"fmt"
	"regexp"
	"strings"
)

// results of a single test
type testResults struct {
	pass int
	fail int
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

	// Test Category
	r, err = regexp.Compile(`@(.+)`)
	s = r.FindStringSubmatch(line)
	if err != nil {
		panic(err)
	}
	if len(s) >= 2 {
		h := catHeader(s[1])
		ti.html = append(ti.html, h)
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

	// individual test passsed
	matched, err = regexp.MatchString(`--- PASS`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.pass++
		ti.html = append(ti.html, testPass(line))
		return
	}

	// individual test failed
	matched, err = regexp.MatchString(`--- FAIL`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.fail++
		ti.html = append(ti.html, testFail(line))
		return
	}

	// test goal
	matched, err = regexp.MatchString(`.*GOAL*`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, testGoal(line))
		return
	}

	matched, err = regexp.MatchString(`    .+_test`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, testInfo(line))
		return
	}

	matched, err = regexp.MatchString(`    test_util.go`, line)
	if err != nil {
		panic(err)
	}
	if matched {
		ti.html = append(ti.html, testInfo(line))
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
		ti.fail++
		ti.html = append(ti.html, fmt.Sprintf("<div class='fail'>compile failed : %s</div>", line))
	}

	return
}

// statistics
func statHeader(pass int, fail int, timestamp string) string {
	var header []string 

	header = make([]string,5)

	// create the pass/fail stats
	var percent float64
	if fail == 0 {
		percent = 100.0
	} else {
		percent = (float64(pass) / float64(pass+fail)) + 0.005
	}

	header[0] = fmt.Sprintf(`<div class="row row-box"><div class="col-5 col-pad"><h5>Go-Assessment Test Results</h5></div>`)
	header[1] = fmt.Sprintf("<div class=\"col-2 badge badge-danger badge-pad\">progress %.0f%%</div>",percent)
	header[2] = fmt.Sprintf(`<div class="col-2 badge badge-success badge-pad">passed %2d</div>`,pass)
	header[3] = fmt.Sprintf(`<div class="col-2 badge badge-danger badge-pad">failed %2d</div>`,fail)
	header[4] = fmt.Sprintf(`</div><div class="row"><div class="col-12 col-pad">%s</div></div>`,timestamp)

	return strings.Join(header,"")
}

// category header
func catHeader(category string) string {
   return  fmt.Sprintf(`<hr/><div class="row"><div class="col-10 col-pad font-weight-bold">%s</div></div>`,category)
}

// individual test passed
func testPass(name string) string {
	return fmt.Sprintf(`<div class="row">
	  <div class="col-10 col-pad text-success">%s</div>
	  <div class="col-2 col-pad"><span class="badge badge-success res-pad">PASS</span></div>
    </div>`,name[9:])
}

func testFail(name string) string {
	return fmt.Sprintf(`<div class="row">
	  <div class="col-10 col-pad text-danger">%s</div>
  	  <div class="col-2 col-pad"><span class="badge badge-danger res-pad">FAIL</span></div>
    </div>`,name[9:])
}

// individual test goal
func testGoal(line string) string {

	index := strings.Index(line,"GOAL")
	location := line[0:index]
	goal := line[index+5:]

	return fmt.Sprintf(`<hr/><div class="row">
	<div class="col-3 col-pad text-info">%s</div>
	<div class="col-9 col-pad text-info">%s</div>
	</div>`,location, goal)
}

// individual test info (line where it failed)
func testInfo(line string) string {
	index := strings.LastIndex(line,":")
	location := line[0:index]
	info := line[index+1:]

	return fmt.Sprintf(`<div class="row">
	<div class="col-3 col-pad text-danger item">%s</div>
	<div class="col-9 col-pad text-danger item">%s</div>
	</div>`,location, info)

	// return fmt.Sprintf(`<div class="row"><div class="col-10 col-pad text-danger item">%s</div></div>`,info)
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

	// return the body as a string
	return statHeader(ti.pass, ti.fail, timestamp) + strings.Join(ti.html, "")
}
