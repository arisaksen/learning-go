package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add increments the [WaitGroup] counter by one.
	wg.Add(1)
	go func() {

		// Done decrements the [WaitGroup] counter by one.
		defer wg.Done()

		time.Sleep(3 * time.Second)
	}()

	// Wait blocks until the [WaitGroup] counter is zero.
	wg.Wait()

	fmt.Println("Done")
}
