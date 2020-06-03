package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var s string
	cmd := exec.Command("./test.sh")
	fmt.Println("executing tests")
	err := cmd.Run()
	if err != nil {
		s = err.Error()
	} else {
		s = TestToHTML("test.txt")
	}
	fmt.Fprintf(w, s, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
