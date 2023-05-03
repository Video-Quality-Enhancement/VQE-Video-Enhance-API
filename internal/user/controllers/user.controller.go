package controllers

import "github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"

type UserController interface {
	interfaces.Controller
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}
