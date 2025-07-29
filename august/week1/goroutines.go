package week1 // August Goroutine Week

import (
	"fmt"
	"time"
)

// routine prints numbers 1 to 5 with a delay
func routine() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond) // Delay to simulate work
		fmt.Print(i, " ")
	}
}

// anonymous launches a goroutine that prints numbers
func anonymous() {
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Print(i, " ")
			// Optional: Add a delay if needed
			// time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Anonymous goroutine finished")
	}()
	time.Sleep(1 * time.Second) // Allow time for goroutine to finish
}

func Goroutine() {
	fmt.Println("Starting routine in a goroutine")
	go routine() // Run in a separate goroutine

	fmt.Println("Running routine in main")
	routine() // Run in main thread

	fmt.Println("\nGoroutine function finished")

	fmt.Println("Starting anonymous goroutine")
	anonymous()// Launch an anonymous function as a goroutine
}
