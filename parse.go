package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func parseTestLine(line string) string {
	var matched bool
	var err error
	var r *regexp.Regexp
	var s []string

	// Test header
	r, err = regexp.Compile(`@(.+)`)
	s = r.FindStringSubmatch(line)
	// matched, err = regexp.MatchString("@(.+)", line)
	if err != nil {
		panic(err.Error)
	}
	if len(s) >= 2 {
		return fmt.Sprintf("<hr/><h3 class='hdr'>%s</h3><hr/>", s[1])
	}

	// === RUN
	matched, err = regexp.MatchString(`=== RUN`, line)
	if err != nil {
		panic(err.Error)
	}
	if matched {
		return fmt.Sprintf("<div class='run'>%s</div>", line)
	}
	matched, err = regexp.MatchString(`--- PASS`, line)
	if err != nil {
		panic(err.Error)
	}
	if matched {
		return fmt.Sprintf("<div class='pass'>%s</div>", line)
	}

	matched, err = regexp.MatchString(`--- FAIL`, line)
	if err != nil {
		panic(err.Error)
	}
	if matched {
		return fmt.Sprintf("<div class='fail'>%s</div>", line)
	}

	matched, err = regexp.MatchString(`    .+_test`, line)
	if err != nil {
		panic(err.Error)
	}
	if matched {
		return fmt.Sprintf("<div class='item'>%s</div>", line)
	}

	matched, err = regexp.MatchString(`^FAIL`, line)
	if err != nil {
		panic(err.Error)
	}
	if matched {
		return ""
	}

	return fmt.Sprintf("<div class='other'>%s</div>", line)
}

func header() string {
	s := `
	<html>
	<style>
	body {
		font-family:monospace;
	}
	.hdr {color:magenta;margin-top:2px;margin-bottom:2px;font-weight:bold;}
	.run {color:blue;margin-top:2px;font-weight:bold;}
	.fail {color:red;margin-top:2px;font-weight:bold;}
	.pass {color:green;margin-top:2px;font-weight:bold;}
	.item {color:blue;margin-top:1px;font-size:small;}
	.other {color:black;margin-top:1px;}
	</style>
	<body>
	`
	return s
}

func trailer() string {
	return `
	</body>
	<script>
	const socket = new WebSocket('ws://' + location.host + '/refresh');

	socket.addEventListener('open', function (event) {
		console.log('websocket open')
	});

	// Listen for messages
	socket.addEventListener('message', function (event) {
		console.log('Message from server ', event.data);
		document.body.innerHTML = '<div>reloading</div>'
		setTimeout( ()=> {
			location.reload()
		},100)
	});

	</script>
	</html>
	`
}

// TestToHTML convert test output to html
func TestToHTML(fname string) string {
	var s string

	// html header
	s = fmt.Sprintf("%s", header())

	// body from test results
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s += parseTestLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// html trailer
	s += fmt.Sprintf("%s", trailer())

	return s
}
