package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

func main() {
	const numberOfResults = 10
	buffer := new(bytes.Buffer)

	// Create channels to store results and errors. This is like a que to store values
	results := make(chan []byte, numberOfResults)
	errorsChan := make(chan error, numberOfResults)

	// A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup

	for i := 0; i < numberOfResults; i++ {

		// +1 to waitGroup to let it know how many Goroutines to wait for
		wg.Add(1)

		// 'Go' starts a goroutine
		go Worker(results, errorsChan, &wg)
	}

	// Wait here for finish and read responses from channels
	wg.Wait()

	// Close channels before reading data
	close(results)
	close(errorsChan)

	// Write all data to buffer
	for result := range results {
		buffer.Write(result)
	}

	// Check for errors
	for err := range errorsChan {
		if err != nil {
			log.Fatalf("Error %v", err)
		}
	}

	fmt.Printf("Length of buffer. %d", buffer.Len())
}
