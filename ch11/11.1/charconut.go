package charcount

import (
	"unicode"
)

// Charcount returns count of Unicode chars from string input
func Charcount(input string) int {
	var count int
	for _, r := range input {
		if r == unicode.ReplacementChar {
			continue
		}
		count++
	}
	return count
}

//!-
