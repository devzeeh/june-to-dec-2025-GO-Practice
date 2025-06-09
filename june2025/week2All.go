package main

import (
	"fmt"
	"GoPractice/june2025/week2"
)

// automatically executed when a package is initialized
func init() {
	// This is the init function for week 2 of June 2025
	fmt.Println("Welcome to June 2025 Week 2!")
}

func main() {
	// This is the main function for week 2 of June 2025
	fmt.Println("Result for june week 2 day1: ", week2.Day1())
	fmt.Println("Result for june week 2 day2: ", week2.Day2())
}