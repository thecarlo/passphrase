package main

import "testing"

func TestInsertNumber(t *testing.T) {
	tests := []struct {
		words []string
		pos   int
		num   int
		want  string
	}{
		// pos=0 (even): before word[0]
		{[]string{"Apple", "River"}, 0, 42, "42Apple-River"},
		// pos=1 (odd): after word[0]
		{[]string{"Apple", "River"}, 1, 42, "Apple42-River"},
		// pos=2 (even): before word[1]
		{[]string{"Apple", "River"}, 2, 42, "Apple-42River"},
		// pos=3 (odd): after word[1]
		{[]string{"Apple", "River"}, 3, 42, "Apple-River42"},
		// three words
		{[]string{"Apple", "River", "Stone"}, 4, 7, "Apple-River-7Stone"},
		{[]string{"Apple", "River", "Stone"}, 5, 7, "Apple-River-Stone7"},
	}
	for _, tt := range tests {
		got := insertNumber(tt.words, tt.pos, tt.num)
		if got != tt.want {
			t.Errorf("insertNumber(%v, %d, %d) = %q, want %q", tt.words, tt.pos, tt.num, got, tt.want)
		}
	}
}
