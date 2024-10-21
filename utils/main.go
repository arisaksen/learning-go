package main

import (
	"fmt"
)

func main() {

	// getMemory
	var strs []string
	for i := 0; i < 50000000; i++ {
		strs = append(strs, "example")
	}
	beforeMemory := getUsedMem()
	fmt.Printf("b: %d bytes\n", beforeMemory)
	fmt.Printf("b: %f MB\n", float64(beforeMemory)/1000000)

}
