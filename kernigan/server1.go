package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("running server on localhost:8000")
	http.HandleFunc("/", handler) // each request requests call handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handling... %q\n", r.URL.Path)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
