package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Tweet struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetMemoryRepository struct {
	tweets []Tweet
}

func (r *TweetMemoryRepository) AddTweet(u Tweet) (int, error) {
	r.tweets = append(r.tweets, u)
	return len(r.tweets), nil
}

type TweetRepository interface {
	AddTweet(u Tweet) (int, error)
}

type server struct {
	repository TweetRepository
}

func main() {
	http.HandleFunc("/tweets", handleTweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTweets(w http.ResponseWriter, r *http.Request) {
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

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
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
	id, err := s.repository.AddTweet(payload)
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
