package diceware

import "testing"

func TestRandomWords(t *testing.T) {
	for _, n := range []int{1, 2, 3, 5} {
		words, err := RandomWords(n)
		if err != nil {
			t.Fatalf("RandomWords(%d) returned error: %v", n, err)
		}
		if len(words) != n {
			t.Errorf("RandomWords(%d) returned %d words, want %d", n, len(words), n)
		}
		for _, w := range words {
			if w == "" {
				t.Errorf("RandomWords(%d) returned an empty word", n)
			}
		}
	}
}

func TestWordlistLoaded(t *testing.T) {
	if len(words) == 0 {
		t.Error("wordlist is empty after init")
	}
}
