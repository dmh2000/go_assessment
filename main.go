package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type testPage struct {
	Body template.HTML
}

// run the tests and return the results as html
func handler(w http.ResponseWriter, r *http.Request) {
	var tests [][]string
	var page testPage

	// execute the tests
	tests = runAllTests()

	// compile the test results into a string
	page.Body = template.HTML(TestToHTML(tests))

	// create the template
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "server error handler:65")
		return
	}

	// execute the template
	err = t.Execute(w, &page)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func main() {
	// set up the http handlers
	http.HandleFunc("/", handler)
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
