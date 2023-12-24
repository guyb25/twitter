package validation

import (
	"github.com/google/uuid"
	"twitter/internal/core/util/uuidutil"
	"twitter/internal/dal/repositories"
	"twitter/internal/dal/schemas"
)

type Update struct {
	userRepository *repositories.User
}

type UpdateValidation int

const (
	FollowersInvalid UpdateValidation = iota
	FollowingInvalid
	Valid
)

func NewUpdate(userRepository *repositories.User) *Update {
	return &Update{
		userRepository: userRepository,
	}
}

// TODO: make sure username is available when changing username
// TODO: when adding a follower/following also update the followee followingee
func (validator *Update) ValidateUserUpdate(user *schemas.User) UpdateValidation {
	if uuidutil.Exists(user.Following, user.Id) || !validator.areIdsValid(user.Following) {
		return FollowingInvalid
	}

	if uuidutil.Exists(user.Followers, user.Id) || !validator.areIdsValid(user.Followers) {
		return FollowersInvalid
	}

	return Valid
}

func (validator *Update) areIdsValid(ids []uuid.UUID) bool {
	for _, id := range ids {
		if !validator.userRepository.ExistsById(id) {
			return false
		}
	}

	return true
}
