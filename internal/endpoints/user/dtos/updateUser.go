package dtos

import "github.com/google/uuid"

type UpdateUser struct {
	Id        uuid.UUID   `json:"id"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	Followers []uuid.UUID `json:"followers"`
	Following []uuid.UUID `json:"following"`
}

func NewUpdateUser(
	userId uuid.UUID,
	username string,
	password string,
	followers []uuid.UUID,
	following []uuid.UUID) *UpdateUser {
	return &UpdateUser{
		Id:        userId,
		Username:  username,
		Password:  password,
		Followers: followers,
		Following: following,
	}
}
