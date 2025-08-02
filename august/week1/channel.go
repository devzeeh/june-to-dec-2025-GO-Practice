package week1 // Channel

import (
	"fmt"
	"time"
)

func Channel() {
	var c = make(chan int, 5) // create a channel of type string
	go process(c)
	for i := range c {
		fmt.Println(i)              // receive from the channel and print the value
		time.Sleep(1 * time.Second) // wait for a second to ensure all messages are printed
	}

}

func process(c chan int) {
	defer close(c) // close the channel when done
	for i := 0; i < 5; i++ {
		c <- i // send value to the channel
	}
	fmt.Println("Exiting process")
}
