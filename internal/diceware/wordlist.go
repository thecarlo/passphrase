package diceware

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand/v2"
	"strings"
)

//go:embed eff_large_wordlist.txt
var wordlistData string

var words []string

func init() {
	scanner := bufio.NewScanner(strings.NewReader(wordlistData))
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) >= 2 {
			words = append(words, parts[len(parts)-1])
		}
	}
}

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
