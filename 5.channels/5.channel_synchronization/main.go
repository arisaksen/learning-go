package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
	fmt.Println()

	done <- true
}

// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
func main() {

	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
	// if we removed the '<-done' line from this program, the program would exit before the worker even started.

}
