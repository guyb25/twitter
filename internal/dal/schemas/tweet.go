package schemas

import (
	"github.com/google/uuid"
	"time"
)

type Tweet struct {
	Id           uuid.UUID
	Owner        uuid.UUID
	Text         string
	CreationTime time.Time
}

func NewTweet(id uuid.UUID, owner uuid.UUID, text string, creationTime time.Time) *Tweet {
	return &Tweet{
		Id:           id,
		Owner:        owner,
		Text:         text,
		CreationTime: creationTime,
	}
}
