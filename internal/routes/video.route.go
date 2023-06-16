package routes

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func testController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API call test successful",
	})
}

func RegisterVideoRoutes(router *gin.RouterGroup, authorization gin.HandlerFunc, controller controllers.VideoEnhanceController) {

	router.Use(authorization)
	router.GET("/test", testController)
	router.POST("/upload", middlewares.SetRequestID(), controller.UploadAndEnhanceVideo)
	router.POST("/", middlewares.SetRequestID(), controller.EnhanceVideo)
	router.GET("/:id", controller.GetVideoEnhance)
	router.GET("/", controller.GetAllVideoEnhance)
	router.DELETE("/:id", controller.DeleteVideoEnhance)

}
