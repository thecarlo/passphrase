package main

import "github.com/carlovan/passphrase/internal/diceware"

func generateWords(n int) ([]string, error) {
	ws, err := diceware.RandomWords(n)
	if err != nil {
		return nil, err
	}
	for i, w := range ws {
		ws[i] = capitalize(w)
	}
	return ws, nil
}
