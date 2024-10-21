package sizeOfStructs

import (
	"fmt"
	"testing"
	"unsafe"
)

type Struct1 struct {
	a bool   // 1 byte
	b string // 16 byte
	c bool   // 1 byte
}

type Struct2 struct {
	a bool   // 1 byte
	c bool   // shares padding with 'a'
	b string // 16 byte
}

func TestSizeOfStructsWithBool(t *testing.T) {
	header := Header{}
	ReadSizeOfStruct(header)

	struct1 := Struct1{}
	struct1Size := unsafe.Sizeof(struct1)
	fmt.Printf("The total size of the struct in bytes is: %d\n", struct1Size)
	if struct1Size != 32 {
		t.Error("Not expected size")
	}

	struct2 := Struct2{}
	struct2Size := unsafe.Sizeof(struct2)
	fmt.Printf("The total size of the struct in bytes is: %d\n", struct2Size)
	if struct2Size != 24 {
		t.Error("Not expected size")
	}
}
