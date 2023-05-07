package services

import (
	"log"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/producers"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/repositories"
	"github.com/google/uuid"
)

type VideoEnhanceService interface {
	VideoEnhanceRequestService
	VideoEnhanceResponseService
	VideoDeleteService
	GetVideoByRequestId(string) (*models.VideoEnhance, error)
	GetVideosByEmail(string) ([]models.VideoEnhance, error)
}

type VideoEnhanceRequestService interface {
	EnhanceVideo(*models.VideoEnhance) error
}

type VideoEnhanceResponseService interface {
	OnVideoEnhancementComplete(string, string) error
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
		log.Println("Error adding video ", video)
		return err
	}
	go service.videoEnhanceProducer.PublishVideo(video)

	log.Println("Added Video ", video.RequestId)
	return nil

}

func (service *videoEnhanceService) GetVideoByRequestId(requestId string) (*models.VideoEnhance, error) {

	video, err := service.repository.FindByRequestId(requestId)
	if err != nil {
		log.Println("Error getting video with id: ", requestId)
		return nil, err
	}

	log.Println("Got video with id: ", requestId)
	return video, nil

}

func (service *videoEnhanceService) GetVideosByEmail(email string) ([]models.VideoEnhance, error) {

	videos, err := service.repository.FindByEmail(email)
	if err != nil {
		log.Println("Error getting videos of ", email)
		return nil, err
	}

	log.Println("Got videos of ", email)
	return videos, nil

}

func (service *videoEnhanceService) OnVideoEnhancementComplete(requestId string, EnhancedVideoUri string) error {

	err := service.repository.Update(requestId, EnhancedVideoUri)
	if err != nil {
		log.Println("Error updating video with id: ", requestId)
		return err
	}

	video, err := service.GetVideoByRequestId(requestId)
	if err != nil {
		log.Println("Error getting video with id: ", requestId)
		return err
	}

	go service.notificationProducer.PublishNotification(video)

	log.Println("Updated video with id: ", requestId)
	return nil

}

func (service *videoEnhanceService) DeleteVideo(string) error {
	// ? should this delete both the original and enhanced video?
	return nil
}
