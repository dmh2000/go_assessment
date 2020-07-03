package goassessment

// write a function that returns true if either argument is true
func either(a bool, b bool) bool {
	return a || b
}

// write a function that returns true only if both arguments are true
func both(a bool, b bool) bool {
	return a && b
}

// write a function that returns true only if both arguments are false
func none(a bool, b bool) bool {
	return (!a) && (!b)
}
