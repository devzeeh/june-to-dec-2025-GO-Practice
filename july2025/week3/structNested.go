package week3

import "fmt"

// Define struct 'student' to represent a student's basic information
type student struct {
	name    string
	program string
	year    int
}

// Define struct 'teacher' that includes a nested 'student' struct
type teacher struct {
	name    string
	subject string
	exp     int
	details student // Nested student struct representing related student info
}

// NestedStruct demonstrates the use of a nested struct by initializing
// a 'teacher' struct that includes a 'student' struct as one of its fields.
func NestedStruct() string {
	// Initialize a teacher instance with nested student data
	result := teacher{
		name:    "Carla",
		subject: "GOLang",
		exp:     5,
		details: student{
			name:    "John",
			program: "BSCpE",
			year:    2,
		},
	}

	nested := fmt.Sprint(result, " Name of teacher: ", result.name) // Output {Carla GOLang 5 {John BSCpE 2}} Name of teacher: Carla
	return nested
}
