// Package week4 demonstrates basic error handling in Go.
package week4

import (
	"errors"
	"fmt"
)

// checkNumber returns an error if the number is negative.
func checkNumber(num int) (string, error) {
	if num < 0 {
		return "", errors.New("number is negative")
	}
	return "number is positive", nil
}

// Error calls checkNumber and returns the result or error message.
func Error() (string, string) {
	result, err := checkNumber(1)

	if err != nil {
		return "", fmt.Sprintln("Error:", err)
	}
	return "", fmt.Sprintln(result)
}
