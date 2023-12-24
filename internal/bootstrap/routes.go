package bootstrap

import (
	"github.com/gin-gonic/gin"
	"twitter/internal/core/models"
	"twitter/internal/core/providers"
	"twitter/internal/dal/repositories"
	"twitter/internal/endpoints/tweet"
	"twitter/internal/endpoints/user"
	"twitter/internal/endpoints/user/validation"
)

type Routes struct {
	controllers []models.Controller
}

func NewRoutes() *Routes {
	return &Routes{}
}

func (routes *Routes) instantiateControllers() {
	userRepository := repositories.NewUser()
	tweetRepository := repositories.NewTweet()
	responseProvider := providers.NewResponse()

	userController := user.NewController(
		user.NewService(
			userRepository, responseProvider, validation.NewUpdate(userRepository),
		),
	)

	tweetController := tweet.NewController(
		tweet.NewService(userRepository, tweetRepository, responseProvider),
	)

	routes.controllers = []models.Controller{userController, tweetController}
}

func (routes *Routes) initializeControllers(engine *gin.Engine) {
	for _, controller := range routes.controllers {
		controller.RegisterRoutes(engine)
	}
}

func (routes *Routes) Bootstrap() *gin.Engine {
	router := gin.Default()
	routes.instantiateControllers()
	routes.initializeControllers(router)
	return router
}
