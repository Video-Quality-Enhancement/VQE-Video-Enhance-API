package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/constants"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type VideoEnhanceService interface {
	EnhanceVideo(video *models.VideoEnhanceRequest) (*models.VideoEnhance, error)
	GetVideoEnhance(userId, requestId string) (*models.VideoEnhance, error)
	GetAllVideoEnhance(userId string) ([]models.VideoEnhance, error)
	DeleteVideoEnhance(userId, requestId string) error
}

type videoEnhanceService struct {
	repository           repositories.VideoEnhanceRepository
	videoEnhanceProducer producers.VideoEnhanceProducer
	firebaseClient       config.FirebaseInfo
	store                VideoStore
}

func NewVideoEnhanceService(repository repositories.VideoEnhanceRepository, firebaseClient config.FirebaseInfo, producer producers.VideoEnhanceProducer) VideoEnhanceService {

	return &videoEnhanceService{
		repository:           repository,
		videoEnhanceProducer: producer,
		firebaseClient:       firebaseClient,
		store:                NewVideoStore(),
	}

}

func (service *videoEnhanceService) EnhanceVideo(request *models.VideoEnhanceRequest) (*models.VideoEnhance, error) {

	videoEnhance := &models.VideoEnhance{
		UserId:    request.UserId,
		RequestId: uuid.New().String(),
		VideoId:   request.VideoId,
		Status:    constants.VideoStatusPending.String(),
	}

	video, err := service.store.GetVideo(request.VideoId, request.UserId)
	if err != nil {
		slog.Error("Error getting video from store", "videoId", request.VideoId, "userId", request.UserId)
		return nil, err
	}

	request.RequestId = videoEnhance.RequestId
	request.VideoResolution = video.VideoResolution

	err = service.repository.Create(videoEnhance)
	if err != nil {
		slog.Error("Error adding videoEnhance to repository", "videoEnhance", videoEnhance)
		return nil, err
	}

	err = service.videoEnhanceProducer.Publish(request)
	if err != nil {
		slog.Error("Error publishing video to enhance", "requestId", videoEnhance.RequestId)
		service.DeleteVideoEnhance(videoEnhance.UserId, videoEnhance.RequestId) // * important to delete the video if it failed to publish
		return nil, err
	}

	slog.Debug("Added Video to enhance", "requestId", videoEnhance.RequestId, "userId", videoEnhance.UserId)
	return videoEnhance, nil

}

func (service *videoEnhanceService) GetVideoEnhance(userId, requestId string) (*models.VideoEnhance, error) {

	video, err := service.repository.FindByRequestId(userId, requestId)
	if err != nil {
		slog.Error("Error getting video with requestId", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Got video with requestId", "requestId", requestId, "userId", userId)
	return video, nil

}

func (service *videoEnhanceService) GetAllVideoEnhance(userId string) ([]models.VideoEnhance, error) {

	videos, err := service.repository.FindAllByUserId(userId)
	if err != nil {
		slog.Error("Error getting videos of user", "userId", userId)
		return nil, err
	}

	slog.Debug("Got videos of user", "userId", userId)
	return videos, nil

}

func (service *videoEnhanceService) DeleteVideoEnhance(userId, requestId string) error {

	videoEnhance, err := service.repository.FindByRequestId(userId, requestId)
	if err != nil {
		slog.Error("Error getting video with requestId", "requestId", requestId)
		return err
	}

	err = service.store.DeleteVideo(videoEnhance.VideoId, userId)
	if err != nil {
		slog.Error("Error deleting uploaded video", "requestId", requestId)
		return err
	}

	if videoEnhance.EnhancedVideoId != "" {
		err = service.store.DeleteVideo(videoEnhance.EnhancedVideoId, userId)
		if err != nil {
			slog.Error("Error deleting enhanced video", "requestId", requestId)
			return err
		}
	}

	err = service.repository.Delete(userId, requestId)
	if err != nil {
		slog.Error("Error deleting video", "requestId", requestId)
		return err
	}

	slog.Debug("Deleted video", "requestId", requestId, "userId", userId)

	return nil
}
