package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()

	_, err := fmt.Fprintln(pipeWriter, "hello")
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, pipeReader)
	if err != nil {
		panic(err)
	}
}
