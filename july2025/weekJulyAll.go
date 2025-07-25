package main

import (
	"fmt"
	"time"

	"GoPractice/july2025/week1"
	"GoPractice/july2025/week2"
	"GoPractice/july2025/week3"
	"GoPractice/july2025/week4"
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
	original, add, update, deletes := week2.AddUpdateDelMaps()
	fmt.Println("Result for maps: -- Add 'Color' key, Update 'Year:Value', Delete 'Color' key\n", original, add, update, deletes)
	fruit, res := week2.LoopOverMap()
	fmt.Println("Result for loop over maps:", res, fruit, res)
	fmt.Println("Result for check key: ", week2.KeyCheck())
	originalPhonebook, add, exist, newPhonebook := week2.Maphonebook()
	fmt.Println("Result for map phonebook: ", originalPhonebook, add, exist, newPhonebook)

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nFor week 3 of july")
	structure := week3.StructInit()
	fmt.Println("Result for Struct:", structure)
	laptopStorage, laptopRam, laptopCpuManu := week3.StructFieldMethod()
	fmt.Print("Result for Struct: ", laptopStorage, laptopRam, laptopCpuManu)
	nested := week3.NestedStruct()
	fmt.Print("\nResult for Nested Struct: ", nested)
	allStruct, oldSalary, updatedSalary := week3.StructPerson()
	fmt.Print("\nResult for Last Struct: ", allStruct, oldSalary, updatedSalary)

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nFor week 4 of july")
	ResultMes, ErrorMes := week4.Error()
	fmt.Print("Result for Error: ", ResultMes, ErrorMes)
	functionError := week4.FunctionError()
	fmt.Print("Result for Function Error: ", functionError)
	fmt.Println("Result for Panic Recover: ")
	week4.PanicRecover()
	validError, resTurn := week4.ValidateError()
	fmt.Print("Result for validate error: ", resTurn, validError)
}
