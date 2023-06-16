package gapi

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/exp/slog"
	"google.golang.org/api/option"
)

type GoogleCloudStorage interface {
	UploadFile(file multipart.File, fileName, email string) (string, string, error)
	DeleteFile(fileName string) error
}

type googleCloudStorage struct {
	storageClient *storage.Client
	bucket        *storage.BucketHandle
	bucketName    string
}

func NewGoogleCloudStorage(bucketName string) GoogleCloudStorage {
	ctx := context.Background()

	credsPath := os.Getenv("STORAGE_SA_KEY_PATH")
	// projectId := os.Getenv("PROJECT_ID")

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		slog.Error("Error creating storage client", "error", err)
		panic(err)
	}

	bucket := client.Bucket(bucketName)

	return &googleCloudStorage{client, bucket, bucketName}
}

func (gcs *googleCloudStorage) UploadFile(file multipart.File, fileName, email string) (string, string, error) {

	object := gcs.bucket.Object(fileName)
	object = object.If(storage.Conditions{DoesNotExist: true})

	ctx := context.Background()
	wc := object.NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		slog.Error("Error copying file to storage", "error", err)
		return "", "", err
	}

	if err := wc.Close(); err != nil {
		slog.Error("Error closing file writer", "error", err)
		return "", "", err
	}

	acl := object.ACL()
	entity := storage.ACLEntity("user-" + email)
	err := acl.Set(ctx, entity, storage.RoleReader)
	if err != nil {
		slog.Error("Error granting access to file", "fileName", fileName, "error", err)
		return "", "", err
	}

	slog.Debug("Granted access to file", "fileName", fileName)

	url := "https://storage.googleapis.com/" + gcs.bucketName + "/" + fileName

	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(1 * time.Minute),
	}

	signedUrl, err := gcs.bucket.SignedURL(fileName, opts)
	if err != nil {
		slog.Error("Error getting signed URL", "error", err)
		return "", "", err
	}

	slog.Debug("Signed URL", "url", url)
	return url, signedUrl, nil

}

func (gcs *googleCloudStorage) DeleteFile(fileName string) error {

	object := gcs.bucket.Object(fileName)
	ctx := context.Background()

	attrs, err := object.Attrs(ctx)
	fmt.Println(attrs)
	fmt.Println(err)
	if err != nil {

		if err == storage.ErrObjectNotExist {
			slog.Debug("File does not exist", "fileName", fileName)
			return nil
		}

		slog.Error("Error getting file attributes", "fileName", fileName, "error", err)
		return err

	} else {

		object = object.If(storage.Conditions{GenerationMatch: attrs.Generation})
		if err := object.Delete(ctx); err != nil {
			slog.Error("Error deleting file", "fileName", fileName, "error", err)
			return err
		}

	}

	return nil

}
