package main

import "math/rand/v2"

func suggestions(wordCount int) ([]string, error) {
	posCount := 2 * wordCount
	result := make([]string, 5)

	for i := 0; i <= 1; i++ {
		words, err := generateWords(wordCount)
		if err != nil {
			return nil, err
		}
		result[i] = generatePassphrase(words)
	}

	for i := 2; i <= 4; i++ {
		words, err := generateWords(wordCount)
		if err != nil {
			return nil, err
		}
		pos := rand.IntN(posCount)
		result[i] = insertNumber(words, pos, randomNumber())
	}

	return result, nil
}
