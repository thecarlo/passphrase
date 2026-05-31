package diceware

import "testing"

const expectedWordCount = 7776

func TestWordlistCount(t *testing.T) {
	if got := len(words); got != expectedWordCount {
		t.Errorf("wordlist has %d words, want %d", got, expectedWordCount)
	}
}

// Spot-check the first few and last few entries against the EFF wordlist.
// Dice roll 11111 = "abacus", 11112 = "abdomen", 66666 = "zoom".
func TestWordlistKnownEntries(t *testing.T) {
	tests := []struct {
		index int
		want  string
	}{
		{0, "abacus"},
		{1, "abdomen"},
		{2, "abdominal"},
		{3, "abide"},
		{4, "abiding"},
		{5, "ability"},
		{expectedWordCount - 1, "zoom"},
	}
	for _, tt := range tests {
		if got := words[tt.index]; got != tt.want {
			t.Errorf("words[%d] = %q, want %q", tt.index, got, tt.want)
		}
	}
}

func TestRandomWords(t *testing.T) {
	for _, n := range []int{1, 2, 3, 5} {
		result, err := RandomWords(n)
		if err != nil {
			t.Fatalf("RandomWords(%d) returned error: %v", n, err)
		}
		if len(result) != n {
			t.Errorf("RandomWords(%d) returned %d words, want %d", n, len(result), n)
		}
		for _, w := range result {
			if w == "" {
				t.Errorf("RandomWords(%d) returned an empty word", n)
			}
		}
	}
}

func TestRandomWordsDistribution(t *testing.T) {
	// Draw a large sample and verify all words come from the wordlist.
	wordSet := make(map[string]struct{}, len(words))
	for _, w := range words {
		wordSet[w] = struct{}{}
	}

	result, err := RandomWords(1000)
	if err != nil {
		t.Fatalf("RandomWords(1000) returned error: %v", err)
	}
	for _, w := range result {
		if _, ok := wordSet[w]; !ok {
			t.Errorf("RandomWords returned word %q not in wordlist", w)
		}
	}
}
