package services

import (
	"mime/multipart"
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/constants"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/services/gapi"
	"golang.org/x/exp/slog"
)

type VideoEnhanceService interface {
	UploadVideo(video *models.VideoEnhanceRequest, file multipart.File, fileExtension string) (string, string, error)
	EnhanceVideo(video *models.VideoEnhanceRequest) error
	GetVideoEnhance(userId, requestId string) (*models.VideoEnhance, error)
	GetAllVideoEnhance(userId string) ([]models.VideoEnhance, error)
	DeleteVideoEnhance(userId, requestId string) error
}

type videoEnhanceService struct {
	repository                  repositories.VideoEnhanceRepository
	videoEnhanceProducer        producers.VideoEnhanceProducer
	firebaseClient              config.FirebaseInfo
	uploadVideoStorageService   gapi.GoogleCloudStorage
	enhancedVideoStorageService gapi.GoogleCloudStorage
}

func NewVideoEnhanceService(repository repositories.VideoEnhanceRepository, firebaseClient config.FirebaseInfo, producer producers.VideoEnhanceProducer) VideoEnhanceService {

	uploadVideoBucketName := os.Getenv("UPLOAD_VIDEO_BUCKET_NAME")
	enhancedVideoBucketName := os.Getenv("ENHANCED_VIDEO_BUCKET_NAME")

	return &videoEnhanceService{
		repository:                  repository,
		videoEnhanceProducer:        producer,
		firebaseClient:              firebaseClient,
		uploadVideoStorageService:   gapi.NewGoogleCloudStorage(uploadVideoBucketName),
		enhancedVideoStorageService: gapi.NewGoogleCloudStorage(enhancedVideoBucketName),
	}

}

func (service *videoEnhanceService) UploadVideo(request *models.VideoEnhanceRequest, file multipart.File, fileExtension string) (string, string, error) {

	fileName := request.RequestId + "." + fileExtension
	email, err := service.firebaseClient.GetEmail(request.UserId)
	if err != nil {
		slog.Error("Error getting email from firebase", "request", request)
		return "", "", err
	}

	videoUrl, signedUrl, err := service.uploadVideoStorageService.UploadFile(file, fileName, email)
	if err != nil {
		slog.Error("Error uploading file to google cloud storage", "request", request)
		return "", "", err
	}

	return videoUrl, signedUrl, nil

}

func (service *videoEnhanceService) EnhanceVideo(request *models.VideoEnhanceRequest) error {

	video := &models.VideoEnhance{
		UserId:        request.UserId,
		RequestId:     request.RequestId,
		VideoUrl:      request.VideoUrl,
		VideoQuality:  request.VideoQuality,
		Status:        constants.VideoStatusPending.String(),
		StatusMessage: "Video is added to the queue to be enhanced",
	}

	err := service.repository.Create(video)
	if err != nil {
		slog.Error("Error adding video to repository", "video", video)
		return err
	}

	err = service.videoEnhanceProducer.Publish(request)
	if err != nil {
		slog.Error("Error publishing video to enhance", "requestId", video.RequestId)
		service.repository.Delete(video.UserId, video.RequestId) // * important to delete video from repository if it was not published to enhance
		return err
	}

	slog.Debug("Added Video to enhance", "requestId", video.RequestId, "userId", video.UserId)
	return nil

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

	fileName := requestId + ".mp4"
	err := service.uploadVideoStorageService.DeleteFile(fileName)
	if err != nil {
		slog.Error("Error deleting uploaded video", "requestId", requestId)
		return err
	}

	err = service.enhancedVideoStorageService.DeleteFile(fileName)
	if err != nil {
		slog.Error("Error deleting enhanced video", "requestId", requestId)
		return err
	}

	err = service.repository.Delete(userId, requestId)
	if err != nil {
		slog.Error("Error deleting video", "requestId", requestId)
		return err
	}

	slog.Debug("Deleted video", "requestId", requestId, "userId", userId)

	return nil
}
