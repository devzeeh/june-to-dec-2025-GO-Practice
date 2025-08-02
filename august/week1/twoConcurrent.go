package week1 // concurrent two tasks example

import (
	"fmt"
	"sync"
	"time"
)

// Task that returns a message through a channel
func task(name string, delay time.Duration, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	fmt.Println(name, "started")
	time.Sleep(delay)
	result := name + " finished"
	ch <- result // Send result to channel
}

func TaskConcurrent() {
	var wg sync.WaitGroup
	ch := make(chan string, 2) // Buffered channel to hold 2 results

	wg.Add(2)

	go task("Task One", 2*time.Second, &wg, ch)
	go task("Task Two", 1*time.Second, &wg, ch)

	wg.Wait()
	close(ch) // Close the channel once all sends are done

	// Receive results from channel
	for msg := range ch {
		fmt.Println("Channel received:", msg)
	}

	fmt.Println("All tasks completed.")
}
