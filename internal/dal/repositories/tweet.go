package repositories

import (
	"github.com/google/uuid"
	"twitter/internal/dal/schemas"
)

type Tweet struct {
	tweets map[uuid.UUID]schemas.Tweet
}

func NewTweet() *Tweet {
	return &Tweet{tweets: make(map[uuid.UUID]schemas.Tweet)}
}

func (repo *Tweet) Create(tweet schemas.Tweet) {
	repo.tweets[tweet.Id] = tweet
}

func (repo *Tweet) GetTweetsByUserId(userId uuid.UUID) []schemas.Tweet {
	tweets := make([]schemas.Tweet, 0)

	for _, tweet := range repo.tweets {
		if tweet.Owner == userId {
			tweets = append(tweets, tweet)
		}
	}

	return tweets
}
