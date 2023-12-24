package schemas

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Username  string
	Password  string
	Followers []uuid.UUID
	Following []uuid.UUID
	Tweets    []uuid.UUID // would be a BST
}

func NewUser(userId uuid.UUID, username string, password string) *User {
	return &User{
		Id:        userId,
		Username:  username,
		Password:  password,
		Followers: make([]uuid.UUID, 0),
		Following: make([]uuid.UUID, 0),
		Tweets:    make([]uuid.UUID, 0),
	}
}
