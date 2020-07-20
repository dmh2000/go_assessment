package goassessment

func listFilesR(data Dir, dirName string) []string {
	var s []string

	s = append(s, data.files...)
	for _, v := range data.dirs {
		s = append(s, listFilesR(v, dirName)...)
	}
	return s
}

func findDirR(data Dir, dirName string) []string {
	var s []string

	if data.name == dirName {
		s = listFilesR(data, dirName)
	} else {
		for _, v := range data.dirs {
			s = append(s, findDirR(v, dirName)...)
		}
	}
	return s
}

// write a function that returns a list of files starting from the named directory
func listFiles(data Dir, dirName string) []string {
	var s []string

	if dirName == "" {
		s = listFilesR(data, dirName)
	} else {
		s = findDirR(data, dirName)
	}
	return s
}

// Topcoder solution
func generate(k int, arr []int, acc *[][]int) {
	if k > 1 {
		for i := 0; i < k; i++ {
			t := arr[i]
			arr[i] = arr[k-1]
			arr[k-1] = t

			generate(k-1, arr, acc)

			t = arr[i]
			arr[i] = arr[k-1]
			arr[k-1] = t
		}
	} else {
		*acc = append(*acc, arr)
	}
}

// wrote a function that returns all permutations of the input array
func permute(arr []int) [][]int {
	var acc [][]int

	acc = make([][]int, 0)

	// generate permutations
	generate(len(arr), arr, &acc)

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

func parens(cur string, open int, close int, n int, acc *[]string) {
	// from leetcode solution
	if len(cur) == (n * 2) {
		*acc = append(*acc, cur)
		return
	}

	if open < n {
		parens(cur+"(", open+1, close, n, acc)
	}
	if close < open {
		parens(cur+")", open, close+1, n, acc)
	}
}

// write a function that returns an array of strings with all valid sets of n parens
func validParentheses(n int) []string {
	var acc []string

	acc = []string{}

	parens("", 0, 0, n, &acc)

	return acc
}
