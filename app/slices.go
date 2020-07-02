package goassessment

// SLICES
// For those functions that take a slice as input and return a slice,
// you can either modify the input slice or make a copy for the return
// the tests don't require that the input slice is not modified

// write a function that returns the index of an item in a slice
func indexOf(a []int, item int) int {
	var index int

	index = -1
	for i := range a {
		if a[i] == item {
			index = i
			break
		}
	}
	return index
}

// write a function that sums the values in a slice
func sum(a []int) int {
	var i int

	i = 0
	for _, v := range a {
		i += v
	}
	return i
}

// write a function that removes all instances of a value from a slice
func remove(a []int, item int) []int {
	var b []int
	for _, v := range a {
		if v != item {
			b = append(b, v)
		}
	}
	return b
}

// write a function that returns the value of the first element in a slice (wihtout removing it)
func front(a []int) int {
	return a[0]
}

// write a function that returns the value of the last element in a slice (wihtout removing it)
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
	if len(a) <= 1 {
		return a
	}
	return a[0 : len(a)-1]
}

// write a function that adds an item to the front of a slice
func pushFront(a []int, item int) []int {
	a = append([]int{item}, a...)
	return a
}

// write a function that removes an item from the front of a slice
func popFront(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	return a[1:]
}

// write a function that concatenates two slices
func concat(a []int, b []int) []int {
	return append(a, b...)
}

// write a function that adds an item to a slice at the specified index
func insert(a []int, item int, index int) []int {
	if len(a) < index {
		return append(a, item)
	}
	var b []int
	b = make([]int, len(a)+1)
	i := 0
	j := 0
	for i < len(b) {
		if i == index {
			b[i] = item
		} else {
			b[i] = a[j]
			j++
		}
		i++

	}
	return b
}

// write a function that returns a count of matching items in a slice
func count(a []int, item int) int {
	var j int
	j = 0
	for i := 0; i < len(a); i++ {
		if item == a[i] {
			j++
		}
	}
	return j
}

// write a function that finds duplicates in a slice
func duplicates(a []int) []int {
	var m map[int]int
	var b []int

	m = make(map[int]int)
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}

	b = []int{}
	for key, val := range m {
		if val > 1 {
			b = append(b, key)
		}
	}

	return b
}

// write a function that sqaures all items in a slice
func square(a []int) []int {
	for i := range a {
		a[i] = a[i] * a[i]
	}
	return a
}

// write a function that returns all the indices in a slice that matches an item
func findAllOccurrences(a []int, item int) []int {
	var b []int

	b = make([]int, 0)
	for i, v := range a {
		if v == item {
			b = append(b, i)
		}
	}
	return b
}
