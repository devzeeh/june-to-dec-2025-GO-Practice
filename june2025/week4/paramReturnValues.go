package week4

import (
	// Importing strconv for converting float to string
	"strconv"
)

func ParamReturnValues(number1 int, number2 float64) (message string) {
	result := float64(number1) * number2
	message = "The result is: " + strconv.FormatFloat(result, 'f', 2, 64)
	return message
}
