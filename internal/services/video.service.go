package services

import (
	"mime/multipart"

	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/constants"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/services/gapi"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/utils"
	"golang.org/x/exp/slog"
)

type VideoEnhanceService interface {
	UploadVideo(video *models.VideoEnhanceRequest, file multipart.File, fileExtension string) (string, error)
	EnhanceVideo(video *models.VideoEnhanceRequest) error
	GetVideoEnhance(userId, requestId string) (*models.VideoEnhance, error)
	GetAllVideoEnhance(userId string) ([]models.VideoEnhance, error)
	DeleteVideoEnhance(userId, requestId string) error
}

type videoEnhanceService struct {
	repository           repositories.VideoEnhanceRepository
	videoEnhanceProducer producers.VideoEnhanceProducer
	driveService         gapi.DriveService
}

func NewVideoEnhanceService(repository repositories.VideoEnhanceRepository, producer producers.VideoEnhanceProducer) VideoEnhanceService {

	return &videoEnhanceService{
		repository:           repository,
		videoEnhanceProducer: producer,
		driveService:         gapi.NewDriveService(),
	}

}

func (service *videoEnhanceService) UploadVideo(request *models.VideoEnhanceRequest, file multipart.File, fileExtension string) (string, error) {

	fileName := request.RequestId + "." + fileExtension
	fileId, err := service.driveService.UploadFile(file, fileName)
	if err != nil {
		slog.Error("Error uploading video to drive", "request", request)
		return "", err
	}
	slog.Debug("Video uploaded to drive", "fileId", fileId, "requestId", request.RequestId, "userId", request.UserId)
	videoUrl := "https://drive.google.com/uc?id=" + fileId
	return videoUrl, nil

}

func (service *videoEnhanceService) EnhanceVideo(request *models.VideoEnhanceRequest) error {

	videoQuality, err := utils.IdentifyQuality(request.VideoUrl)
	if err != nil {
		slog.Error("Error identifying the quality of the video", "request", request)
		return err
	}

	video := &models.VideoEnhance{
		UserId:        request.UserId,
		RequestId:     request.RequestId,
		VideoUrl:      request.VideoUrl,
		VideoQuality:  videoQuality,
		Status:        constants.VideoStatusPending.String(),
		StatusMessage: "Video is added to the queue to be enhanced",
	}

	err = service.repository.Create(video)
	if err != nil {
		slog.Error("Error adding video to repository", "video", video)
		return err
	}

	request.VideoQuality = videoQuality
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

	// TODO: call delete video producer

	err := service.repository.Delete(userId, requestId)
	if err != nil {
		slog.Error("Error deleting video", "requestId", requestId)
		return err
	}

	slog.Debug("Deleted video", "requestId", requestId, "userId", userId)

	return nil
}
