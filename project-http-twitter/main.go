package main

import (
	"log"
	"net/http"
	"twitter/server"
)

func main() {
	s := server.NewServer()
	http.HandleFunc("/tweets", s.AddTweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
