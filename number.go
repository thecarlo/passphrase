package main

import "math/rand/v2"

func randomNumber() int {
	return rand.IntN(100) + 1
}
