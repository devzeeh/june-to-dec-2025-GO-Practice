package week2

import (
	"fmt"
	"strconv"
)

var pi = 3.14
var number = 4.29
var result = pi + number 

func week2Result() string{
	const day = "5"
	const name = "Devzeeh"
	const age = 21
	const month = "june"
	const week = "2"

	var result1 = result
	//fmt.Println("Result is:", strconv.FormatFloat(float64(result1),'f', 1, 64))
	result := fmt.Sprintf("Week %s Day %s of %s Result is: %s", week, day, month, strconv.FormatFloat(float64(result1),'f', 1, 64))
	return result
}

func Day3toDay5() string {
	return week2Result()
}