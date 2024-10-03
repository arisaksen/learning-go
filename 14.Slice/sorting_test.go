package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSorting(t *testing.T) {
	test := Numbers{1, 10, 4, 6, 8, 2, 3}

	sort.Sort(byInc{test})
	fmt.Println(test)

	sort.Sort(byDec{test})
	fmt.Println(test)
}

func TestBonusMethods(t *testing.T) {
	test := Numbers{1, 10, 4, 6, 8, 2, 3}

	fmt.Println(test.Len())

	test.Swap(0, len(test)-1)
	fmt.Println(test)

	if test.Greater(0, len(test)-1) {
		t.Log("OK")
	} else {
		t.Error("Not expected result for Greater method")
	}
}
