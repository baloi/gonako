package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index called...")
	fmt.Fprintf(w, "Welcome %s!", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", about)
	fmt.Println("Handlers added")
	http.ListenAndServe(":8080", nil)
}

type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about called")
	m := Message{ "Welcome to Baloi's API, buld 0.2.2.3, 11/22/2017" }
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}