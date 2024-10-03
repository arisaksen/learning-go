package main

import "fmt"

type outPutstrategy interface {
	createOutput(s string) string
}
type stringStrategy struct{}

func (rcv stringStrategy) createOutput(s string) string {
	return s
}

type byteStrategy struct{}

func (rcv byteStrategy) createOutput(s string) string {
	return fmt.Sprintf("%v", []byte(s))
}

type hexStrategy struct{}

func (rcv hexStrategy) createOutput(s string) string {
	return fmt.Sprintf("%x", []byte(s))
}
