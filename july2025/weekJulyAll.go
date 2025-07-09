package main

import (
	"fmt"
	"time"

	"GoPractice/july2025/week1"
	"GoPractice/july2025/week2"
)
func init() {
	fmt.Println("Welcome to July Go Practice")
}

func main() {
	fmt.Println("\nFor Week 1 of July")
	fmt.Println("Result for array: ", week1.Array(), "--Calling all String in Array") // call the 4th element of the array maximum index is 4 ["Freya", "Athena", "Angel", "Mika", "Kazumi"]
	fmt.Println("Result for slices: ", week1.Slices(), "--Remove 'Mika', 'Kazumi'")
	fmt.Println("Result for appending: ", week1.AppendSlice(), "--Add 'Mika', 'Kazumi'")
	fmt.Println("Result for slices using 'For' & 'Range': ", week1.RangeSlice())
	min, max := week1.MinMaxSlice()
	fmt.Println("Result for slices using 'For' & 'Range': Min =", min, ", Max =", max)

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nFor Week 2 of July")
	original, add, update, deletes:= week2.AddUpdateDelMaps()
	fmt.Println("Result for maps: -- Add 'Color' key, Update 'Year:Value', Delete 'Color' key\n",original,add,update,deletes)
	fruit, res := week2.LoopOverMap()
	fmt.Println("Result for loop over maps:",res,fruit,res)

}