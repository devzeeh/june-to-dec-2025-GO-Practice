package week3

import (
	"fmt"
	_ "fmt"
)

func ForLoops() string {
	n := 10           // Number of iterations
	var result string // Initialize an empty string to store the result
	for i := 1; i <= n; i++ {
		result += fmt.Sprintf("%d ", i)
	}
	return result
}
