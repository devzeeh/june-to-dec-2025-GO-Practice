package june

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

func week1() string {
	var result string
	
	result += "Welcome to Go programming Language!\n"
	result += "Week: " + strconv.FormatInt(int64(week), 10) + ", using 'strconv' pacakge\n"
	result += fmt.Sprintf("week: %d, using 'fmt' package \n", week)
	result += fmt.Sprintf("Name: Hello %s, Hello %s, Hello %s\n", name[0], name[1], name[2])
	
	// Call the day function to get the current day, date, and time
	day, date, timeStr := day()
	result += fmt.Sprintf("Day: %d, Date: %s, Time: %s", day, date, timeStr)
	
	return result
}

func Week1Result() string {
	// Call the Week1 function to get the week number and names
	return week1()
	
	// Format the result string
	//return fmt.Sprintf("Weeksdfghjk %d Result: %s", week, result)
}
