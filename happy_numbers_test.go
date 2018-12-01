package main

import "testing"

var numbers = []struct {
	in  int
	out bool
}{
	{7, true},
	{8, false},
}

func TestIsHappyNumber(t *testing.T) {
	// It is a bit hard to test this function. I did not find
	// a way to mathematically prove that 7 is a number happy and 8 not.
	for _, tt := range numbers {
		r := IsHappyNumber(tt.in)
		if r != tt.out {
			t.Errorf("For number %d, expected %t but got %t", tt.in, tt.out, r)
		}
	}
}
