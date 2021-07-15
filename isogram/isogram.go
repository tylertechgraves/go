package isogram

import (
	"strings"
)

func IsIsogram(input string) bool {
	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	if len(input) == 0 {
		return true
	}

	for i := 0; i < len(input); i++ {
		if strings.LastIndex(input, string(input[i])) != i {
			return false
		}
	}

	return true
}
