package june

import "fmt"

var name = []string{
	"John",
	"Frey",
	"Angela",
}

// Call the Week1 function to get the week number and names
func Week1Result() string {
	result := fmt.Sprintf("Name: Hello %s, Hello %s, Hello %s", name[0], name[1], name[2])
	return result
}
