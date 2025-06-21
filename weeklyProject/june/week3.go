package june

import "fmt"

// Function to generate the first 10 Fibonacci numbers
func PrintFibonacci() string {
	// Check if n is less than or equal to 0
	a, b := 0, 1 // Initialize the first two Fibonacci numbers
	n := 10      // Number of Fibonacci numbers to generate
	result := ""
	// for loop to generate Fibonacci numbers
	for i := 0; i < n; i++ {
		result += fmt.Sprintf("%d", a)
		if i < n-1 {
			result += " "
		}
		a, b = b, a+b
	}
	return result
}

// Called by the main function to get the result for week 3
// This function returns the first 10 Fibonacci numbers as a string
func Week3Result() string {
	return PrintFibonacci()
}
