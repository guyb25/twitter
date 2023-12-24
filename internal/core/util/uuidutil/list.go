package uuidutil

import (
	"github.com/google/uuid"
)

func Exists(list []uuid.UUID, id uuid.UUID) bool {
	for _, idFromList := range list {
		if idFromList == id {
			return true
		}
	}

	return false
}
