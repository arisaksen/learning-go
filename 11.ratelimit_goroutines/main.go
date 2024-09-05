package main

import (
	"fmt"
	"time"
)

// See YouTube video:
// https://www.youtube.com/watch?v=tSjnf6l8cq8&list=WL&index=2&t=288s
func main() {
	manager := CreateManager(1 * time.Second)

	for event := range manager.Stream() {
		fmt.Println(event)
	}

}
