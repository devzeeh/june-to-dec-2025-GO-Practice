package week4

import "fmt"

// This example shows how panic halts function execution, and how recover can
// catch that panic to prevent the program from crashing completely.

// zeeh calls panicking and shows control flow around a panic-recovered call.
func zeeh() {
	fmt.Println("zeeh called")
	panicking()
	fmt.Println("zeeh finished") // This will still run after recover
}

// panicking simulates a runtime panic and recovers from it using defer.
func panicking() {
	fmt.Println("panic called")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover") // This handles the panic
		}
	}()
	panic("OH NO")                // Triggers the panic
	fmt.Println("panic finished") // Will NOT be executed
}

func PanicRecover() string {
	// Last to execute

	// Will execute first
	fmt.Println("start!") // First to execute
	zeeh()
	fmt.Println("end main") // Will run after recover in panicking()
	return ""
}
