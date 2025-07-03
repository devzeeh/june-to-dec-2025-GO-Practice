package week1

import (
	_ "fmt"
)

func array() []string {
	array := [5]string{"Freya", "Athena", "Angel", "Mika", "Kazumi"} 
	result := array[:5]
	return result
}

func Array() []string{
	return array()
}
