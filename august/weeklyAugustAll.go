package main

import (
	"fmt"
	"time"

	"GoPractice/august/week1"
	"GoPractice/august/week2"
	"GoPractice/august/week3"
)

func init() {
	fmt.Println("Welcome to August Go Practice")
}

func main() {
	fmt.Println("\nFor Week 1 of August")
	fmt.Print("Result of Goroutine execution:")
	week1.Goroutine() // Call the Goroutine function from week1 package
	fmt.Print("\nResult of WaitGroup execution:")
	week1.SyncWg() // Call the SyncWg function from week1 package
	fmt.Print("\nResult of Channel execution:")
	week1.Channel() // Call the Channel function from week1 package
	fmt.Println("\nResult of Two Task Concurrent: ")
	week1.TaskConcurrent() // Call the TaskConcurrent function from week1 package
	fmt.Println("\nResult of Open, Read and Write: ")

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nFor Week 2 of August")
	week2.OpenReadWrite() // Call the OpenReadWrite function from week2 package
	fmt.Println("\nResult of Append Data to File: ")
	week2.AppendFileText() // Call the AppendFileText function from week2 package
	fmt.Println("\nResult of Copy File: ")
	week2.CopyFile() // Call the CopyFile function from week2 package
	fmt.Println("\nResult of net/http server: ")

	time.Sleep(3 * time.Second)
	fmt.Println("\nFor Week 3 of August")
	fmt.Println("For this we can use Curl command to test the server")
	week3.NetHttp() // Call the NetHttp function from week3 package

}
