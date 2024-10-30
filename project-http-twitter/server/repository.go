package server

import "sync"

type Tweet struct {
	ID       int    `json:"-"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetMemoryRepository struct {
	tweets []Tweet
	lock   sync.RWMutex
}

func (r *TweetMemoryRepository) AddTweet(u Tweet) (int, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	u.ID = len(r.tweets) + 1
	r.tweets = append(r.tweets, u)
	return len(r.tweets), nil
}
func (r *TweetMemoryRepository) Tweets() ([]Tweet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.tweets, nil
}

type TweetRepository interface {
	AddTweet(u Tweet) (int, error)
	Tweets() ([]Tweet, error)
}
