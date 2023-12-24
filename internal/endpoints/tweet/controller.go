package tweet

import (
	"github.com/gin-gonic/gin"
	"twitter/internal/core/util/ginutil"
	"twitter/internal/endpoints/tweet/dtos"
)

type Controller struct {
	service *Service
}

func (controller *Controller) RegisterRoutes(engine *gin.Engine) {
	engine.POST(`/tweet/create`, controller.Create)
	engine.POST(`/tweet/feed/view`, controller.GetFeed)
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Create(context *gin.Context) {
	dto, _ := ginutil.ExtractDto[dtos.CreateTweet](context)
	res := controller.service.CreateTweet(dto)
	context.JSON(res.Status, res.Message)
}

func (controller *Controller) GetFeed(context *gin.Context) {
	dto, _ := ginutil.ExtractDto[dtos.ViewFeed](context)
	res := controller.service.GetFeed(dto)
	context.JSON(res.Status, res.Message)
}
