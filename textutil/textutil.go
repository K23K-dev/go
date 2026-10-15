// Package textutil has small helpers for working with text. It's the
// example package for lesson 9.
//
// The comment right above the package line, like this one, documents the
// package. "go doc" shows it, and so does VS Code when you hover over the
// import.
package textutil

import (
	"strings"
	"unicode"
)

// Reverse returns s with its characters in reverse order.
// It's exported because its name starts with a capital letter.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case, spaces and punctuation.
func IsPalindrome(s string) bool {
	cleaned := normalize(s)
	return cleaned == Reverse(cleaned)
}

// WordCount returns the number of words in s.
func WordCount(s string) int {
	return len(strings.Fields(s))
}

// normalize lowercases s and keeps only letters and digits.
// Its name starts with a lowercase letter, so only code inside the textutil
// package can call it.
func normalize(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
