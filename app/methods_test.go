package goassessment

import (
	"container/heap"
	"container/list"
	"fmt"
	"sort"
	"testing"
)

type Person struct {
	age  int
	name string
}

type PersonSlice []Person

type Employee struct {
	p  Person
	id int
}

type EmployeeSlice []Employee

// Sort : implement Sort for PersonSlice
func TestMethodsSort(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to implement the Sort interface for a specified type")

	// unsorted
	p := PersonSlice{
		{20, "January"}, {18, "April"}, {45, "June"}, {30, "August"},
	}
	// sorted
	q := PersonSlice{
		{18, "April"}, {20, "January"}, {30, "August"}, {45, "June"},
	}

	// =========================================
	// fix the sort implementation in methods.go
	// sort by age
	// =========================================
	sort.Sort(p)

	for i := 0; i < len(p); i++ {
		if p[i] != q[i] {
			t.Error(shouldBe(p[i], q[i]))
		}
	}
}

// fix the Heap interface for PersonSlice sorted by ascending age
func TestMethodsHeap(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to implement the Heap interface for a specified type")

	p := PersonSlice{
		{20, "January"}, {18, "April"}, {45, "June"}, {30, "August"},
	}
	q := PersonSlice{
		{18, "April"}, {20, "January"}, {30, "August"}, {45, "June"},
	}

	// ===================================================
	// fix the container/heap implementation in methods.go
	// sort by age ascending
	// ====================================================
	// init the heap
	var h *PersonSlice
	h = &p
	heap.Init(h)

	// load the heap
	for i := 0; i < len(p); i++ {
		heap.Push(h, p[i])
	}

	// pop in ascending order by age
	for i := 0; i < len(q); i++ {
		r := heap.Pop(h)
		if r == nil {
			t.Error("heap should not be empty")
			return
		}
		if r.(Person).age != q[i].age {
			t.Error(shouldBe(r, q[i]))
		}
	}

	// should be empty
	if h.Len() > 0 {
		t.Error("heap should be empty")
	}
}

// write a function that creates and populates a list
func TestMethodsList(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to create and populate a list")
	var p *list.List
	var e *list.Element

	q := PersonSlice{
		{18, "April"}, {20, "January"}, {30, "August"}, {45, "June"},
	}

	// populate a list in order of increasing age
	p = populateList(q)
	if p == nil {
		t.Error("should return a list")
		return
	}

	e = p.Front()
	i := 0
	for e != nil {
		if i >= len(q) {
			t.Error("list is too long")
		}
		if e.Value.(Person).name != q[i].name {
			t.Error(shouldBe(e.Value.(Person), q[i]))
		}
		i++
		e = e.Next()
	}
}

// Notifier interface
type Notifier interface {
	notify() string
}

func (p Person) notify() string {
	return fmt.Sprintf("new mail for %s", p.name)
}

func (e Employee) notify() string {
	return fmt.Sprintf("new mail for %v:%s", e.id, e.p.name)
}

func TestMethodsNotify(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to use an interface with different types")
	var s string
	var r string

	p := Person{name: "October", age: 99}
	e := Employee{p: Person{name: "September", age: 31}, id: 12345}

	// =========================
	// your function
	s = sendNotification(e)
	r = "new mail for 12345:September"
	if s != r {
		t.Error(shouldBe(s, r))
	}

	// =========================
	// your function
	s = sendNotification(p)
	r = "new mail for October"
	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// write a function that returns an integer value from an interface{}
func TestIntFromInterface(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you shoud be able to extract a value from an interface")
	var x interface{}
	var i int
	var r int

	x = 5
	r = 5
	i = intFromInterface(x)
	if i != r {
		t.Error(shouldBe(i, r))
	}
}
