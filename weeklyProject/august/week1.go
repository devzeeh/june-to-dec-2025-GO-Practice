package august

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simulateDownload simulates fetching a URL with random delay
func simulateDownload(url string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	// Simulate random download time
	delay := time.Duration(rand.Intn(2000)+500) * time.Millisecond
	time.Sleep(delay)

	// Send result to channel
	ch <- fmt.Sprintf("Finished downloading %s in %v", url, delay)
}

func AugustWeek1Result() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	urls := []string{
		"https://example.com/file1",
		"https://example.com/file2",
		"https://example.com/file3",
		"https://example.com/file4",
		"https://example.com/file5",
	}

	var wg sync.WaitGroup
	ch := make(chan string)

	// Launch goroutines
	for _, url := range urls {
		wg.Add(1)
		go simulateDownload(url, &wg, ch)
	}

	// Collect results concurrently
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Print results
	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("All downloads completed.")
}
