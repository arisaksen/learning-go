package main

import (
	"fmt"
	"testing"
)

func TestDotDotDot(t *testing.T) {

	addUsers("Alice", "Bob", "Foo")

	numbers := []int{1, 2, 3, 4, 5}
	test := removeFromSliceReturnUnordered(numbers, 2)
	fmt.Println(test)

	testOrdered := removeFromSliceReturnOrdered(numbers, 2)
	fmt.Println(testOrdered)

	// another example - clean slice method
	newNumbers := MySlice{5, 6, 7, 8, 9, 10}
	newNumbers.Remove(2)
	fmt.Println(newNumbers)
}
