package goassessment

import (
	"fmt"
	"testing"
)

func x(s string) string {
	return s
}

func TestX(t *testing.T) {
	s := fmt.Sprintf("%T", x)
	fmt.Println(s)
}
