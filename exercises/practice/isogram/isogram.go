package isogram

import "strings"

func IsIsogram(word string) bool {
	visited := make(map[rune]bool, 0)
	for _, letter := range strings.ToLower(word) {
		if letter == '-' || letter == ' ' {
			continue
		}
		if visited[letter] {
			return false
		}
		visited[letter] = true
	}
	return true
}
