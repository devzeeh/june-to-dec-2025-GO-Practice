package main

import (
	"fmt"
	// Importing the week2 package to access its functions "moduleName/folderPath/packageName"
	"GoPractice/weeklyProject/august"
	"GoPractice/weeklyProject/july"
	"GoPractice/weeklyProject/june"
)

// automatically executed when a package is initialized
func init() {
	// This is the init function for week 2 of June 2025
	fmt.Println("Welcome to June - December 2025 Weekly Mini Project!")
}

func main() {
	// This is the main function "packageName.functionName()"
	fmt.Println("Result for june week 1: ", june.Week1Result())
	fmt.Println("Result for june week 2: ", june.Week2Result())
	fmt.Println("Result for june week 3: ", june.Week3Result())
	fmt.Println("Result for june week 4: ", june.Week4Result())

	// For the month of July
	fmt.Println("\n Weekly project for month of July")
	july.Week1Result1()
	a, initialInventory, initial, addInv, addStock, updateStock, totalQty := july.Inventory()
	fmt.Println(a, initialInventory, initial, addInv, addStock, updateStock, totalQty)
	fmt.Println("Result for Week3 Struct", july.Week3Result())
	fmt.Println("Result for Week3 Struct", july.Week4Result())

	// For the month of August
	fmt.Println("\n Weekly project for month of August")
	fmt.Println("Result for August Week 1:")
	august.AugustWeek1Result()

}
