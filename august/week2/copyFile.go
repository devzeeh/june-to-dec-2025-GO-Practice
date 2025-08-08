package week2 // copyFile.go

import (
	"fmt"
	"os"
)

var existingFile = "august/week2/sample.text"

// fileOpen opens the existing file and prints a success message.
func fileOpen() {
	file, err := os.Open(existingFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}

// fileRead reads the content of the existing file and returns it as a string.
func fileRead() string {
	data, err := os.ReadFile(existingFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("File read successfully")
	return string(data)
}

// copyFile creates a copy of the file and writes the provided content into a new file.
func copyFile(content string) {
	// Create a file (this one is just to demonstrate creation)
	file, err := os.Create("august/week2/Filecopy.text")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("New file created successfully")

	// Write to another file
	err = os.WriteFile("august/week2/Filecopy.text", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data written to new file successfully")
	fmt.Println("Content of the new file:", content)
}

func CopyFile() {
	fileOpen()

	content := fileRead() // Read once
	fmt.Println(content)

	copyFile(content) // Pass content instead of reading again
}
