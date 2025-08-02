package main

import (
	"fmt"

	"GoPractice/august/week1"
)

func init() {
	fmt.Println("Welcome to August Go Practice")
}

func main() {
	fmt.Println("\nFor Week 1 of August")
	fmt.Print("Result of Goroutine execution:")
	week1.Goroutine() // Call the Goroutine function from week1 package
	fmt.Print("\nResult of WaitGroup execution:")
	week1.SyncWg()
	fmt.Print("\nResult of Channel execution:")
	week1.Channel()
	fmt.Println("\nResult of Two Task Concurrent: ")
	week1.TaskConcurrent()

	//fmt.Println("\nFor Week 2 of August")
	// Add calls for week2 functions here when implemented

	//fmt.Println("\nFor Week 3 of August")
	// Add calls for week3 functions here when implemented

	//fmt.Println("\nFor Week 4 of August")
	// Add calls for week4 functions here when implemented
}
