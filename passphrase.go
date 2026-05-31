package main

import "strings"

func generatePassphrase(words []string) string {
	return strings.Join(words, "-")
}
