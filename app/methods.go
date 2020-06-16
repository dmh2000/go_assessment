package goassessment

import "container/list"

// fix the Sort, Heap and List interfaces for PersonSlice
// Len ...
func (p PersonSlice) Len() int { return 0 }

// Less ...
func (p PersonSlice) Less(i int, j int) bool { return true }

// Swap ...
func (p PersonSlice) Swap(i int, j int) {}

// Push ...
func (p *PersonSlice) Push(x interface{}) {}

// Pop ...
func (p *PersonSlice) Pop() interface{} { return nil }

// create and populate a list
func populateList(p PersonSlice) *list.List {
	return nil
}

// notification
func sendNotification(n Notifier) string {
	return ""
}
