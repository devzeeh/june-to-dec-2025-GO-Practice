package week1

import (
	_ "fmt"
)

func array(number int) string {
	array := [5]string{"Freya", "Athena", "Angel", "Mika", "Kazumi"} 
	result := array[number]
	return result
}

func Array(number int) string{
	return array(number)
}
