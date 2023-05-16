package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func SetUserId(c *gin.Context, userId string) {

	if userId == "" {
		slog.Warn("User ID missing, cannot set userId")
	} else {
		c.Set("X-User-ID", userId)
	}

}

func GetUserId(c *gin.Context) string {

	userId := c.GetString("X-User-ID")

	if userId == "" {
		slog.Warn("User ID missing, cannot get userId")
		return ""
	}

	return userId
}
