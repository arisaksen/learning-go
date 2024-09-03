package main

import (
	"fmt"
	"time"
)

type EmailTask struct {
	Email string
}

func (e *EmailTask) Process() {
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)

	fmt.Printf("Processing email. Sending to %s\n", e.Email)
}
