package main

import (
	"runtime"
)

func getUsedMem() uint64 {
	// A MemStats records statistics about the memory allocator.
	var m runtime.MemStats
	//runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
