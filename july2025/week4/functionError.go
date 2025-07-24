package week4

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("file not found")
}

func startApp() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("failed to start app: %w", err) // ← wrapping here
	}
	return nil
}

func FunctionError() string {
	err := startApp()
	if err != nil {
		return fmt.Sprintln("Error:", err)
	}
	return ""
}
