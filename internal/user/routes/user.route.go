package routes

import "github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"

type UserRouter interface {
	interfaces.Router
}

type userRouter struct {
}

func NewUserRouter() UserRouter {
	return &userRouter{}
}

func (router *userRouter) RegisterRoutes() {

}
