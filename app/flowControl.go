package goassessment

import "strconv"

// write a function that receives a number as its argument;
// if the number is divisible by 3, the function should return 'fizz';
// if the number is divisible by 5, the function should return 'buzz';
// if the number is divisible by 3 and 5, the function should return
// 'fizzbuzz';
//
// otherwise the function should return the number as a string
func fizzBuzz(num int) string {
	var r string
	if (num%3) == 0 && (num%5) == 0 {
		r = "fizzbuzz"
	} else if (num % 5) == 0 {
		r = "buzz"
	} else if (num % 3) == 0 {
		r = "fizz"
	} else {
		r = strconv.Itoa(num)
	}
	return r
}
