package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	hello := strings.NewReader("hello")
	space := strings.NewReader(" ")
	world := strings.NewReader("world!")

	reader := io.MultiReader(hello, space, world)
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}
}
