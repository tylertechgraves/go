// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	if len(s) == 0 {
		return ""
	}

	s = removePunctuation(s)
	words := strings.Fields(s)
	abbreviation := ""

	for i := range words {
		abbreviation += string(words[i][0])

	}

	return strings.ToUpper(abbreviation)
}

func removePunctuation(s string) string {
	returnString := ""
	s = strings.ReplaceAll(s, "-", " ")

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			returnString = returnString + string(r)
		}
	}

	return returnString
}
