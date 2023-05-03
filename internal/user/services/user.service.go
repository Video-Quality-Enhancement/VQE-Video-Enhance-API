package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/repositories"
)

type UserService interface {
	interfaces.Service
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}
