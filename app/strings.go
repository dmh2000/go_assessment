package goassessment

// write a function that reduces adjacent repeated characters
func reduceString(s string, count int) string {

	var i int
	var j int
	var a byte
	var r byte
	var b []byte

	b = make([]byte, 0, len(s))
	// append and continue
	a = s[0]
	b = append(b, a)
	r = a
	j = 1
	i = 1
	for i < len(s) {
		a = s[i]
		// if repeated character
		if a == r {
			// count reached?
			if j == count {
				// yes, skip it
			} else {
				// no, append it
				b = append(b, a)
				j++
			}
		} else {
			// different or count exceeded
			b = append(b, a)
			r = a
			j = 1
		}
		i++
	}

	return string(b)
}

// write a function that wraps lines at a given number of columns without breaking works
func wordWrap(s string, column int) string {
	var q []string
	var a string
	var b string
	var r string
	// find words and add them to a queue
	a = ""
	q = make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			a += string(s[i])
		} else {
			if len(a) > 0 {
				q = append(q, a)
				a = ""
			}
		}
	}
	if len(a) > 0 {
		q = append(q, a)
	}

	a = q[0]
	b = ""
	r = ""
	for i := 1; i < len(q); i++ {
		r = q[i]
		if (len(a) + len(r)) >= column {
			// add wrap character
			b += (a + "_")
			a = r
		} else {
			// continue with space between
			a += " " + r
		}
	}
	// if a word is left unappended, add it
	if len(a) > 0 {
		b += a
	}
	return b
}

// write a function that reverses the characters in a string
func reverseString(s string) string {
	var i int
	var j int
	var b []byte

	if len(s) <= 1 {
		return s
	}

	i = 0
	j = len(s) - 1
	b = []byte(s)
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
