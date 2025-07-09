package week2

import (
	"fmt"
)

func LoopOverMap() (string, string) {
	fruitDetails := map[string]string{
		"Banana": "20",
		"Apple": "2",
		"Melon": "50",
		"lyche": "54",
	}

	fruit := fmt.Sprintln("\nOriginal:",fruitDetails)

	//loop over map
	var result string
	for ky, val := range fruitDetails {
		result += fmt.Sprintf("%v: %v, ", ky, val)
	}
	var res = result

	return fruit, res

	// if fruit value is lower than 30 print "low: {key}
	/*for key, value := range fruits {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Invalid number for key:", value)
			continue
		}
		if num <= 20 {
			fmt.Println("Low:",key)
		}
	}*/
}
