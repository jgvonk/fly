package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jgvonk/fly/flying"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unhandled route: \"%s\"", r.URL.Path[1:])
	fmt.Printf("Unhandled route: \"%s\"\n", r.URL.Path[1:])
}

func takeoff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "takeoff")
	fmt.Println("takeoff")
	fl := flying.NewMyDrone()
	fl.Start()
}

func topos(w http.ResponseWriter, r *http.Request) {
	position := r.URL.Path[7:]
	fmt.Fprintf(w, "to position %s", position)
	fmt.Printf("to position %s\n", position)
}

func land(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "land")
	fmt.Println("land")
}

func main() {
	http.HandleFunc("/takeoff/", takeoff)
	http.HandleFunc("/topos/", topos)
	http.HandleFunc("/land/", land)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
