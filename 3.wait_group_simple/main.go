package main

import (
	"fmt"
	"sync"
	"time"
)

// Preferred way with anonymousFunction for more reusable functions compared to 'worker2'.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	workers := 10

	for i := 1; i <= workers; i++ {

		// You should always call wg.Add() before you launch the goroutine that will call wg.Done().
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("NEW WaitGroup")
	var wg2 sync.WaitGroup
	for i := 1; i <= workers; i++ {
		wg2.Add(1)
		go worker2(i, &wg2)
	}

	wg2.Wait()
}

// WaitGroup inside function
func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
