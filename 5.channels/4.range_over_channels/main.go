package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	totalOrders := 5
	orders := make(chan string, totalOrders)

	orders <- "Order1"
	orders <- "Order2"
	orders <- "Order3"
	orders <- "Order4"
	orders <- "Order5"
	close(orders) //  NB! to range over the channel, it needs to be closed

	// process
	for order := range orders {
		fmt.Println("Handling: " + order)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All orders processed.")
	fmt.Println()

	withAnonymousFunction()
}

func withAnonymousFunction() {
	totalOrders := 7
	orders := make(chan string, totalOrders)

	go func() {
		defer close(orders) //  NB! to range over the channel, it needs to be closed

		for i := 1; i <= totalOrders; i++ {
			order := "Order" + strconv.Itoa(i)
			orders <- order
			fmt.Println("Placed: " + order)
		}
	}()

	for order := range orders {
		fmt.Println("Handling: " + order)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All orders processed inside function.")
}
