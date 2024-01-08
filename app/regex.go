package goassessment

import (
	"regexp"
)

// write a function that checks if the string contains a decimal number
func containsNumber(s string) bool {
	m, err := regexp.Match(`.*[0-9].*`, []byte(s))
	if err != nil {
		return false
	}

	return m
}

// write a function that checks if the string ends with a vowel
func endsWithVowel(s string) bool {
	m, err := regexp.Match(`.*[aeiouAEIOU]$`, []byte(s))
	if err != nil {
		return false
	}

	return m
}

// write a function that captures and returns the first string of 3 decimal digits
func captureThreeNumbers(s string) string {
	re := regexp.MustCompile(`([0-9][0-9][0-9])`)
	m := re.FindSubmatch([]byte(s))
	if len(m) == 0 {
		return ""
	}

	return string(m[1])
}

// write a function that checks if the string matches a specified pattern
func matchesPattern(s string) bool {
	// XXX-XXX-XXXX
	re := regexp.MustCompile(`^\d\d\d\-\d\d\d\-\d\d\d\d$`)
	m := re.Match([]byte(s))
	return m
}

// write a function that checks if the string is a correctly-formatted monetary amounts in USD
func isUSD(s string) bool {
	re := regexp.MustCompile(`^\$(0|[1-9][0-9]{0,2})(,\d{3})*(\.\d{2})?$`)
	m := re.Match([]byte(s))
	return m
}
