package goassessment

func findDirectory(data *Dir, dirName string) *Dir {
	if data.name == dirName {
		return data
	}

	var d *Dir = nil
	for _, v := range data.dirs {
		d = findDirectory(&v, dirName)
		if d != nil {
			break
		}
	}
	return d
}

// write a function that returns a list of files starting either at
// the top level directory or a specified directory
// dirName == 0 means use top level
func listFiles(data Dir, dirName string) []string {

	files := []string{}
	var target *Dir = &data
	if dirName == "" {
		// start  from here
		target = &data
	} else {
		// find directory
		for _, v := range target.dirs {
			d := findDirectory(&v, dirName)
			if d != nil {
				target = d
				break
			}
		}
	}
	if target == nil {
		return []string{}
	}

	for _, v := range target.files {
		files = append(files, v)
	}

	for _, v := range target.dirs {
		e := listFiles(v, dirName)
		files = append(files, e...)
	}
	return files
}

func swap(a []int, x int, y int) {
	u := a[x]
	a[x] = a[y]
	a[y] = u
}

// Generating permutation using Heap Algorithm
// Geek to Geek
func heapPermutation(acc *[][]int, arr []int, size int, n int) {
	// if size becomes 1 then prints the obtained
	// permutation
	if size == 1 {
		r := append([]int{}, arr...)
		*acc = append(*acc, r)
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(acc, arr, size-1, n)

		// if size is odd, swap 0th i.e (first) and
		// (size-1)th i.e (last) element
		if size%2 == 1 {
			swap(arr, 0, size-1)
		} else {
			// If size is even, swap ith and
			// (size-1)th i.e (last) element
			swap(arr, i, size-1)
		}
	}
}

// wrote a function that returns all permutations of the input array
func permute(arr []int) [][]int {
	acc := [][]int{}
	n := len(arr)
	heapPermutation(&acc, arr, len(arr), n)
	return acc
}

// write a function that returns the fibonnaci number of n
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func parens(acc *[]string, s *[]byte,
	pos int, n int, open int, close int) {

	if close == n {
		*acc = append(*acc, string(*s))
		return
	}

	if open > close {
		(*s)[pos] = ')'
		parens(acc, s, pos+1, n, open, close+1)
	}

	if open < n {
		(*s)[pos] = '('
		parens(acc, s, pos+1, n, open+1, close)
	}

}

// write a function that returns an array of strings with all valid sets of n parens
func validParentheses(n int) []string {
	r := []string{}
	s := make([]byte, n*2)
	parens(&r, &s, 0, n, 0, 0)
	return r
}
