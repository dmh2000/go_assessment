package goassessment

import (
	"testing"
)

// write a function that composes a string from arguments
func TestStringCompose(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to compose a string from arguments")
	var s string
	var r string
	s = composeString("Hello", "World")
	r = "Hello, World\n"

	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// write a function that returns a string from a byte array
func TestStringFromBytes(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to convert a byte array to a string")

	b := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 10}
	s := fromBytes(b)
	r := "Hello, World\n"

	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// write a function that takes a string and returns an array of runes
func TestRunes(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to decompose a UTF-8 string into an array of runes")
	var s string
	var r []rune
	var u []rune

	s = "✋ 👍 👎 ✊"
	u = []rune{9995, 32, 128077, 32, 128078, 32, 9994}
	r = runesFromString(s)

	// are they the same length
	if len(u) != len(r) {
		t.Error(shouldBe(len(u), len(r)))
		return
	}

	// are they equal
	b := true
	for i := 0; i < len(u); i++ {
		if r[i] != u[i] {
			b = false
			break
		}
	}

	if !b {
		t.Error(shouldBe(r, u))
	}
}

// write a function that splits a string into words
func TestSplitString(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to split a string into words")
	var s string
	var r []string
	var u []string
	s = "Hello, World"
	u = []string{"Hello", "World"}

	r = splitString(s)

	// check length equal
	if len(r) != len(u) {
		t.Error(shouldBe(len(r), len(u)))
		return
	}

	// check words
	for i := 0; i < len(u); i++ {
		if u[i] != r[i] {
			t.Error(shouldBe(u, r))
		}
	}
}

// write a function that converts a string to Title Case
func TestTitleCase(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to convert a string to title case")
	var s string
	var r string
	var u string
	s = "this is the title"
	u = "This Is The Title"

	r = titleString(s)

	// check length equal
	if r != u {
		t.Error(shouldBe(r, u))
	}
}

// write a function that reduces repeated characters
func TestReduceString(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to reduce adjacent repeated characters to a desired minimum")
	var s string
	var r string

	s = "aaaabbbb"
	r = "aabb"
	s = reduceString(s, 2)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "xaaabbbb"
	r = "xaabb"
	s = reduceString(s, 2)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "aaaabbbb"
	r = "ab"
	s = reduceString(s, 1)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "aaxxxaabbbb"
	r = "aaxxaabb"
	s = reduceString(s, 2)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "aaxxxaabbbb"
	r = "axab"
	s = reduceString(s, 1)
	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// write a function that wraps lines at a given number of columns without breaking works
func TestWordWrap(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to wrap lines at a given number of columns, without breaking words")
	// !!! use underscore to indicate where a wrap would occur

	column := 5
	s := "abcdef abcde abc def"
	r := "abcdef_abcde_abc_def"
	s = wordWrap(s, column)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "abc abc abc"
	r = "abc_abc_abc"
	s = wordWrap(s, column)
	if s != r {
		t.Error(shouldBe(s, r))
	}
	s = "a b c def"
	r = "a b c_def"

	s = wordWrap(s, column)
	if s != r {
		t.Error(shouldBe(s, r))
	}

}

// write a function that reverses the characters in a string
func TestReverseString(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to reverse the characters in a string")
	var s string
	var r string

	s = "abc"
	r = "cba"
	s = reverseString(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "i am a string of characters"
	r = "sretcarahc fo gnirts a ma i"
	s = reverseString(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}

}
