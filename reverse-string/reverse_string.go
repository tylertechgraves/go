package reverse

func Reverse(s string) string {
	runeSlice := []rune(s)
	var reversedRuneSlice []rune

	for i := len(runeSlice) - 1; i > -1; i-- {
		reversedRuneSlice = append(reversedRuneSlice, runeSlice[i])
	}

	return string(reversedRuneSlice)
}
