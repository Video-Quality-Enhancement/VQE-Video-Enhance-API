package middlewares

import (
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/utils"
	"github.com/gin-gonic/gin"
)

// TODO: jwt firebase verifyIdToken
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SetUserId(c, "1234")
	}
}
