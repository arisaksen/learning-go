package main

import (
	"embed"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

func main() {
	server := newMapServer()
	doc := MyDocument(server.center, server.locations, server.zoom)

	http.Handle("/", templ.Handler(doc))
	http.Handle("/add-location", http.HandlerFunc(server.addLocationHandler))
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
