package week2

import (
	"fmt"
)

func Maphonebook() (string, string, string, string) {
	var name, number string
	phonebook := map[string]string{
		"Angela":  "121",
		"Mikaela": "910",
		"Freya":   "456",
	}
	originalPhonebook := fmt.Sprintln("Original Phonebook: ", phonebook)

	// adding new phonebook
	name = "Ashera"
	number = "069"
	add := fmt.Sprintln("Add:", name, number)

	//check if name is exist (key:)
	_, ok := phonebook[name]
	var exist string
	var newPhonebook string
	if ok == true {
		exist += fmt.Sprintln("exist", ok)
	} else {
		//if not exist add the new phonebook
		phonebook[name] = number

		for key, value := range phonebook {
			newPhonebook += fmt.Sprintln(key, value)
		}
	}

	return originalPhonebook, add, exist, newPhonebook
}
