package config

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func GetRequestID(c *gin.Context) string {
	requestID := c.Writer.Header().Get("X-Request-ID")

	if requestID == "" {
		slog.Warn("Request ID not found in header")
		return ""
	}

	return requestID
}
