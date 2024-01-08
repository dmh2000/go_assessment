package goassessment

// SLICES
// For those functions that take a slice as input and return a slice,
// you can either modify the input slice or make a copy for the return
// the tests don't require that the input slice is not modified

// write a function that returns the index of an item in a slice
func indexOf(a []int, item int) int {
	for i, v := range a {
		if v == item {
			return i
		}
	}
	return -1
}

// write a function that sums the values in a slice
func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// write a function that removes all instances of a value from a slice
func remove(a []int, item int) []int {
	b := []int{}
	for _, v := range a {
		if v != item {
			b = append(b, v)
		}
	}
	return b
}

// write a function that returns the value of the first element in a slice (without removing it)
func front(a []int) int {
	return a[0]
}

// write a function that returns the value of the last element in a slice (without removing it)
func back(a []int) int {
	return a[len(a)-1]
}

// write a function that adds an item to the end of a slice
func pushBack(a []int, item int) []int {
	a = append(a, item)
	return a
}

// write a function that removes an item to the end of a slice
func popBack(a []int) []int {
	a = a[0 : len(a)-1]
	return a
}

// write a function that adds an item to the front of a slice
func pushFront(a []int, item int) []int {
	a = append([]int{item}, a...)
	return a
}

// write a function that removes an item from the front of a slice
func popFront(a []int) []int {
	a = a[1:]
	return a
}

// write a function that concatenates two slices
func concat(a []int, b []int) []int {
	c := append(a, b...)
	return c
}

// write a function that adds an item to a slice at the specified index
func insert(a []int, item int, index int) []int {
	if index == 0 {
		return append([]int{item}, a...)
	}
	if index >= len(a) {
		return append(a, item)
	}
	b := []int{}
	b = append(b, a[0:index]...)
	c := []int{}
	c = append(c, a[index:]...)
	d := append(b, item)
	d = append(d, c...)
	return d
}

// write a function that returns a count of matching items in a slice
func count(a []int, item int) int {
	ct := 0
	for _, v := range a {
		if v == item {
			ct++
		}
	}
	return ct
}

// write a function that finds duplicates in a slice
func duplicates(a []int) []int {
	m := make(map[int]int)

	for _, v := range a {
		u := m[v]
		u++
		m[v] = u
	}

	b := []int{}
	for k := range m {
		if m[k] > 1 {
			b = append(b, k)
		}
	}
	return b
}

// write a function that squares all items in a slice
func square(a []int) []int {
	b := []int{}
	for _, v := range a {
		b = append(b, v*v)
	}
	return b
}

// write a function that returns all the indices in a slice that matches an item
func findAllOccurrences(a []int, item int) []int {
	b := []int{}
	for i, v := range a {
		if v == item {
			b = append(b, i)
		}
	}
	return b
}
