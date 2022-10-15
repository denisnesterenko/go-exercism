// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"unicode"
)

// Transforms trimmed sentence to the abbreviation
func Abbreviate(s string) string {
	var result string
	wordEnded := true
	for _, char := range s {
		if unicode.IsLetter(char) && wordEnded {
			result += string(unicode.ToUpper(char))
			wordEnded = false
		} else if char == ' ' || char == '-' {
			wordEnded = true
		} else {
			continue
		}
	}
	return result
}
