package week3

import "fmt"

var category = "fruit"
var product = "apple"

// function for switchstatement
func switchStatement() string {
	switch category {
	case "fruit":
		switch product {
		case "apple":
			return fmt.Sprintf("Category: %s, Item: Apple", category)
		default:
			return "Unknown Product"
		}
	case "vegetable":
		switch product {
		case "carrot":
			return fmt.Sprintf("Category: %s, Item: carrot", category)
		default:
			return "Unknown Product"
		}
	default:
		return fmt.Sprintf("unknown category: %s", category)
	}
}

func Switch() string {
	return switchStatement()
}
