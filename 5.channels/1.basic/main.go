package main

import "fmt"

// channels are for passing data between goroutines
// two types of channels: unbuffered, buffered
func main() {
	// unbuffered channel is default behavior
	messages := make(chan string)

	// write to channel
	go func() { messages <- "ping1" }()

	// Read single value from channel
	msg := <-messages
	fmt.Println(msg)

}
