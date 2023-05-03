package repositories

import "github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"

type UserRepository interface {
	interfaces.Repository
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
