package services

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/repositories"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type VideoEnhanceService interface {
	EnhanceVideo(video *models.VideoEnhance) error
	GetVideoByRequestId(userId, requestId string) (*models.VideoEnhance, error)
	GetVideosByUserId(userId string) ([]models.VideoEnhance, error)
	DeleteVideo(userId, requestId string) error
}

type videoEnhanceService struct {
	repository           repositories.VideoEnhanceRepository
	videoEnhanceProducer producers.VideoEnhanceProducer
}

func NewVideoEnhanceService(repository repositories.VideoEnhanceRepository, ch *amqp.Channel) VideoEnhanceService {

	return &videoEnhanceService{
		repository:           repository,
		videoEnhanceProducer: producers.NewVideoEnhanceProducer(ch),
	}

}

func (service *videoEnhanceService) EnhanceVideo(video *models.VideoEnhance) error {

	err := service.repository.Create(video)
	if err != nil {
		slog.Error("Error adding video to repository", "video", video)
		return err
	}

	request := &models.VideoEnhanceRequest{
		UserId:    video.UserId,
		RequestId: video.RequestId,
		VideoUrl:  video.VideoUrl,
	}
	err = service.videoEnhanceProducer.Publish(request)
	if err != nil {
		slog.Error("Error publishing video to enhance", "requestId", video.RequestId)
		service.repository.Delete(video.UserId, video.RequestId) // * important to delete video from repository if it was not published to enhance
		return err
	}

	slog.Debug("Added Video to enhance", "requestId", video.RequestId)
	return nil

}

func (service *videoEnhanceService) GetVideoByRequestId(userId, requestId string) (*models.VideoEnhance, error) {

	video, err := service.repository.FindByRequestId(userId, requestId)
	if err != nil {
		slog.Error("Error getting video with requestId", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Got video with requestId", "requestId", requestId)
	return video, nil

}

func (service *videoEnhanceService) GetVideosByUserId(userId string) ([]models.VideoEnhance, error) {

	videos, err := service.repository.FindAllByUserId(userId)
	if err != nil {
		slog.Error("Error getting videos of user", "userId", userId)
		return nil, err
	}

	slog.Debug("Got videos of user", "userId", userId)
	return videos, nil

}

func (service *videoEnhanceService) DeleteVideo(userId, requestId string) error {

	// TODO: call delete video producer

	err := service.repository.Delete(userId, requestId)
	if err != nil {
		slog.Error("Error deleting video", "requestId", requestId)
		return err
	}

	slog.Debug("Deleted video", "requestId", requestId)

	return nil
}
