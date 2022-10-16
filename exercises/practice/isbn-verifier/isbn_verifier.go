package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	counter, isbnSum := 0, 0
	isbnCleared := strings.Replace(isbn, "-", "", 13)
	if len(isbnCleared) != 10 {
		return false
	}

	for i := 0; i < 10; i++ {
		num, err := strconv.Atoi(string(isbnCleared[i]))
		// Handle the case when the last element could be X
		// which represents the number 10
		if err != nil {
			if i != 9 {
				continue
			}
			if isbnCleared[i] == 'X' {
				num = 10
			} else {
				return false
			}
		}
		counter++
		isbnSum += num * (11 - counter)
	}
	return isbnSum%11 == 0
}
