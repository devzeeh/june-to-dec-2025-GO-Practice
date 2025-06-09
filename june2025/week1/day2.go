package main 

import (
	"fmt"
	"time"
)

func day3() (int, string, string) {
	day := 3
	// Get the current date
	//loc, _:= time.LoadLocation("Asia/Manila")
	time := time.Now()//.In(loc)
	date := time.Format("2006-01-02") // year-month-day format
	timeStr := time.Format("03:04:05 PM")// 12-hour format (12hrs, minutes, seconds, AM/PM)
	return day, date, timeStr
}

func main(){
	fmt.Print("Hello, Go! \n")
	fmt.Print("Hello, Gopher!")
	day, date, timeStr := day3()
	fmt.Printf("Day: %d, Date: %s, Time: %s", day, date, timeStr)
}

