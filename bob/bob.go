// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimRight(remark, " ")

	// Return "Fine. Be that way!" if empty string or all spaces
	if len(remark) == 0 || isAllSpaces(remark) {
		return "Fine. Be that way!"
	}

	// Return "Sure." if there are no letters and ends in "?"
	if !containsLetters(remark) && strings.HasSuffix(remark, "?") {
		return "Sure."
	}

	// Return "Whatever." if no letters were sent
	if !containsLetters(remark) {
		return "Whatever."
	}

	// Return "Calm down, I know what I'm doing!" if ALL CAPS and ends in "?"
	if strings.HasSuffix(remark, "?") && isUpper(remark) {
		return "Calm down, I know what I'm doing!"
	}

	// Return "Sure." if it ends in "?"
	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}

	// Return "Whoa, chill out!" if ALL CAPS
	if isUpper(remark) {
		return "Whoa, chill out!"
	}

	// Return "Whatever." for anything else
	return "Whatever."
}

func isAllSpaces(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func containsLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}

	return false
}

func isUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}

	return true
}
