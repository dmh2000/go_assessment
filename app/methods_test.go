package goassessment

import (
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

func TestMethodsSort(t *testing.T) {
	t.Log("you should be able to implement the Sort interface for a specified type")

	p := PersonSlice{
		{20, "January"}, {18, "April"}, {45, "June"}, {30, "August"},
	}

	// fix the sort implementation in methods.go
	sort.Sort(p)

	q := PersonSlice{
		{18, "April"}, {20, "January"}, {30, "August"}, {45, "June"},
	}
	for i := 0; i < len(p); i++ {
		if p[i] != q[i] {
			t.Error(shouldBe(p[i], q[i]))
		}
	}
}

type Notifier interface {
	notify() string
}

func (p Person) notify() string {
	return fmt.Sprintf("new mail for %s", p.name)
}

func (e Employee) notify() string {
	return fmt.Sprintf("new mail for %v:%s", e.id, e.p.name)
}

func TestMethodsInterface(t *testing.T) {
	t.Log("you should be able to use an interface")
	var s string
	var r string

	p := Person{name: "October", age: 99}
	e := Employee{p: Person{name: "September", age: 31}, id: 12345}

	// =========================
	// your function
	// =========================
	s = sendNotification(e)
	// =========================

	r = "new mail for 12345:September"
	if s != r {
		t.Error(shouldBe(s, r))
	}

	// =========================
	// your function
	// =========================
	s = sendNotification(p)
	// =========================

	r = "new mail for October"
	if s != r {
		t.Error(shouldBe(s, r))
	}
}
