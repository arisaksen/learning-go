package main

import "fmt"

const size = 100

func DynamicArray() (int, int) {
	var s []int
	for i := 0; i < size; i++ {
		s = append(s, i)
	}

	// cap() returns the possible capacity of the array
	return len(s), cap(s)
}

func StaticArray() (int, int) {
	s := [size]int{}

	for i := 0; i < size; i++ {
		s[i] = i
	}

	return len(s), cap(s)
}

func main() {
	fmt.Println("Dynamic: ")
	fmt.Println(DynamicArray())

	fmt.Println("Static: ")
	fmt.Println(StaticArray())
}
