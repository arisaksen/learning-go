package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	done := make(chan interface{})
	defer close(done)

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	// these will run for as long as the done channel is open
	go func() {
		for {
			select {
			// breaks when done is closed. (when main stops)
			case <-done:
				return
			case cows <- "moo":
			}
		}
	}()

	// these will run for as long as the done channel is open
	go func() {
		for {
			select {
			// breaks when done is closed. (when main stops)
			case <-done:
				return
			case pigs <- "oink":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows, &wg)
	wg.Add(1)
	go consumePigs(done, pigs, &wg)

	wg.Wait()
}

func consumePigs(done chan interface{}, pigs chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for pig := range orDone(done, pigs) {

		// do some complex logic
		fmt.Println(pig)
	}
}

func consumeCows(done chan interface{}, cows chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for cow := range orDone(done, cows) {

		// do some complex logic
		fmt.Println(cow)
	}
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return relayStream
}

// without the orDone() function
//func consumePigs(done chan interface{}, pigs chan interface{}, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		select {
//		case <-done:
//			return
//		case pig, ok := <-pigs: // if pigs channel closed ok=false
//			if !ok {
//				fmt.Println("channel closed")
//				return
//			}
//
//			// do some complex logic
//			fmt.Println(pig)
//		}
//	}
//
//}
