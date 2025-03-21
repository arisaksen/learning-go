package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
)

type mapServer struct {
	locations []Location
	center    Location
	zoom      float64
}

func newMapServer() *mapServer {
	s := new(mapServer)
	s.locations = []Location{
		{"Point A", 69.65 + 0.01, 18.94 - 0.07},
		{"Point B", 69.65, 18.94 - 0.07},
		{"Point C", 69.65 - 0.01, 18.94 - 0.07},
	}
	s.center = Location{
		Name: "start",
		Lat:  69.65,
		Lng:  18.94,
	}
	s.zoom = 13

	return s
}

func (ms *mapServer) ListenAndServe() error {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Print(err)
		return err
	}
	defer listener.Close()

	log.Printf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)
	if err = http.Serve(listener, nil); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (ms *mapServer) addLocationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addLocationHandler called")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	lat, err1 := strconv.ParseFloat(r.FormValue("lat"), 64)
	lng, err2 := strconv.ParseFloat(r.FormValue("lng"), 64)
	newCenterLat, err3 := strconv.ParseFloat(r.FormValue("newCenterLat"), 64)
	newCenterLng, err4 := strconv.ParseFloat(r.FormValue("newCenterLng"), 64)
	zoom, err4 := strconv.ParseFloat(r.FormValue("zoom"), 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		http.Error(w, "Invalid latitude or longitude", http.StatusBadRequest)
		return
	}

	newLocation := Location{Name: name, Lat: lat, Lng: lng}
	log.Println("newLocation:", newLocation)
	ms.locations = append(ms.locations, newLocation)
	ms.center = Location{"Center", newCenterLat, newCenterLng}

	w.Header().Set("Content-Type", "text/html")
	if err := MyDocument(ms.center, ms.locations, zoom).Render(context.Background(), w); err != nil {
		log.Fatal(err)
	}
}
