package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Authorization(firebaseClient config.FirebaseClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := strings.Split(c.GetHeader("Authorization"), " ")[1]

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		uid, err := firebaseClient.VerifyIDToken(ctx, token)

		if err != nil {
			slog.Error("error verifying ID token", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		utils.SetUserId(c, uid)

	}
}
