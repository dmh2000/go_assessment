package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"
)

// anything but / and /index.html
func handle404(w http.ResponseWriter, r *http.Request) {
	log.Println("404:", r.RequestURI)
	w.WriteHeader(404)
	fmt.Fprint(w, "404 not found")
}

// something broke
func handle500(w http.ResponseWriter, r *http.Request, err string) {
	log.Println(err)
	w.WriteHeader(500)
	fmt.Fprintf(w, "500 server error")
}

// page template
type testPage struct {
	Body template.HTML
}

// execute tests and render the results as html
func handleTests(w http.ResponseWriter, r *http.Request) {
	defer func() {
		rec := recover()
		if rec != nil {
			handle500(w, r, "error in handleTests")
		}
	}()

	// allow only / or /index.html
	var m bool
	m, _ = regexp.MatchString(`^/$|^/index.html$`, r.RequestURI)
	if !m {
		handle404(w, r)
		return
	}

	log.Println("req:", r.RequestURI)

	var tests [][]string
	var page testPage

	log.Println("Tests Started")

	// execute the tests
	t0 := time.Now()
	tests = runAllTests(t0)
	t1 := time.Now()
	// tests finished

	// elapsed time
	et := t1.Sub(t0)

	// print time for tests to execute
	log.Println("Tests Finished:", et)

	// compile the test results into a string
	timestamp := fmt.Sprintf("%s : ET %.3f sec", t0.Local().Format("Mon Jan 2 15:04:05 -0700 MST 2006"), et.Seconds())

	// timestamp := t0.Local().Format("Mon Jan 2 15:04:05 -0700 MST 2006") + ": ET " + (t1.Sub(t0).Seconds())
	page.Body = template.HTML(TestToHTML(timestamp, tests))

	// create the template
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		handle500(w, r, err.Error())
		return
	}

	// execute the template
	err = t.Execute(w, &page)
	if err != nil {
		log.Println(err)
		handle500(w, r, err.Error())
	}
}

func main() {
	// get specified port
	port := flag.String("port", "8080", "http://ipaddr:port")
	flag.Parse()

	// set up the http handlers
	http.HandleFunc("/index.html", handleTests)
	http.HandleFunc("/", handleTests)
	log.Println("listening on :", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
