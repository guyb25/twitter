package tweet

import (
	"github.com/google/uuid"
	"sort"
	"time"
	"twitter/internal/core/models"
	"twitter/internal/core/providers"
	"twitter/internal/dal/repositories"
	"twitter/internal/dal/schemas"
	"twitter/internal/endpoints/tweet/dtos"
)

type Service struct {
	userRepository   *repositories.User
	tweetRepository  *repositories.Tweet
	responseProvider *providers.Response
}

func NewService(
	userRepository *repositories.User,
	tweetRepository *repositories.Tweet,
	responseProvider *providers.Response,
) *Service {
	return &Service{
		userRepository:   userRepository,
		tweetRepository:  tweetRepository,
		responseProvider: responseProvider,
	}
}

func (service *Service) CreateTweet(dto dtos.CreateTweet) *models.Response {
	tweet := schemas.NewTweet(uuid.New(), dto.UserId, dto.Text, time.Now())

	if !service.userRepository.ExistsById(dto.UserId) {
		return service.responseProvider.UserResponse.NotFound()
	}

	service.tweetRepository.Create(*tweet)
	service.userRepository.AddTweet(dto.UserId, tweet.Id)
	return service.responseProvider.TweetResponse.Created()
}

func (service *Service) GetFeed(dto dtos.ViewFeed) *models.Response {
	user := service.userRepository.GetById(dto.UserId)

	if user == nil {
		return service.responseProvider.UserResponse.NotFound()
	}

	tweets := make([]schemas.Tweet, 0)
	for _, followingId := range user.Following {
		tweets = append(tweets, service.tweetRepository.GetTweetsByUserId(followingId)...)
	}

	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].CreationTime.Before(tweets[j].CreationTime)
	})

	return service.responseProvider.TweetResponse.Feed(tweets)
}
