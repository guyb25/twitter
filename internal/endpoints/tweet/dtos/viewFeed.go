package dtos

import "github.com/google/uuid"

type ViewFeed struct {
	UserId uuid.UUID `json:"userId"`
}
