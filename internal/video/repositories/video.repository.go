package repositories

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoRepository interface {
	interfaces.Repository
	Create(models.Video) error
	FindAll(email string) ([]models.Video, error)
	Update(string, models.Video) error
	Delete(models.Video) error
}

type videoRepository struct {
	database *mongo.Database
}

func NewVideoRepository(database *mongo.Database) VideoRepository {
	return &videoRepository{database}
}

func (repository *videoRepository) Create(video models.Video) error {
	return nil
}

func (repository *videoRepository) FindAll(email string) ([]models.Video, error) {
	return []models.Video{}, nil
}

func (repository *videoRepository) Update(id string, video models.Video) error {
	return nil
}

func (repository *videoRepository) Delete(models.Video) error {
	return nil
}
