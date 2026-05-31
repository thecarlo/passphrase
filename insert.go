package main

import (
	"strconv"
	"strings"
)

// Positions are encoded as 0..2N-1:
//   even pos → before word[pos/2]
//   odd  pos → after  word[pos/2]
func insertNumber(words []string, pos int, num int) string {
	result := make([]string, len(words))
	copy(result, words)
	numStr := strconv.Itoa(num)
	wordIdx := pos / 2
	if pos%2 == 0 {
		result[wordIdx] = numStr + result[wordIdx]
	} else {
		result[wordIdx] = result[wordIdx] + numStr
	}
	return strings.Join(result, "-")
}
