package main

import "fmt"

func okChannelUse1() {
	defer fmt.Println("Finished okChannelUse")

	var channel = make(chan int)
	go func(channel chan int) {
		channel <- 42
		channel <- 43
		channel <- 44
	}(channel)
	fmt.Println(<-channel)
}

func okChannelUse2() {
	defer fmt.Println("Finished okChannelUse")

	var channel = make(chan int)

	// here we create another goroutine to receive the value before sending it
	go func() { fmt.Println(<-channel) }()

	// so now this is ok
	channel <- 42
}

// https://coffeebytes.dev/en/go-channels-understanding-the-goroutines-deadlocks/
func main() {

	// In go, operations that send or receive channel values are blocking inside their own goroutine (remember that the main function is a goroutine), i.e., they keep code execution waiting:

	// fatal error: all goroutines are asleep - deadlock!
	// This error occurs when:
	//
	// A channel sends information, but not channel is there to receive it.
	//	There is a channel that receives information, but not channel that sends it.
	//	When we are not inside a goroutine other than the one from the main function.

	okChannelUse1()
	okChannelUse2()

	// Blocking or deadlock due to lack of sender
	ch1 := make(chan string)
	fmt.Println(<-ch1)

	// Blocking or deadlock for lack of recipient
	ch2 := make(chan string)
	ch2 <- "text"

}
