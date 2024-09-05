package main

import (
	"fmt"
	"time"
)

func main() {
	manager := CreateManager(1 * time.Second)

	for event := range manager.Stream() {
		fmt.Println(event)
	}

}
