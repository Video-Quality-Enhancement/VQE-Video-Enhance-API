package repositories

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	interfaces.Repository
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}
