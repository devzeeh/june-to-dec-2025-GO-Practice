package week3

import "fmt"

// Person struct defines basic details about a person
type person struct {
	Name       string
	Age        int
	Occupation string
	Salary     int
}

// updateSalary is a method with a pointer receiver that updates the Salary field.
// Since it's a pointer receiver, it modifies the original struct.
func (update *person) updateSalary(newSalary int) {
	update.Salary = newSalary
}

func StructPerson() (string, string, string) {
	// Create a new Person instance with initial values
	person := person{
		Name:       "John",
		Age:        22,
		Occupation: "Developer",
		Salary:     20000,
	}
	// Print the complete Person struct and the current salary
	allStruct := fmt.Sprintln("Complete Struct of 'Person':", person)
	oldSalary := fmt.Sprintln("Old Salary:", person.Salary)

	// Update the salary using the pointer receiver method
	person.updateSalary(50000)
	// Print the updated salary to verify the change
	updatedSalary := fmt.Sprintln("Updated Salary:", person.Salary, "\tPointer", &person.Age)

	return allStruct, oldSalary, updatedSalary
}
