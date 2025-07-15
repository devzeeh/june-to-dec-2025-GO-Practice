package week3

import "fmt"

type Person struct {
	name       string
	age        int
	occupation string
	salary     int
}

func StructInit() string {
	person := Person{
		name:       "John",
		age:        22,
		occupation: "Student",
		salary:     0,
	}

	structure := fmt.Sprintf("Name: %s, Age: %d, Occupation: %s, Salary: %d", person.name, person.age, person.occupation, person.salary)
	return structure

	//using pointer
	//fmt.Println("Person1 name:",person1.name, "\n", person2.age)

}
