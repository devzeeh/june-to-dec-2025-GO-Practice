package june

import (
	"fmt"
)

// function for adding number
func addNumbers(a float64, b float64) (result float64) {
	result = float64(a + b)
	return result
}

// function for adding number
func subtractNumbers(a float64, b float64) (result float64) {
	result = float64(a - b)
	return result
}

// function for adding number
func multipleNumbers(a float64, b float64) (result float64) {
	result = float64(a + b)
	return result
}

// function for adding number
func divideNumbers(a float64, b float64) (result float64) {
	result = float64(a + b)
	return result
}

func Week4Result() string {
	var number, number1, number2 float64
	fmt.Println("Basic Calculator")
	for {
		fmt.Println("Choose operation: \n Addition(1), Subtraction(2), Multiplication(3), Division(4).")
		fmt.Scanln(&number)

		switch number {
		case 1:
			fmt.Println("Addition")
			fmt.Print("First number: ")
			fmt.Scan(&number1)
			fmt.Print("Second number: ")
			fmt.Scan(&number2)
			fmt.Println("Result is: ", addNumbers(number1, number2))
		case 2:
			fmt.Println("Subtraction")
			fmt.Print("First number: ")
			fmt.Scan(&number1)
			fmt.Print("Second number: ")
			fmt.Scan(&number2)
			fmt.Println("Result is: ", subtractNumbers(number1, number2))
		case 3:
			fmt.Println("Multiplication")
			fmt.Print("First number: ")
			fmt.Scan(&number1)
			fmt.Print("Second number: ")
			fmt.Scan(&number2)
			fmt.Println("Result is: ", multipleNumbers(number1, number2))
		case 4:
			fmt.Println("Division")
			fmt.Print("First number: ")
			fmt.Scan(&number1)
			fmt.Print("Second number: ")
			fmt.Scan(&number2)
			fmt.Println("Result is: ", divideNumbers(number1, number2))
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
		break
	}
	return "Calculator finished successfully"
}
