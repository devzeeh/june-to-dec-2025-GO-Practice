package week2

import (
	"fmt"
)
 
func practiceDay1() string {
	// declaring variable using not inferred type
	var name string = "DevZeeh"
	var month string = "june"
	var day int = 1
	var week int = 2
	var message string = "Welcome to GO!,"

	// result for ":="
	result := fmt.Sprintf("%s %s, this is your week %d day %d for the month of %s", message, name, week, day, month)
	return result
}

func Day1() string {
	return practiceDay1()
}
