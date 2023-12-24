package ginutil

import "github.com/gin-gonic/gin"

func ExtractDto[T any](context *gin.Context) (T, error) {
	var dto T
	err := context.BindJSON(&dto)
	return dto, err
}
