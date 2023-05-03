package controllers

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/services"
)

type UserController interface {
	interfaces.Controller
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
}
