package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/models"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/exp/slog"
)

type VideoStore interface {
	GetVideo(videoId string, userId string) (*models.Video, error)
	DeleteVideo(videoId string, userId string) error
}

type videoStore struct {
	baseUrl string
}

func NewVideoStore() *videoStore {

	baseUrl, ok := os.LookupEnv("VIDEO_STORE_BASE_URL")
	if !ok {
		slog.Error("Error getting video store base url")
		panic("Error getting video store base url")
	}

	return &videoStore{
		baseUrl: baseUrl,
	}
}

func (store *videoStore) generateA2AToken(userId string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{}
	claims["iat"] = time.Now().Unix()
	claims["uid"] = userId
	claims["sub"] = userId
	claims["iss"] = "video-enhance-api"
	claims["aud"] = "video-store-api"
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["role"] = "a2a"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("VIDEO_STORE_SECRET_KEY")))
	if err != nil {
		slog.Error("Error generating a2a token", "userId", userId)
		return "", err
	}

	return tokenString, nil
}

func (store *videoStore) GetVideo(videoId string, userId string) (*models.Video, error) {

	url := store.baseUrl + "/videos/" + videoId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Error creating get request", "url", url)
		return nil, err
	}

	token, err := store.generateA2AToken(userId)
	if err != nil {
		slog.Error("Error generating a2a token", "userId", userId)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error getting video", "url", url)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Error getting video", "url", url, "status", resp.StatusCode)
		return nil, errors.New("Error getting video")
	}

	video := &models.Video{}
	err = json.NewDecoder(resp.Body).Decode(video)
	if err != nil {
		slog.Error("Error decoding video", "url", url)
		return nil, err
	}

	return video, nil
}

func (store *videoStore) DeleteVideo(videoId string, userId string) error {

	url := store.baseUrl + "/videos/" + videoId

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		slog.Error("Error creating delete request", "url", url)
		return err
	}

	token, err := store.generateA2AToken(userId)
	if err != nil {
		slog.Error("Error generating a2a token", "userId", userId)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error deleting video", "url", url)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Error deleting video", "url", url, "status", resp.StatusCode)
		return errors.New("Error deleting video")
	}

	return nil

}
