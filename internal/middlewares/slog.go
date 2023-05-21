package middlewares

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func JSONlogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		userId, _ := utils.GetUserId(c)

		attributes := []slog.Attr{
			slog.String("gin-env", os.Getenv("GIN_ENV")),
			slog.String("service-name", os.Getenv("SERVICE_NAME")),
			slog.Int("status", c.Writer.Status()),
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("user-id", userId),
			slog.String("ip", c.ClientIP()),
			slog.Duration("latency", latency),
			slog.String("user-agent", c.Request.UserAgent()),
		}

		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			slog.LogAttrs(context.Background(), slog.LevelWarn, c.Errors.String(), attributes...)
		case c.Writer.Status() >= http.StatusInternalServerError:
			slog.LogAttrs(context.Background(), slog.LevelError, c.Errors.String(), attributes...)
		default:
			slog.LogAttrs(context.Background(), slog.LevelInfo, "Incoming request", attributes...)
		}

	}
}
