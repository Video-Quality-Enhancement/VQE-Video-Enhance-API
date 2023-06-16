package middlewares

import (
	"net/http"
	"strings"

	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Authorization(firebaseClient config.FirebaseAuth) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := strings.Split(c.GetHeader("Authorization"), " ")[1]

		uid, err := firebaseClient.VerifyIDToken(token)

		if err != nil {
			slog.Error("error verifying ID token", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		utils.SetUserId(c, uid)

	}
}
