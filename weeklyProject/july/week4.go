package july

import (
	_ "errors"
	"fmt"
	"strconv"
)

func Week4Result() string{
	var input string
	fmt.Print("int: ")
	fmt.Scan(&input) // Read input from the user (input will be stored as a string)

	// Try to convert the string input to an integer using strconv.Atoi
	// This returns the integer value and an error (if conversion fails)
	number, err := strconv.Atoi(input)
	if err != nil {
		// If the conversion fails, print an error message and exit
		fmt.Println("Error", err)
		return ""
	}
	// If conversion is successful, print the valid integer
	fmt.Println("Valid", number)

	return ""
}
