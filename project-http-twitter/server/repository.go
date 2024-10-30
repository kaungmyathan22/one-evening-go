package server

type Tweet struct {
	ID       int    `json:"-"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetMemoryRepository struct {
	tweets []Tweet
}

func (r *TweetMemoryRepository) AddTweet(u Tweet) (int, error) {
	u.ID = len(r.tweets) + 1
	r.tweets = append(r.tweets, u)
	return len(r.tweets), nil
}
func (r *TweetMemoryRepository) Tweets() ([]Tweet, error) {
	return r.tweets, nil
}

type TweetRepository interface {
	AddTweet(u Tweet) (int, error)
	Tweets() ([]Tweet, error)
}
