package main

import "testing"

func TestGeneratePassphrase(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{[]string{"Apple", "River"}, "Apple-River"},
		{[]string{"Apple", "River", "Stone"}, "Apple-River-Stone"},
		{[]string{"One"}, "One"},
		{[]string{"Apple", "River", "Stone", "Cloud"}, "Apple-River-Stone-Cloud"},
	}
	for _, tt := range tests {
		got := generatePassphrase(tt.words)
		if got != tt.want {
			t.Errorf("generatePassphrase(%v) = %q, want %q", tt.words, got, tt.want)
		}
	}
}
