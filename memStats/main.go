package main

import (
	"fmt"
	"runtime"
)

func printMemStats(title string) {
	fmt.Println(title)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v MB\n\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1000 / 1000 // SI unit decimal form
	//return b / 1024 / 1024 // binary form
}

func main() {

	printMemStats("Before:")
	s := make([]int, 100_000_000)
	for i := range s {
		s[i] = i
	}

	printMemStats("After:")
}
