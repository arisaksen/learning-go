package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	timeout := 7 * time.Second
	counterMax := 6
	//counterMax := 8

	ctx, cancel := context.WithCancel(context.Background())
	var counter int
	go func() {
		for {
			if counter > counterMax {
				cancel()
			}
			counter++
			time.Sleep(1 * time.Second)
			log.Println("counter", counter)
		}
	}()

	log.Println("waiting for done or timeout")
	select {
	case <-time.After(timeout):
		log.Println("timeout")
	case <-ctx.Done():
		log.Println("done")
	}

}
