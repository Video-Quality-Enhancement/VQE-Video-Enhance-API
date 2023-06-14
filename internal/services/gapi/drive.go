package gapi

import (
	"context"
	"mime/multipart"
	"os"

	drive "google.golang.org/api/drive/v3"

	"golang.org/x/exp/slog"
	"google.golang.org/api/option"
)

type DriveService interface {
	UploadFile(file multipart.File, fileName string) (string, error)
}

type driveService struct {
	service *drive.Service
}

func NewDriveService() DriveService {

	service, err := drive.NewService(
		context.Background(),
		option.WithCredentialsFile(os.Getenv("DRIVE_SA_KEY_PATH")),
		option.WithScopes(drive.DriveScope),
	)

	if err != nil {
		slog.Error("Unable to create drive Client", "error", err)
	}

	return &driveService{service}

}

func (d *driveService) UploadFile(file multipart.File, fileName string) (string, error) {

	f := &drive.File{
		Name:    fileName,
		Parents: []string{os.Getenv("DRIVE_UPLOAD_VIDEO_FOLDER_ID")},
	}
	res, err := d.service.Files.
		Create(f).
		Media(file).
		ProgressUpdater(func(now, size int64) { slog.Debug("Uploading "+fileName, "now", now, "size", size) }).
		Do()
	if err != nil {
		slog.Error("Error uploading file to drive", "error", err)
		return "", err
	}
	slog.Debug("Video Uploaded to drive", "file Id", res.Id)
	return res.Id, nil

}
