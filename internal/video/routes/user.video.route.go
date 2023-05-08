package routes

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/controllers"
	"github.com/gin-gonic/gin"
)

type UserVideoRouter interface {
	interfaces.Router
}

type userVideoRouter struct {
	routerGroup *gin.RouterGroup
	controller  controllers.UserVideoEnhanceController
}

func NewUserVideoRouter(routerGroup *gin.RouterGroup, controller controllers.UserVideoEnhanceController) UserVideoRouter {
	return &userVideoRouter{routerGroup, controller}
}

func (r *userVideoRouter) RegisterRoutes() {
	r.routerGroup.POST("/enhance", r.controller.EnhanceVideo)
	// r.routerGroup.GET("/:id", r.controller.GetVideo)
	// r.routerGroup.GET("/", r.controller.GetVideos)
}
