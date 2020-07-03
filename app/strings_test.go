package goassessment

import "testing"

// write a function that reduces repeated characters
func TestReduceString(t *testing.T) {
	t.Log("you should be able to reduce repeated characters to a desired minimum")

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

}

// write a function that wraps lines at a given number of columns without breaking works
func TestWordWrap(t *testing.T) {
	t.Log("you should be able to wrap lines at a given number of columns, without breaking words")
	var s string
	var r string
	var column int

	// use underscore to indicate where a wrap would occur

	column = 5
	s = "abcdef abcde abc def"
	r = "abcdef_abcde_abc_def"
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
	t.Log("you should be able to reverse the characters in a string")
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

	s = "A man, a plan, a canal: Panama"
	r = "amanaP :lanac a ,nalp a ,nam A"
	s = reverseString(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}
}
