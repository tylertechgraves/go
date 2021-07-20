package luhn

import (
	"strconv"
	"strings"
)

func Valid(s string) bool {
	// Remove all spaces
	s = strings.ReplaceAll(s, " ", "")

	// If the length of the number is less than 2, it's invalid.
	if len(s) < 2 {
		return false
	}

	// Only numbers are allowed, otherwise it's invalid
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	// Start from the right and double every other digit, starting with the second digit
	// If doubling the digit results in something greater than 9, subtract 9 from it
	// Sum all the digits
	counter := 1
	sum := 0
	for i := len(s) - 1; i > -1; i-- {
		if counter%2 == 0 {
			doubleThis, _ := strconv.Atoi(string(s[i]))
			doubleThis = doubleThis * 2
			if doubleThis > 9 {
				doubleThis = doubleThis - 9
			}
			sum += doubleThis
		} else {
			digit, _ := strconv.Atoi(string(s[i]))
			sum += digit
		}
		counter++
	}

	// If the sum is evenly divisible by 10, return true.  Otherwise, return false
	return sum%10 == 0
}
