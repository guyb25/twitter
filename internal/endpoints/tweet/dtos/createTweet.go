package dtos

import "github.com/google/uuid"

type CreateTweet struct {
	UserId uuid.UUID `json:"userId"`
	Text   string    `json:"text"`
}
