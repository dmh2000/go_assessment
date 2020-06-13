package goassessment

import "testing"

func TestReduceString(t *testing.T) {
	t.Log("you should be able to reduce duplicate characters to a desired minimum")

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

func TestReverseString(t *testing.T) {
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
