package week3

import (
	"fmt"
	_ "fmt"
)


func ForLoops() string {
	var result string
	var count int
	for i := 1; i < 5; i++ {
		result += fmt.Sprintf("For Loop %d\n", i)
		count++
	}
	result += fmt.Sprintf("Total iterations: %d\n", count)
	return result
}