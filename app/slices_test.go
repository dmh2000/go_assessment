package goassessment

import (
	"sort"
	"testing"
)

// For those functions that take a slice as input and return a slice,
// you can either modify the input slice or make a copy
// the tests don't check for returning the origina input slice

func TestIndexOf(t *testing.T) {
	t.Log("you should be able to determine the location of an item in a slice")

	a := []int{1, 2, 3, 4}

	// your function
	index := indexOf(a, 3)

	if index != 2 {
		t.Error(shouldBe(index, 2))
	}
	index = indexOf(a, 5)
	if index != -1 {
		t.Error(shouldBe(index, -1))
	}
}

func TestSum(t *testing.T) {
	t.Log("you should be able to sum the items of a slice")
	a := []int{1, 2, 3, 4}

	// your function
	v := sum(a)

	if v != 10 {
		t.Error(shouldBe(v, 10))
	}
}

func TestRemove(t *testing.T) {
	t.Log("you should be able to remove all instances of a value from a slice")

	a := []int{1, 2, 3, 4}
	// make sure the value appears more than one time
	// make sure the value appears more than one time in a row
	a = append(a, 2)
	a = append(a, 2)

	// your function
	a = remove(a, 2)

	answer := []int{1, 3, 4}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestFront(t *testing.T) {
	t.Log("you should be able to get the value of the first element of a slice")

	a := []int{1, 2, 3, 4}

	// your function
	v := front(a)

	if v != 1 {
		t.Error(shouldBe(v, 1))
	}
}

func TestBack(t *testing.T) {
	t.Log("you should be able to get the value of the last element of a slice")

	a := []int{1, 2, 3, 4}

	// your function
	v := back(a)

	if v != 4 {
		t.Error(shouldBe(v, 1))
	}
}

func TestPushBack(t *testing.T) {
	t.Log("you should be able to add an item to the end of an slice")

	a := []int{1, 2, 3, 4}

	// your function
	a = pushBack(a, 2)

	answer := []int{1, 2, 3, 4, 2}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestPopBack(t *testing.T) {
	t.Log("you should be able to remove the last item of an slice")

	a := []int{1, 2, 3, 4}

	// your function
	a = popBack(a)

	answer := []int{1, 2, 3}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestPushFront(t *testing.T) {
	t.Log("you should be able to add an item to the front of an slice")

	a := []int{1, 2, 3, 4}

	// your function
	a = pushFront(a, 2)

	answer := []int{2, 1, 2, 3, 4}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}

}

func TestPopFront(t *testing.T) {
	t.Log("you should be able to remove the last item of an slice")

	a := []int{1, 2, 3, 4}

	// your function
	a = popBack(a)

	answer := []int{2, 3, 4}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestConcat(t *testing.T) {
	t.Log("you should be able to join together two slices")
	t.Log("you should be able to remove the last item of an slice")
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}

	// your function
	a = concat(a, b)

	answer := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestInsert(t *testing.T) {
	t.Log("you should be able to add an item anywhere in a slice")

	a := []int{1, 2, 3, 4}

	// your function
	a = insert(a, 5, 2)

	answer := []int{1, 2, 5, 3, 4}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}

}

func TestCount(t *testing.T) {
	t.Log("you should be able to count the occurences of an item in a slice")

	a := []int{1, 4, 2, 3, 4, 4}

	// your function
	v := count(a, 4)

	if v != 3 {
		t.Error(shouldBe(v, 3))
	}
}

func TestDuplicates(t *testing.T) {
	t.Log("you should be able to find duplicates in a slice")

	a := []int{1, 2, 4, 4, 3, 3, 1, 5, 3}

	// your function
	a = duplicates(a)

	sort.IntSlice(a).Sort()
	answer := []int{1, 3, 4}
	if len(a) != len(answer) {
		t.Error(shouldBe(len(a), len(answer)))
	}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestSquare(t *testing.T) {
	t.Log("you should be able to square each number in a slice")

	a := []int{1, 2, 3, 4}

	// your function
	a = square(a)

	answer := []int{1, 4, 9, 16}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}

func TestFindAllOccurrences(t *testing.T) {
	t.Log("you should be able to find all occurrences of an item in an array and return their indices")
	a := []int{1, 2, 3, 4, 5, 6, 1, 7}

	// your function
	a = findAllOccurrences(a, 1)

	answer := []int{0, 6}
	if !intSliceEqual(a, answer) {
		t.Error(shouldBe(a, answer))
	}
}
