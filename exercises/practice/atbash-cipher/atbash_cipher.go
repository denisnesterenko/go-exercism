package atbash

import "unicode"

func Atbash(s string) string {
	var encodedString string
	counter := 0
	for _, char := range s {
		if char == ' ' || unicode.IsPunct(char) {
			continue
		}
		if counter == 5 {
			encodedString += " "
			counter = 0
		}
		if unicode.IsLetter(char) {
			lowered := unicode.ToLower(char)
			encodedString += string('z' + ('a' - lowered))
		} else {
			encodedString += string(char)
		}
		counter++
	}
	return encodedString
}
