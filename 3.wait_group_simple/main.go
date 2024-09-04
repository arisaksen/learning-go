package main

import (
	"fmt"
	"sync"
	"time"
)

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
}

// WaitGroup inside function
//func worker(id int, wg *sync.WaitGroup) {
//	wg.Add(1)
//	defer wg.Done()
//
//	fmt.Printf("Worker %d starting\n", id)
//	time.Sleep(time.Second)
//	fmt.Printf("Worker %d done\n", id)
//}
//
//func main() {
//	var wg sync.WaitGroup
//	workers := 10
//
//	for i := 1; i <= workers; i++ {
//		go worker(i, &wg)
//	}
//
//	wg.Wait()
//}
