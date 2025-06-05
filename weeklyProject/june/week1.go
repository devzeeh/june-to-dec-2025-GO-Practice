package main

import (
	"fmt"
	"strconv"
	"time"
)

var name = []string{
	"John",
	"Frey",
	"Angela",
}

// Declare a global variable for the week number. 	Will not function if dont have variable declaration
var week int = 1

func day() (int, string, string) {
	var day = 3
	//Get the current date
	//loc, _:= time.LoadLocation("Asia/Taipei") // Load the location for Manila timezone
	time := time.Now()                    //.In(loc)            // Load the location for Manila timezone
	date := time.Format("2006-01-02")     // year-month-day format
	timeStr := time.Format("03:04:05 PM") // 12-hour format (12hrs, minutes, seconds, AM/PM)
	return day, date, timeStr
}

func main() {
	fmt.Println("Welcome to Go programming Language!")                                       // FormatInt converts int to string
	fmt.Println("Week: " + strconv.FormatInt(int64(week), 10) + ", using 'strconv' pacakge") // Get the current day of the month and convert to base 10
	fmt.Printf("week: %d, using 'fmt' package \n", week)                                     // Print the week number using fmt package
	fmt.Println("Name: " + "Hello", name[0] + ", " + "Hello", name[1] + ", " + "Hello", name[2])                                                               // Print the names from the slice
	// Call the day function to get the current day, date, and time
	day, date, timeStr := day()
	fmt.Printf("Day: %d, Date: %s, Time: %s", day, date, timeStr)
}
