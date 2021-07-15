package raindrops

import "fmt"

func Convert(input int) string {
	returnString := ""

	if input%3 == 0 {
		returnString += "Pling"
	}

	if input%5 == 0 {
		returnString += "Plang"
	}

	if input%7 == 0 {
		returnString += "Plong"
	}

	if len(returnString) == 0 {
		return fmt.Sprint(input)
	}

	return returnString
}
