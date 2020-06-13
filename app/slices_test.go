package goassessment

import (
	"sort"
	"testing"
)

func TestIndexOf(t *testing.T) {
	t.Log("you should be able to determine the location of an item in a slice")

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

func TestSum(t *testing.T) {
	t.Log("you should be able to sum the items of a slice")
	a := []int{1, 2, 3, 4}

	v := sum(a)

	if v != 10 {
		t.Error(shouldBe(v, 10))
	}
}

func TestRemove(t *testing.T) {
	t.Log("you should be able to remove all instances of a value from a slice")

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

func TestFront(t *testing.T) {
	t.Log("you should be able to get the value of the first element of a slice")

	a := []int{1, 2, 3, 4}

	v := front(a)

	if v != 1 {
		t.Error(shouldBe(v, 1))
	}
}

func TestBack(t *testing.T) {
	t.Log("you should be able to get the value of the last element of a slice")

	a := []int{1, 2, 3, 4}

	v := back(a)

	if v != 4 {
		t.Error(shouldBe(v, 1))
	}
}

func TestPushBack(t *testing.T) {
	t.Log("you should be able to add an item to the end of a slice")

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

func TestPopBack(t *testing.T) {
	t.Log("you should be able to remove the last item of a slice")

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

func TestPushFront(t *testing.T) {
	t.Log("you should be able to add an item to the front of a slice")

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

func TestPopFront(t *testing.T) {
	t.Log("you should be able to remove the first item of a slice")

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

func TestConcat(t *testing.T) {
	t.Log("you should be able to join together two slices")
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

func TestInsert(t *testing.T) {
	t.Log("you should be able to add an item anywhere in a slice")

	a := []int{1, 2, 3, 4}

	a = insert(a, 5, 2)

	b := []int{1, 2, 5, 3, 4}
	if len(a) != len(b) {
		t.Error(shouldBe(len(a), len(b)))
	}
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}

}

func TestCount(t *testing.T) {
	t.Log("you should be able to count the occurences of an item in a slice")

	a := []int{1, 4, 2, 3, 4, 4}

	v := count(a, 4)

	if v != 3 {
		t.Error(shouldBe(v, 3))
	}
}

func TestDuplicates(t *testing.T) {
	t.Log("you should be able to find duplicates in a slice")

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

func TestSquare(t *testing.T) {
	t.Log("you should be able to square each number in a slice")

	a := []int{1, 2, 3, 4}
	b := []int{1, 4, 9, 16}

	a = square(a)
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}

func TestFindAllOccurrences(t *testing.T) {
	t.Log("you should be able to find all occurrences of an item in an array and return their indices")

	a := []int{1, 2, 3, 4, 5, 6, 1, 7}
	b := []int{0, 6}

	a = findAllOccurrences(a, 1)
	if !testIntSliceEqual(a, b) {
		t.Error(shouldBe(a, b))
	}
}
