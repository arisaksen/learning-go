package main

import "fmt"

// this was fixed in go 1.22. But if you switch to go 1.21 the rusult will be:
// three three three
func main() {
	orig := []string{"one", "two", "three"}

	var ptr []*string
	for _, item := range orig {
		ptr = append(ptr, &item)
	}

	for _, p := range ptr {
		fmt.Println(*p)
	}

}
