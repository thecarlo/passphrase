package main

import "testing"

func TestRandomNumber(t *testing.T) {
	for range 1000 {
		n := randomNumber()
		if n < 1 || n > 100 {
			t.Errorf("randomNumber() = %d, want value in [1, 100]", n)
		}
	}
}
