package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

func Worker(results chan<- []byte, errorsChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := DummyGetResponse()
	if err != nil {
		errorsChan <- err
	}

	results <- response
}

func DummyGetResponse() ([]byte, error) {
	response := []byte{byte(0b11111111)} // 1 byte - 8 bit response. Decimal: 255, Hex: 0xFF, Bin: 0b 1111.1111
	return response, nil
}

func main() {
	const numberOfResults = 10
	buffer := new(bytes.Buffer)

	results := make(chan []byte, numberOfResults)
	errorsChan := make(chan error, numberOfResults)

	var wg sync.WaitGroup
	for i := 0; i < numberOfResults; i++ {
		wg.Add(1)
		go Worker(results, errorsChan, &wg)
	}

	wg.Wait()
	close(results)
	close(errorsChan)
	for result := range results {
		buffer.Write(result)
	}
	for err := range errorsChan {
		if err != nil {
			log.Fatalf("Error %v", err)
		}
	}
	fmt.Printf("Length of buffer. %d", buffer.Len())
}
