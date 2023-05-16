package middlewares

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/utils"
	"github.com/gin-gonic/gin"
)

// TODO: jwt firebase verifyIdToken
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SetUserId(c, "1234")
	}
}
