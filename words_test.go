package main

import (
	"testing"
	"unicode"
)

func TestGenerateWords(t *testing.T) {
	for _, n := range []int{2, 3, 4, 5} {
		words, err := generateWords(n)
		if err != nil {
			t.Fatalf("generateWords(%d) returned error: %v", n, err)
		}
		if len(words) != n {
			t.Errorf("generateWords(%d) returned %d words, want %d", n, len(words), n)
		}
		for _, w := range words {
			if len(w) == 0 {
				t.Errorf("generateWords(%d) returned an empty word", n)
				continue
			}
			if !unicode.IsUpper(rune(w[0])) {
				t.Errorf("generateWords(%d): word %q is not capitalized", n, w)
			}
		}
	}
}
