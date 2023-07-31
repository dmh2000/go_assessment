package goassessment

import (
	"testing"
)

// write a function that returns and empty map
func TestCollectionMap(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

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
		t.Error("map['c'] is not 2")
	}
}

// write a function that returns a linked list supporting any type
func TestCollectionLinkedList(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	data := []interface{}{1, 2, 3, 4}
	list := MakeList(data)
	if list.head == nil {
		t.Error("list.head is nil")
	}

	node := list.head

	for _, v := range data {
		if node == nil {
			t.Error("node is nil")
			break
		}

		if node.value.(int) != v {
			t.Error("node.value is not", v)
			break
		}
		node = node.next
	}

	if node != nil {
		t.Error("node is not nil")
	}
}
