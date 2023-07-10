package goassessment

import (
	"strings"
)

// write a function that composes a string from arguments
func composeString(a string, b string) string {
	return a + ", " + b + "\n"
}

// write a function that returns a string from a byte array
func fromBytes(b []byte) string {
	return string(b)
}

// write a function that splits a string into words
func splitString(s string) []string {
	return strings.Split(s, ", ")
}

// write a function that converts a string to Title Case
func titleString(s string) string {
	return strings.Title((s))
}

func runesFromString(s string) []rune {
	return []rune(s)
}

// write a function that reduces adjacent repeated characters
func reduceString(s string, count int) string {
	a := []rune(s)
	b := a[0]
	r := string(b)
	n := 1
	for _, v := range a[1:] {
		if b != v {
			b = v
			r += string(b)
			n = 1
		} else if n < count {
			// accumulate
			r += string(v)
			n++
		} else {
			// full count, discard
		}
	}

	return r
}

// write a function that wraps lines at a given number of columns without breaking words
func wordWrap(s string, column int) string {
	words := strings.Split(s, " ")
	line := ""
	line_len := 0
	sep := ""
	state := 0
	for _, v := range words {
		switch state {
		case 0:
			// first line
			line_len = len(v)
			line += v
			sep = "_"
			state = 1
			break
		case 1:
			// continue line
			if (line_len + len(v)) < (column - 1) {
				line += " "
				line += v
			} else {
				line += sep
				line += v
				line_len = len(v)
			}
		}
	}
	return line
}

// write a function that reverses the characters in a string
func reverseString(s string) string {
	r := ""
	a := []rune(s)
	for i := len(a) - 1; i >= 0; i-- {
		r += string(a[i])
	}
	return r
}
