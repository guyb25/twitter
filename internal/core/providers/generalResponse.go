package providers

import (
	"net/http"
	"twitter/internal/core/models"
)

type generalResponse struct{}

func (res *generalResponse) InternalServerError() *models.Response {
	return models.NewResponse(http.StatusInternalServerError, "internal server error")
}
