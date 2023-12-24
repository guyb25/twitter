package models

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(engine *gin.Engine)
}
