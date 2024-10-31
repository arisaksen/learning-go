package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()

	go func() {
		defer pipeWriter.Close() // will cause EOF error

		// write each word to the list
		inputText := []string{"hello", " ", "world", "!"}
		for _, input := range inputText {
			_, err := fmt.Fprint(pipeWriter, input)
			if err != nil {
				panic(err)
			}
		}

	}()

	buffer := new(bytes.Buffer)
	// Because Copy is defined to read from src until EOF, it does
	// not treat an EOF from Read as an error to be reported.
	_, err := io.Copy(buffer, pipeReader)
	if err != nil {
		panic(err)
	}

	for _, text := range buffer.String() {
		fmt.Print(string(text))
	}
}
