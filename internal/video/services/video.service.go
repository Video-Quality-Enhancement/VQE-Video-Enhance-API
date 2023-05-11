package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/producers"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/repositories"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type VideoEnhanceService interface {
	videoEnhanceRequestService
	VideoEnhanceResponseService
	VideoDeleteService
}

type videoEnhanceRequestService interface { // ? should this be exported?
	EnhanceVideo(*models.VideoEnhance) error
	GetVideoByRequestId(string) (*models.VideoEnhance, error)
	GetVideosByEmail(string) ([]models.VideoEnhance, error)
}

type VideoEnhanceResponseService interface {
	OnVideoEnhancementComplete(response *models.VideoEnhanceResponse) error
}

type VideoDeleteService interface {
	DeleteVideo(string) error
}

type videoEnhanceService struct {
	repository           repositories.VideoEnhanceRepository
	videoEnhanceProducer producers.VideoEnhanceProducer
	notificationProducer producers.NotificationProducer
}

func NewVideoEnhanceService(repository repositories.VideoEnhanceRepository) VideoEnhanceService {

	return &videoEnhanceService{
		repository:           repository,
		videoEnhanceProducer: producers.NewVideoEnhanceProducer(),
		notificationProducer: producers.NewNotificationProducer(),
	}

}

func (service *videoEnhanceService) EnhanceVideo(video *models.VideoEnhance) error {

	video.RequestId = uuid.NewString()

	err := service.repository.Create(video)
	if err != nil {
		slog.Error("Error adding video ", video)
		return err
	}

	request := &models.VideoEnhanceRequest{
		RequestId:        video.RequestId,
		UploadedVideoUri: video.UploadedVideoUri,
	}
	go service.videoEnhanceProducer.PublishVideo(request)

	slog.Debug("Added Video ", video.RequestId)
	return nil

}

func (service *videoEnhanceService) GetVideoByRequestId(requestId string) (*models.VideoEnhance, error) {

	video, err := service.repository.FindByRequestId(requestId)
	if err != nil {
		slog.Error("Error getting video with id: ", requestId)
		return nil, err
	}

	slog.Debug("Got video with id: ", requestId)
	return video, nil

}

func (service *videoEnhanceService) GetVideosByEmail(email string) ([]models.VideoEnhance, error) {

	videos, err := service.repository.FindByEmail(email)
	if err != nil {
		slog.Error("Error getting videos of ", email)
		return nil, err
	}

	slog.Debug("Got videos of ", email)
	return videos, nil

}

func (service *videoEnhanceService) OnVideoEnhancementComplete(response *models.VideoEnhanceResponse) error {

	err := service.repository.Update(response)
	if err != nil {
		slog.Error("Error updating video with id: ", response.RequestId)
		return err
	}

	video, err := service.GetVideoByRequestId(response.RequestId)
	if err != nil {
		slog.Error("Error getting video with id: ", response.RequestId)
		return err
	}

	go service.notificationProducer.PublishNotification(video)

	slog.Debug("Updated video with id: ", response.RequestId)
	return nil

}

func (service *videoEnhanceService) DeleteVideo(requestId string) error {
	// ? should this delete both the original and enhanced video?
	return nil
}
