package main

import (
	"fmt"
)

func work(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	work("direct")

	go work("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("goroutine2")

	//time.Sleep(time.Second) // If we add the wait here. The goroutines will not have time to finish
	fmt.Println("done")
}
