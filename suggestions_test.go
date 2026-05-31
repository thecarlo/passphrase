package main

import (
	"testing"
	"unicode"
)

func containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func TestSuggestions(t *testing.T) {
	for _, wordCount := range []int{2, 3, 4, 5, 6, 7, 8, 9, 10} {
		result, err := suggestions(wordCount)
		if err != nil {
			t.Fatalf("suggestions(%d) returned error: %v", wordCount, err)
		}
		if len(result) != 5 {
			t.Errorf("suggestions(%d) returned %d results, want 5", wordCount, len(result))
		}
		for i, phrase := range result[:2] {
			if containsDigit(phrase) {
				t.Errorf("suggestions(%d): result[%d] = %q should not contain a digit", wordCount, i, phrase)
			}
		}
		for i, phrase := range result[2:] {
			if !containsDigit(phrase) {
				t.Errorf("suggestions(%d): result[%d] = %q should contain a digit", wordCount, i+2, phrase)
			}
		}
	}
}
