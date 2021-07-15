package reverse

func Reverse(input string) string {
	reversedString := ""

	for i := len(input) - 1; i >= 0; i-- {
		reversedString += string(input[i])
	}

	return reversedString
}
