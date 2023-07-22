package main

import (
	"fmt"
	"os"
)

func main() {
	error := validateInput(os.Args[1:])
	if error != nil {
		fmt.Println("Something went wrong: ", error)
		return
	}

	name := reorderNameByCountryCode(os.Args[1:])
	fmt.Println("Ordered name is: ", name)
}

func validateInput(argsWithoutProg []string) error {
	if len(argsWithoutProg) < 3 || len(argsWithoutProg) > 4 {
		return fmt.Errorf("your input should have at least 3 args and 4 args at maximum")
	}

	var lengthOfArgs = len(os.Args)
	var countryCode = os.Args[lengthOfArgs-1]
	if len(countryCode) != 2 {
		return fmt.Errorf("invalid country code")
	}

	return nil
}

func reorderNameByCountryCode(argsWithoutProg []string) string {
	var argsLen = len(argsWithoutProg)
	var countryCode = argsWithoutProg[argsLen-1]
	var firstName = argsWithoutProg[0]
	var lastName = argsWithoutProg[1]
	var middleName = argsWithoutProg[2]
	var hasMiddleName = middleName != countryCode
	var nameWithoutMiddle = firstName + " " + lastName

	switch countryCode {
	case "VN":
		if hasMiddleName {
			return lastName + " " + middleName + " " + firstName
		} else {
			return lastName + " " + firstName
		}
	case "US":
		if !hasMiddleName {
			return nameWithoutMiddle
		} else {
			return firstName + " " + lastName + " " + middleName
		}
	default:
		return nameWithoutMiddle
	}

	return nameWithoutMiddle
}
