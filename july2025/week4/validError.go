package week4

import (
	"errors"
	"fmt"
	"strings"
)

func validateUsername(username string) error {
	username = strings.TrimSpace(username)

	if username == "" {
		return errors.New("Username cannot be empty")
	}
	if len(username) < 3 {
		return errors.New("Username must be atleast 3 character long")
	}
	return nil
}

func ValidateError() (string, string) {
	input := "DevZeeh"

	// call the validation function
	if err := validateUsername(input); err != nil {
		return "", ""
	}

	res := fmt.Sprintln("Username Valid")
	return res, ""
}
