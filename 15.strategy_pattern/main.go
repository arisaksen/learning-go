package main

import (
	"flag"
	"fmt"
	"log"
)

type printer struct {
	strategy outPutstrategy
}

func (rcv *printer) setStrategy(s outPutstrategy) {
	rcv.strategy = s
}

func (rcv *printer) print(input string) {
	output := rcv.strategy.createOutput(input)
	fmt.Println(output)
}

// go run . -i "hello" -s string
// go run . -i "hello" -s byte

func main() {

	// https://pkg.go.dev/flag
	input := flag.String("i", "", "the input")
	strat := flag.String("s", "", "the input")

	flag.Parse()

	p := printer{}

	switch {
	case *strat == "string":
		p.setStrategy(stringStrategy{})
	case *strat == "byte":
		p.setStrategy(byteStrategy{})
	case *strat == "hex":
		p.setStrategy(hexStrategy{})
	default:
		log.Fatal("no strategy specified")
	}

	p.print(*input)

}
