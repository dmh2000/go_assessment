package goassessment

import "container/list"

// fix the Sort interface for PersonSlice
// fix the Heap interface for PersonSlice sorted by ascending age
// fix the List interface for PersonSlice

// Len ...
func (p PersonSlice) Len() int {
	return len(p)
}

// Less ...
func (p PersonSlice) Less(i int, j int) bool {
	return p[i].age < p[j].age
}

// Swap ...
func (p PersonSlice) Swap(i int, j int) {
	t := p[i]
	p[i] = p[j]
	p[j] = t
}

// Push ...
func (p *PersonSlice) Push(x interface{}) {
	*p = append(*p, x.(Person))
}

// Pop ...
func (p *PersonSlice) Pop() interface{} {
	n := len(*p)       // copy of slice header
	x := (*p)[n-1]     // save last element in list
	*p = (*p)[0 : n-1] // truncate list
	return x           // return last element
}

// write a function that creates and populates a list
func populateList(p PersonSlice) *list.List {
	var a *list.List
	a = list.New()
	for _, v := range p {
		a.PushBack(v)
	}
	return a
}

// write a function that sends a notification for different types
func sendNotification(n Notifier) string {
	return n.notify()
}

// write a function that returns an integer value from an interface{}
func intFromInterface(i interface{}) int {
	return i.(int)
}
