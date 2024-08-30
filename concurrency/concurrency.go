package main

import (
	"sync"
)

func Worker(results chan<- []byte, errorsChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate call to get data
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
