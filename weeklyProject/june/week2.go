package june

import "strconv"

// Declare a global variable for the temperature in Celsius
var temperature float64 = 25.00

// Function to convert Celsius to Fahrenheit F = C × 9/5 + 32
func celsiusToFahrenheit() float64 {
	//celsius := 515.35
	fahrenheit := (temperature * 9/5) + 32 // Convert Celsius to Fahrenheit and Kelvin
	//kelvin := tempFahrenheit + 273.15 // Convert Celsius to Kelvin K = C + 273.15
	//celsius := (tempCelsius * 9/5) + 32 // Convert Fahrenheit to Celsius (25°C × 9/5) + 32 = 77°F
	return fahrenheit
}

// Convert f
func celsiusToKelvin() float64{
	kelvin := temperature + 273.15 // Convert Celsius to Kelvin K = C + 273.15
	return kelvin
}

func fahrenheitToCelsius() float64{
	celsius := (temperature - 32) * 5/9 // Convert Fahrenheit to Celsius (25°F − 32) × 5/9
	return celsius
}


func Week2Result() string{
	celsius := strconv.FormatFloat(float64(fahrenheitToCelsius()), 'f', 2, 64)
	fahrenheit := strconv.FormatFloat(float64(celsiusToFahrenheit()), 'f', 0, 64)
	kelvin := strconv.FormatFloat(float64(celsiusToKelvin()), 'f', 2, 64)
	return "Temperature Convertion Celsius: " + celsius + "°F, Fahrenheit: " + fahrenheit + "°F, Kelvin: " + kelvin + "K"
	//return convertTemperature(25.0) // Example temperature in Celsius
}