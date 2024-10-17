package main

import (
	"fmt"
	"unsafe"
)

type Header struct {
	id      uint8  // 1 byte
	counter uint16 // 2 bytes
	pulse   uint32 // 4 bytes
	payload uint64 // 8 bytes
	_       uint64 // 8 bytes
}

func ReadSizeOfStruct(header Header) {
	totalSize := unsafe.Sizeof(header)
	fmt.Printf("The total size of the struct in bytes is: %d\n", totalSize)
}
