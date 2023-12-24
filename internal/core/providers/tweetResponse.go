package providers

import (
	"fmt"
	"net/http"
	"twitter/internal/core/models"
	"twitter/internal/dal/schemas"
)

type tweetResponse struct{}

func (res *tweetResponse) Created() *models.Response {
	return models.NewResponse(http.StatusCreated, "created tweet")
}

func (res *tweetResponse) Feed(tweets []schemas.Tweet) *models.Response {
	return models.NewResponse(http.StatusCreated, fmt.Sprintf("tweets: %v", tweets))
}
