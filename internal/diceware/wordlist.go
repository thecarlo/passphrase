//go:generate go run ./cmd/gen

package diceware

import (
	"fmt"
	"math/rand/v2"
)

func RandomWords(n int) ([]string, error) {
	if len(words) == 0 {
		return nil, fmt.Errorf("wordlist is empty")
	}
	result := make([]string, n)
	for i := range result {
		result[i] = words[rand.IntN(len(words))]
	}
	return result, nil
}
