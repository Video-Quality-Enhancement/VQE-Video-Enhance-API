package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/interfaces"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
)

type VideoService interface {
	interfaces.Service
	CreateVideo(*models.Video) error
	GetVideoByCorrelationId(string) (*models.Video, error)
	GetVideosByEmail(string) ([]models.Video, error)
	UpdateVideo(string, string) error
	DeleteVideo(string) error
}
