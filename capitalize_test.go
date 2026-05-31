package main

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"already", "Already"},
		{"Hello", "Hello"},
		{"a", "A"},
		{"", ""},
	}
	for _, tt := range tests {
		got := capitalize(tt.input)
		if got != tt.want {
			t.Errorf("capitalize(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
