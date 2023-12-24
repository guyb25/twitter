package providers

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"twitter/internal/core/models"
	"twitter/internal/dal/schemas"
)

type userResponse struct{}

func (res *userResponse) UsernameTaken() *models.Response {
	return models.NewResponse(http.StatusConflict, "username is taken")
}

func (res *userResponse) Created() *models.Response {
	return models.NewResponse(http.StatusCreated, "user created")
}

func (res *userResponse) Updated() *models.Response {
	return models.NewResponse(http.StatusCreated, "user updated")
}

func (res *userResponse) UserInfo(user *schemas.User) *models.Response {
	return models.NewResponse(http.StatusOK, user)
}

func (res *userResponse) NotFound() *models.Response {
	return models.NewResponse(http.StatusNotFound, "user not found")
}

func (res *userResponse) InvalidFollowers(ids []uuid.UUID) *models.Response {
	return models.NewResponse(http.StatusNotFound, fmt.Sprintf("invalid followers: %v", ids))
}

func (res *userResponse) InvalidFollowing(ids []uuid.UUID) *models.Response {
	return models.NewResponse(http.StatusNotFound, fmt.Sprintf("invalid following: %v", ids))
}

func (res *userResponse) SelfFollow() *models.Response {
	return models.NewResponse(http.StatusBadRequest, "You can't follow yourself!")
}
