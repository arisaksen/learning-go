package main

import (
	"fmt"
	"time"
)

type ImageTask struct {
	ImageUrl string
}

func (i *ImageTask) Process() {
	// Simulate a time consuming process
	time.Sleep(3 * time.Second)

	fmt.Printf("Processing image %s\n", i.ImageUrl)
}
