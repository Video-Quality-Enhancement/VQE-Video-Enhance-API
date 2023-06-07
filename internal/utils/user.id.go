package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func SetUserId(c *gin.Context, userId string) error {

	if userId == "" {
		slog.Error("User ID missing, cannot set userId")
		return errors.New("user id missing, cannot set userId")
	}

	c.Set("X-User-ID", userId)
	return nil

}

func GetUserId(c *gin.Context) (string, error) {

	userId := c.GetString("X-User-ID")

	if userId == "" {
		slog.Error("User ID missing, cannot get userId")
		return "", errors.New("user id missing, cannot get userId")
	}

	return userId, nil
}
