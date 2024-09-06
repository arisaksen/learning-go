package main

import "fmt"

// channels are for passing data between goroutines
// two types of channels: unbuffered, buffered
func main() {
	// unbuffered channel is default behavior
	messages := make(chan string)

	// must use go routine to write to 'unbuffered channel'
	go func() { messages <- "ping1" }()

	// Read single value from channel
	msg := <-messages
	fmt.Println(msg)

	bufferedChannel()
}

func bufferedChannel() {

	messages := make(chan string, 10)

	// Ok to not use goroutine to write to buffered channel.
	messages <- "ping2"

	// Read single value from channel
	msg := <-messages
	fmt.Println(msg)
}
