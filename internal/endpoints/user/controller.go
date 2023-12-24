package user

import (
	"github.com/gin-gonic/gin"
	"twitter/internal/core/util/ginutil"
	"twitter/internal/endpoints/user/dtos"
)

type Controller struct {
	service *Service
}

func (controller *Controller) RegisterRoutes(engine *gin.Engine) {
	engine.POST(`/user/register`, controller.Create)
	engine.GET(`/user/request/:username`, controller.Request)
	engine.PUT(`/user/update`, controller.Update)
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Create(context *gin.Context) {
	dto, _ := ginutil.ExtractDto[dtos.CreateUser](context)
	res := controller.service.CreateUser(dto)
	context.JSON(res.Status, res.Message)
}

func (controller *Controller) Request(context *gin.Context) {
	username := context.Param("username")
	res := controller.service.RequestUser(username)
	context.JSON(res.Status, res.Message)
}

func (controller *Controller) Update(context *gin.Context) {
	dto, _ := ginutil.ExtractDto[dtos.UpdateUser](context)
	res := controller.service.UpdateUser(dto)
	context.JSON(res.Status, res.Message)
}
