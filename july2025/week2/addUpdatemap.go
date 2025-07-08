package week2
import (
	"fmt"
)

func AddUpdateDelMaps() (string, string, string, string) {
	// create empty map
	//carDetails := make(map[string]string)
	//carDetails["Brand"] = "Toyota"

	// create maps (original maps)
	carDetails := map[string]string{"Brand": "Toyota", "Model": "Corolla", "Year": "1997"}
	original := fmt.Sprint("Original:", carDetails)

	// Add key:value Color: Grey
	carDetails["Color"] = "Grey"
	add := fmt.Sprintf("\nAdd: %v -- Add 'Color' key", carDetails)

	//update the Year
	carDetails["Year"] = "2020"
	update := fmt.Sprintf("\nUpdate: %v -- Update 'Year:{Value}'", carDetails)

	//delete the color
	delete(carDetails, "Color")
	deletes := fmt.Sprintf("\nDelete: %v -- Delete 'Color' key",carDetails)
	
	return original, add, update, deletes
}
