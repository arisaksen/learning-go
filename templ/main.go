package main

import (
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

func main() {
	component := hello("world")

	http.Handle("/", templ.Handler(component))

	fmt.Println("http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
