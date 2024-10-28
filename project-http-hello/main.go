package main

import (
	"fmt"
	"log"
	"net/http"
)

var calls = []string{}
var stats = map[string]int{}

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", greet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	calls = append(calls, name)
	count, ok := stats[name]
	if !ok {
		stats[name] = 1
	} else {
		stats[name] = count + 1
	}
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)

	fmt.Fprint(w, "Hello, ", name)

}
