package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Repository TweetRepository
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func HandleTweets(w http.ResponseWriter, r *http.Request) {
	payload := Tweet{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := json.Marshal(struct {
		ID int `json:"ID"`
	}{ID: 1})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Tweet: `%s` from %s\n", payload.Message, payload.Location)
	w.Write(p)
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()
	if r.Method == "GET" {
		tweets, err := s.Repository.Tweets()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		p, err := json.Marshal(tweetsList{Tweets: tweets})
		if err != nil {
			log.Println("Failed to marshal:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(p)
		return
	}
	payload := Tweet{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := s.Repository.AddTweet(payload)
	if err != nil {
		log.Println("Failed to add user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p, err := json.Marshal(struct {
		ID int `json:"ID"`
	}{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Tweet: `%s` from %s\n", payload.Message, payload.Location)
	w.Write(p)
}

func NewServer() Server {
	return Server{
		Repository: &TweetMemoryRepository{},
	}
}
