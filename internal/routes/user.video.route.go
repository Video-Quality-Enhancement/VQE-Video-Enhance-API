package routes

import (
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/controllers"
	"github.com/gin-gonic/gin"
)

type UserVideoRouter interface {
	RegisterRoutes()
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
