package week2 // Read and write file

import (
	"fmt"
	"os"
)

// filePath defines the path to the sample text file used for demonstration.
var filePath = "august/week2/sample.text"

// data contains the string that will be written to the file.
var data = ("Sample text read and write file in Golang")

// createFile creates a new file if it does not already exist.
// If the file exists, it skips creation and prints a message.
func createFile() {
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("File already exists, not creating a new one")
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close() // Ensure the file is closed to avoid resource leaks

	fmt.Println("File created:", filePath)
	fmt.Println("File created successfully")
}

// openFile opens the existing file in read-write mode.
// This function ensures that the file can be accessed for both reading and writing.
func openFile() {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}

// writeFile writes the predefined data string to the file.
// If writing fails, the function panics.
func writeFile() {
	data := []byte(data) // Convert string to byte slice for writing
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("File written successfully")
}

// readFile reads the contents of the file and prints them to the console.
// It handles any read errors by panicking.
func readFile() {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("File read successfully")
	fmt.Println("Content of the file:", string(data))
}

// main is the entry point of the program.
// It sequentially calls the functions to create, open, write to, and read from a file.
func OpenReadWrite() {
	createFile() // Create the file if it does not exist
	openFile()   // Open the file for reading and writing
	writeFile()  // Write data to the file
	readFile()   // Read the contents of the file
}
