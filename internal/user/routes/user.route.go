package routes

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	interfaces.Router
}

type userRouter struct {
	router     *gin.RouterGroup
	controller controllers.UserController
}

func NewUserRouter(router *gin.RouterGroup, controller controllers.UserController) UserRouter {
	return &userRouter{router, controller}
}

func (router *userRouter) RegisterRoutes() {

}
