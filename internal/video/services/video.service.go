package services

import (
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/video/models"
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/video/producers"
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/video/repositories"
	"golang.org/x/exp/slog"
)

type VideoEnhanceService interface {
	videoEnhanceRequestService
	VideoDeleteService
}

type videoEnhanceRequestService interface { // ? should this be exported?
	EnhanceVideo(video *models.VideoEnhance) error
	GetVideoByRequestId(requestId string) (*models.VideoEnhance, error)
	GetVideosByUserId(userId string) ([]models.VideoEnhance, error)
}

type VideoDeleteService interface {
	DeleteVideo(requestId string) error
}

type videoEnhanceService struct {
	repository           repositories.VideoEnhanceRepository
	videoEnhanceProducer producers.VideoEnhanceProducer
}

func NewVideoEnhanceService(repository repositories.VideoEnhanceRepository) VideoEnhanceService {

	return &videoEnhanceService{
		repository:           repository,
		videoEnhanceProducer: producers.NewVideoEnhanceProducer(),
	}

}

func (service *videoEnhanceService) EnhanceVideo(video *models.VideoEnhance) error {

	err := service.repository.Create(video)
	if err != nil {
		slog.Error("Error adding video to repository", "video", video)
		return err
	}

	request := &models.VideoEnhanceRequest{
		RequestId:        video.RequestId,
		UploadedVideoUri: video.UploadedVideoUri,
	}
	go service.videoEnhanceProducer.PublishVideo(request)

	slog.Debug("Added Video to enhance", "requestId", video.RequestId)
	return nil

}

func (service *videoEnhanceService) GetVideoByRequestId(requestId string) (*models.VideoEnhance, error) {

	video, err := service.repository.FindByRequestId(requestId)
	if err != nil {
		slog.Error("Error getting video with requestId", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Got video with requestId", "requestId", requestId)
	return video, nil

}

func (service *videoEnhanceService) GetVideosByUserId(userId string) ([]models.VideoEnhance, error) {

	videos, err := service.repository.FindByUserId(userId)
	if err != nil {
		slog.Error("Error getting videos of user", "userId", userId)
		return nil, err
	}

	slog.Debug("Got videos of user", "userId", userId)
	return videos, nil

}

func (service *videoEnhanceService) DeleteVideo(requestId string) error {
	// ? should this delete both the original and enhanced video?
	return nil
}
