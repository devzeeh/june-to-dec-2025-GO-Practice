package week3

import _"fmt"

func ifElse() string{
	number := 911938341
	var result string

	if number %2 == 0 {
		result = "even"
	} else {
		result = "odd"
	}
	return result
}

func Week3Result() string {
	return ifElse()
}
