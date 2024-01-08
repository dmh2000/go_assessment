package goassessment

import (
	"testing"
)

// write a function that returns and empty map
func TestCollectionMap(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to populate a map with keys and values")

	keys := []interface{}{"a", "b", "c"}
	values := []interface{}{1, 2, 3}
	m := MakeMap(keys, values)
	if len(m) == 0 {
		t.Error("map is empty")
	}

	v, ok := m["a"]
	if !ok {
		t.Error("map does not contain key 'a'")
	}
	if v != 1 {
		t.Error("map['a'] is not 1")
	}

	v, ok = m["b"]
	if !ok {
		t.Error("map does not contain key 'b'")
	}
	if v != 2 {
		t.Error("map['b'] is not 2")
	}

	v, ok = m["c"]
	if !ok {
		t.Error("map does not contain key 'c'")
	}
	if v != 3 {
		t.Error("map['c'] is not 3")
	}
}

// write PushFront and PopFront functions for a linked list
func TestCollectionFront(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to push and pop values from the front of a linked list")

	list := &LinkedList{}

	PushFront(list, 1)
	PushFront(list, 2)
	PushFront(list, 3)

	if list.head == nil {
		t.Error("list.head is nil")
	}

	node := list.head
	if node.value != 3 {
		t.Error("node.value is not 3")
	}

	node = node.next
	if node.value != 2 {
		t.Error("node.value is not 2")
	}

	node = node.next
	if node.value != 1 {
		t.Error("node.value is not 1")
	}

	v := PopFront(list)
	if v != 3 {
		t.Error("PopFront did not return 3")
	}

	v = PopFront(list)
	if v != 2 {
		t.Error("PopFront did not return 2")
	}

	v = PopFront(list)
	if v != 1 {
		t.Error("PopFront did not return 1")
	}

	if list.head != nil {
		t.Error("list.head is not nil")
	}
}

// write PushBack and PopBack functions for a linked list
func TestCollectionBack(t *testing.T) {
	// write PushFront and PushBack functions for a linked list
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to push and pop values from the back of a linked list")

	list := &LinkedList{}

	PushBack(list, 1)
	PushBack(list, 2)
	PushBack(list, 3)

	if list.head == nil {
		t.Error("list.head is nil")
	}

	node := list.head
	if node.value != 1 {
		t.Error("node.value is not 1")
	}

	node = node.next
	if node.value != 2 {
		t.Error("node.value is not 2")
	}

	node = node.next
	if node.value != 3 {
		t.Error("node.value is not 3")
	}

	v := PopBack(list)
	if v != 3 {
		t.Error("PopBack did not return 3")
	}

	v = PopBack(list)
	if v != 2 {
		t.Error("PopBack did not return 2")
	}

	v = PopBack(list)
	if v != 1 {
		t.Error("PopBack did not return 1")
	}

	if list.head != nil {
		t.Error("list.head is not nil")
	}
}

func TestCollectionFrontBack(t *testing.T) {
	// write PushFront and PushBack functions for a linked list
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to push and pop values from the front and back of a linked list")

	list := &LinkedList{}

	PushBack(list, 1)
	PushFront(list, 2)
	PushBack(list, 3)
	PushFront(list, 4)

	// should be 4, 2, 1, 3

	if list.head == nil {
		t.Error("list.head is nil")
	}

	node := list.head
	if node.value != 4 {
		t.Error("node.value is not 1")
	}

	node = node.next
	if node.value != 2 {
		t.Error("node.value is not 2")
	}

	node = node.next
	if node.value != 1 {
		t.Error("node.value is not 1")
	}

	node = node.next
	if node.value != 3 {
		t.Error("node.value is not 3")
	}

	v := PopFront(list)
	if v != 4 {
		t.Error("PopBack did not return 4")
	}

	v = PopBack(list)
	if v != 3 {
		t.Error("PopBack did not return 3")
	}

	v = PopFront(list)
	if v != 2 {
		t.Error("PopBack did not return 2")
	}

	v = PopBack(list)
	if v != 1 {
		t.Error("PopBack did not return 1")
	}

	if list.head != nil {
		t.Error("list.head is not nil")
	}
}
