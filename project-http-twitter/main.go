package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.NewServer()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/tweets", s.ListTweets)
	r.With(httprate.LimitByIP(10, time.Minute)).Group(func(r chi.Router) {
		r.Post("/tweets", s.AddTweet)
	})
	go spamTweets()
	log.Fatal(http.ListenAndServe(":8080", r))

}

func spamTweets() {
	for {
		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		marshaledPayload, err := json.Marshal(addTweetPayload)
		// ...

		url := "http://localhost:8080/tweets"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
	}
}
