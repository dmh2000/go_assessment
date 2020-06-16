package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

// channel for file watcher -> websocket handler
var refresh chan string

// watch for a change in any files under 'app'
// and notify the websocket handler
func watch() {

	var w *fsnotify.Watcher

	// create a watcher
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err.Error())
	}
	// add the ./app directory to be watched
	err = w.Add("./app")
	if err != nil {
		panic(err.Error())
	}

	// wait for file change notifications
	var event fsnotify.Event
	for {
		select {
		case event = <-w.Events:
			// notification : signal the websocket handler
			fmt.Println(event)
			refresh <- "refresh"
		case err = <-w.Errors:
			fmt.Println(err.Error())
			w.Close()
			return
		}
	}
}

// run the tests and return the results as html
func handler(w http.ResponseWriter, r *http.Request) {
	var t []string
	var s string

	t = runAllTests()
	s = TestToHTML(t)
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
