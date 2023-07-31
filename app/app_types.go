package goassessment

// Person definition
type Person struct {
	age  int
	name string
}

// PersonSlice slice of
type PersonSlice []Person

// Employee subtype of Person
type Employee struct {
	p  Person
	id int
}

// Notifier interface
type Notifier interface {
	notify() string
}

// Dir recursive definition of a file structure
type Dir struct {
	name  string
	files []string
	dirs  []Dir
}

type LinkedListElement struct {
	next  *LinkedListElement
	value interface{}
}

type LinkedList struct {
	head *LinkedListElement
}
