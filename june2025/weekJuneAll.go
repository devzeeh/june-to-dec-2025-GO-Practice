package main

import (
	"fmt"
	"time"

	// Importing the week2 package to access its functions "moduleName/folderPath/packageName"
	"GoPractice/june2025/week2"
	"GoPractice/june2025/week3"
)

// automatically executed when a package is initialized
func init() {
	// This is the init function for week 2 of June 2025
	fmt.Println("Welcome to My Go Programming Basic Guide")
}

func main() {
	// This is the main function for week 2 of June 2025 "packageName.functionName()"
	fmt.Println("For week to of July All About Go Installations")

	time.Sleep(3000 * time.Millisecond) // Sleep for 3 seconds to simulate a delay
	fmt.Println("\nFor week 2 of july")
	time.Sleep(2000) 
	fmt.Println("Result for june week 2 day1: ", week2.Day1())
	fmt.Println("Result for june week 2 day2: ", week2.Day2())
	fmt.Println("Result for june week 2 day3: ", week2.Day3toDay5())
	
	time.Sleep(3000 * time.Millisecond) // Sleep for 3 seconds to simulate a delay
	fmt.Println("\nFor week 3 of july")
	time.Sleep(2000) 
	fmt.Println("Result for june week 2 day1: ", week3.Week3Result())
	fmt.Println("Result for Switch Statement: ", week3.Switch())
	fmt.Println("Result for For Loops: ", week3.ForLoops())
	fmt.Println("Result for Fibonacci: ", week3.Fibonacci())
}