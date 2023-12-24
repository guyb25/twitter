package repositories

import (
	"github.com/google/uuid"
	"twitter/internal/dal/schemas"
)

type User struct {
	users map[uuid.UUID]*schemas.User
}

func NewUser() *User {
	return &User{users: make(map[uuid.UUID]*schemas.User)}
}

func (repo *User) Create(user *schemas.User) {
	repo.users[user.Id] = user
}

func (repo *User) Delete(userId uuid.UUID) {
	delete(repo.users, userId)
}

func (repo *User) Update(user *schemas.User) {
	repo.users[user.Id] = user
}

func (repo *User) GetById(userId uuid.UUID) *schemas.User {
	return repo.users[userId]
}

func (repo *User) ExistsById(id uuid.UUID) bool {
	_, exists := repo.users[id]
	return exists
}

func (repo *User) ExistsByUsername(username string) bool {
	return repo.GetByUsername(username) != nil
}

func (repo *User) GetByUsername(username string) *schemas.User {
	for _, user := range repo.users {
		if user.Username == username {
			return user
		}
	}

	return nil
}

func (repo *User) AddTweet(userId uuid.UUID, tweetId uuid.UUID) {
	repo.users[userId].Tweets = append(repo.users[userId].Tweets, tweetId)
}
