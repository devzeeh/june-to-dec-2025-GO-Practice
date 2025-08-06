package week2 // Append data to a file

import (
	"fmt"
	"os"
)

// appendDataFile appends a message to the file
func appendDataFile() {
	message := []byte("\nHello, Go!") // Message to append

	file, err := os.OpenFile("august/week2/sample.text", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("Appending %s to file\n", message)

	// Append the message to the file
	_, err = file.Write(message)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message appended successfully")
}

// readTheFile reads and prints the content of the file
func readTheFile() {
	data, err := os.ReadFile("august/week2/sample.text")
	if err != nil {
		panic(err)
	}
	fmt.Println("File read successfully")
	fmt.Println("Content of the file:")
	fmt.Println(string(data))
}

// AppendFileText coordinates appending and reading
func AppendFileText() {
	appendDataFile() // Append the message to the file
	readTheFile()    // Read and print the file content
}
