package main

import (
	"fmt"
)

func main() {
	// buffered channel
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	// messages <- "channel"   // Writing when the buffer is full will still block the main goroutine same as an unbuffered channel

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	notOkChannelUse()
}

func notOkChannelUse() {
	// unbuffered channel
	numbers := make(chan int)

	numbers <- 1
	numbers <- 2

	fmt.Println(<-numbers)
	fmt.Println(<-numbers)
}
