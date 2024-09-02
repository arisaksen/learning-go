package main

import (
	"fmt"
	"runtime"
)

func getUsedMem() uint64 {
	// A MemStats records statistics about the memory allocator.
	var m runtime.MemStats
	//runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
func main() {
	var strs []string
	for i := 0; i < 50000000; i++ {
		strs = append(strs, "example")
	}

	beforeMemory := getUsedMem()
	fmt.Printf("b: %d bytes\n", beforeMemory)
	fmt.Printf("b: %f MB\n", float64(beforeMemory)/1000000)

}
