package goassessment

import (
	"regexp"
)

// write a function that checks if the string contains a decimal number
func containsNumber(s string) bool {

	ok, _ := regexp.Match(`\d`, []byte(s))
	return ok
}

// write a function that checks if the string ends with a vowel
func endsWithVowel(s string) bool {
	ok, _ := regexp.Match(`.*[aeiou,AEIOU]$`, []byte(s))
	return ok
}

// write a function that captures and returns the first string of 3 decimal digits
func captureThreeNumbers(s string) string {
	re := regexp.MustCompile(`\d\d\d`)
	m := re.Find([]byte(s))
	return string(m)
}

// write a function that checks if the string matches a specified pattern
func matchesPattern(s string) bool {
	// the pattern is XXX-XXX-XXXX where all X"s are decimal digits
	re := regexp.MustCompile(`^\d\d\d-\d\d\d-\d\d\d\d$`)
	ok := re.Match([]byte(s))
	return ok
}

// write a function that checks if the string is a correctly-formatted monetary amounts in USD
func isUSD(s string) bool {
	re := regexp.MustCompile(`^\$[123456789]([0123456789]{0,2})(\,[0123456789]{3})*(\.\d\d)?$`)
	b := re.Match([]byte(s))

	return b
}
