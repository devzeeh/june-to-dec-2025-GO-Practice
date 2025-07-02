package main

import (
	"fmt"

	"GoPractice/july2025/week1"
)
func init() {
	fmt.Println("Welcome to July Go Practice")
}

func main() {
	fmt.Println("Result for array: ", week1.Array(4)) // call the 4th element of the array maximum index is 4 ["Freya", "Athena", "Angel", "Mika", "Kazumi"]
	fmt.Println("Result for slices: ", week1.Slices())
}