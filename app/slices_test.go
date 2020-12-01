package goassessment

import (
	"sort"
	"testing"
)

// write a function that retursns the index of an item in a slice
func TestIndexOf(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to determine the location of an item in a slice")

	a := []int{1, 2, 3, 4}

	index := indexOf(a, 3)

	if index != 2 {
		t.Error(shouldBe(index, 2))
	}
	index = indexOf(a, 5)
	if index != -1 {
		t.Error(shouldBe(index, -1))
	}
}

// write a function that sums the values in a slice
func TestSum(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to sum the items of a slice")

	a := []int{1, 2, 3, 4}

	v := sum(a)

	if v != 10 {
		t.Error(shouldBe(v, 10))
	}
}

// write a function that removes all instances of a value from a slice
func TestRemove(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to remove all instances of a value from a slice")

	a := []int{1, 2, 3, 4, 2, 2}
	b := []int{1, 3, 4}

	a = remove(a, 2)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that returns the value of the first element in a slice (wihtout removing it)
func TestFront(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to get the value of the first element of a slice")

	a := []int{1, 2, 3, 4}

	v := front(a)

	if v != 1 {
		t.Error(shouldBe(v, 1))
	}
}

// write a function that returns the value of the last element in a slice (wihtout removing it)
func TestBack(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to get the value of the last element of a slice")

	a := []int{1, 2, 3, 4}

	v := back(a)

	if v != 4 {
		t.Error(shouldBe(v, 1))
	}
}

// write a function that adds an item to the end of a slice
func TestPushBack(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to add an item to the end of a slice")

	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 2}

	a = pushBack(a, 2)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that removes an item to the end of a slice
func TestPopBack(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to remove the last item of a slice")

	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3}

	a = popBack(a)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that adds an item to the front of a slice
func TestPushFront(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to add an item to the front of a slice")

	a := []int{1, 2, 3, 4}
	b := []int{2, 1, 2, 3, 4}

	a = pushFront(a, 2)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}

}

// write a function that removes an item from the front of a slice
func TestPopFront(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to remove the first item of a slice")

	a := []int{1, 2, 3, 4}
	b := []int{2, 3, 4}

	a = popFront(a)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that concatenates two slices
func TestConcat(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to join together two slices")
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	c := []int{1, 2, 3, 4, 5, 6, 7, 8}

	a = concat(a, b)

	if len(a) != len(c) {
		t.Error(shouldBe(len(a), len(c)))
	}
	if !testIntSliceEqual(a, c) {
		t.Error(shouldBe(a, c))
	}
}

// write a function that adds an item to a slice at the specified index
func TestInsert(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to add an item anywhere in a slice")

	a := []int{1, 2, 3, 4}

	a = insert(a, 5, 2)

	b := []int{1, 2, 5, 3, 4}
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}

	// insert at front
	a = []int{1}
	b = []int{5, 1}
	a = insert(a, 5, 0)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}

	// insert at front
	a = []int{1}
	b = []int{1, 5}
	a = insert(a, 5, 1)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}

	// if index > length of slice, append to end
	a = []int{}
	b = []int{5}
	a = insert(a, 5, 2)
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that returns a count of matching items in a slice
func TestCount(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to count the occurrences of an item in a slice")

	a := []int{1, 4, 2, 3, 4, 4}

	v := count(a, 4)

	if v != 3 {
		t.Error(shouldBe(v, 3))
	}
}

// write a function that finds duplicates in a slice
func TestDuplicates(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to find duplicates in a slice")

	a := []int{1, 2, 4, 4, 3, 3, 1, 5, 3}

	a = duplicates(a)

	sort.IntSlice(a).Sort()
	b := []int{1, 3, 4}
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that sqaures all items in a slice
func TestSquare(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to square each number in a slice")

	a := []int{1, 2, 3, 4}
	b := []int{1, 4, 9, 16}

	a = square(a)
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

// write a function that returns all the indices in a slice that matches an item
func TestFindAllOccurrences(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to find all occurrences of an item in an array and return their indices")

	a := []int{1, 2, 3, 4, 5, 6, 1, 7}
	b := []int{0, 6}

	a = findAllOccurrences(a, 1)
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}
