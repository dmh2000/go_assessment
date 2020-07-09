package goassessment

import "testing"

// write a function that checks if the string contains a decimal number
func TestContainsNumber(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to detect a number in a string")
	var b bool

	b = containsNumber("abc123")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = containsNumber("abc")
	if b {
		t.Error(shouldBe(b, false))
	}
}

// write a function that checks if the string contains repeated letters
func TestRepeatingLetter(t *testing.T) {
	t.Log("you should be able to detect a repeating letter in a string")
	var b bool

	b = containsRepeatingLetter("bookkeeping")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = containsRepeatingLetter("rattler")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = containsRepeatingLetter("ZEPPELIN")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = containsRepeatingLetter("cats")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = containsRepeatingLetter("l33t")
	if b {
		t.Error(shouldBe(b, false))
	}
}

// write a function that checks if the string ends with a vowel
func TestEndsWithVowel(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to determine whether a string ends with a vowel (aeiou)")
	var b bool

	b = endsWithVowel("cats")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = endsWithVowel("gorilla")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = endsWithVowel("I KNOW KUNG FU")
	if !b {
		t.Error(shouldBe(b, true))
	}
}

// write a function that captures and returns the first string of 3 decimal digits
func TestCaptureThreeNumbers(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to capture the first series of three numbers")
	var s string
	var r string

	s = "abc123"
	r = "123"
	s = captureThreeNumbers(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "9876543"
	r = "987"
	s = captureThreeNumbers(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "abcdef"
	r = ""
	s = captureThreeNumbers(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}

	s = "12ab12ab"
	r = ""
	s = captureThreeNumbers(s)
	if s != r {
		t.Error(shouldBe(s, r))
	}
}

// write a function that checks if the string matches a specified pattern
func TestMatchesPattern(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to determine whether a string matches a pattern")
	// the pattern is XXX-XXX-XXXX where all X"s are digits
	var b bool

	b = matchesPattern("800-555-1212")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = matchesPattern("451-933-7899")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = matchesPattern("33-444-5555")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = matchesPattern("abc-def-hijk")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = matchesPattern("1800-555-1212")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = matchesPattern("800-555-12121")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = matchesPattern("800-5555-1212")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = matchesPattern("800-55-1212")
	if b {
		t.Error(shouldBe(b, false))
	}
}

// write a function that checks if the string is a correctly-formatted monetary amounts in USD
func TestIsUSD(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors

	t.Log("you should be able to detect correctly-formatted monetary amounts in USD")
	var b bool
	// TRUE

	b = isUSD("$132.03")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = isUSD("$32.03")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = isUSD("$2.03")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = isUSD("$1,023,032.03")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = isUSD("$20,933,209.93")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = isUSD("$20,933,209")
	if !b {
		t.Error(shouldBe(b, true))
	}

	b = isUSD("$459,049,393.21")
	if !b {
		t.Error(shouldBe(b, true))
	}

	// FALSE
	b = isUSD("34,344.34")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$,344.34")
	if b {
		t.Error(shouldBe(b, false))
	}

	if b {
		t.Error(shouldBe(b, false))
	}

	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$34,344.3")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$34,344.344")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$34,344_34")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$3,432,12.12")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$3,432,1,034.12")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("4$3,432,034.12")
	if b {
		t.Error(shouldBe(b, false))
	}

	b = isUSD("$2.03.45")
	if b {
		t.Error(shouldBe(b, false))
	}
}
