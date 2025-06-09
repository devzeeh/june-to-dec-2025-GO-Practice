package week2

import (
	"fmt"
)

func showDataType() string {
	var name string = "John"
	var gpa float64 = 1.63
	var yearLevel int = 3
	var enrolled bool = true

	result := fmt.Sprintf("Name of student: %s, GPA: %.2f, Year Level: %d, Enrolled: %t", name, gpa, yearLevel, enrolled)
	return result
}

func Day2() string {
	return showDataType()
}
