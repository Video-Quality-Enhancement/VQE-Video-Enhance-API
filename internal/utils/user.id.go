package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func SetUserId(c *gin.Context, userId string) {

	if userId == "" {
		slog.Warn("User ID missing, cannot set userId")
	} else {
		c.Set("x-userId", userId)
	}

}

func GetUserId(c *gin.Context) string {

	userId := c.GetString("x-userId")

	if userId == "" {
		slog.Warn("User ID missing, cannot get userId")
		return ""
	}

	return userId
}
