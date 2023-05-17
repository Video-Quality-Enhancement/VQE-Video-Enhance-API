package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func GetRequestID(c *gin.Context) (string, error) {
	requestID := c.Writer.Header().Get("X-Request-ID")

	if requestID == "" {
		slog.Error("Request ID not found in header")
		return "", errors.New("request ID not found in header")
	}

	return requestID, nil
}
