package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var refresh chan string

func watch() {

	var w *fsnotify.Watcher
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err.Error())
	}
	err = w.Add("./app")
	if err != nil {
		panic(err.Error())
	}

	var event fsnotify.Event
	for {
		select {
		case event = <-w.Events:
			fmt.Println(event)
			refresh <- "refresh"
		case err = <-w.Errors:
			fmt.Println(err.Error())
			w.Close()
			return
		}
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	var s string
	cmd := exec.Command("./all.sh")
	fmt.Println("executing tests")
	err := cmd.Run()
	if err != nil {
		s = err.Error()
	} else {
		s = TestToHTML("test.txt")
	}
	fmt.Fprint(w, s) // , r.URL.Path[1:])
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// wait for a file change
		select {
		case <-refresh:
			// refresh the page
			err = conn.WriteJSON(true)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {

	refresh = make(chan string)

	// set up the http handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/refresh", wshandler)
	fmt.Println("listening on :8080")

	// start the file watcher
	go watch()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
