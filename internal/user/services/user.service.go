package services

import "github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"

type UserService interface {
	interfaces.Service
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}
