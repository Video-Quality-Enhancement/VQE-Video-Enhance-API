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
	UploadVideo(video *models.VideoEnhance, file multipart.File, fileExtension string) (string, error)
	EnhanceVideo(video *models.VideoEnhance) error
	GetVideoByRequestId(userId, requestId string) (*models.VideoEnhance, error)
	GetVideosByUserId(userId string) ([]models.VideoEnhance, error)
	DeleteVideo(userId, requestId string) error
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

func (service *videoEnhanceService) UploadVideo(video *models.VideoEnhance, file multipart.File, fileExtension string) (string, error) {

	fileName := video.RequestId + "." + fileExtension
	fileId, err := service.driveService.UploadFile(file, fileName)
	if err != nil {
		slog.Error("Error uploading video to drive", "video", video)
		return "", err
	}
	slog.Debug("Video uploaded to drive", "fileId", fileId, "requestId", video.RequestId, "userId", video.UserId)
	videoUrl := "https://drive.google.com/uc?id=" + fileId
	return videoUrl, nil

}

func (service *videoEnhanceService) EnhanceVideo(video *models.VideoEnhance) error {

	videoQuality, err := utils.IdentifyQuality(video.VideoUrl)
	if err != nil {
		slog.Error("Error identifying the quality of the video", "video", video)
		return err
	}

	video.VideoQuality = videoQuality
	video.Status = constants.VideoStatusPending.String()
	video.StatusMessage = "Video is added to the queue to be enhanced"

	err = service.repository.Create(video)
	if err != nil {
		slog.Error("Error adding video to repository", "video", video)
		return err
	}

	request := &models.VideoEnhanceRequest{
		UserId:       video.UserId,
		RequestId:    video.RequestId,
		VideoUrl:     video.VideoUrl,
		VideoQuality: video.VideoQuality,
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

func (service *videoEnhanceService) GetVideoByRequestId(userId, requestId string) (*models.VideoEnhance, error) {

	video, err := service.repository.FindByRequestId(userId, requestId)
	if err != nil {
		slog.Error("Error getting video with requestId", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Got video with requestId", "requestId", requestId, "userId", userId)
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

	slog.Debug("Deleted video", "requestId", requestId, "userId", userId)

	return nil
}
