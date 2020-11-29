package goassessment

import "container/list"

// fix the Sort interface for PersonSlice
// fix the Heap interface for PersonSlice sorted by ascending age
// fix the List interface for PersonSlice

// Len ...
func (p PersonSlice) Len() int { return len(p) }

// Less ...
func (p PersonSlice) Less(i int, j int) bool {
	return false
}

// Swap ...
func (p PersonSlice) Swap(i int, j int) {

}

// Push ...
func (p *PersonSlice) Push(x interface{}) {

}

// Pop ...
func (p *PersonSlice) Pop() interface{} {
	return nil
}

// write a function that creates and populates a list
func populateList(p PersonSlice) *list.List {
	var a *list.List
	return a
}

// write a function that sends a notification for different types
func sendNotification(n Notifier) string {
	return ""
}

// write a function that returns an integer value from an interface{}
func intFromInterface(i interface{}) int {
	return -1
}
