package routes

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func testController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API call test successful",
	})
}

func RegisterUserVideoRoutes(router *gin.RouterGroup, authorization gin.HandlerFunc, controller controllers.VideoEnhanceController) {

	router.Use(authorization)
	router.GET("/test", testController)
	router.POST("/upload-and-enhance", middlewares.SetRequestID(), controller.UploadAndEnhanceVideo)
	router.POST("/enhance", middlewares.SetRequestID(), controller.EnhanceVideo)
	router.GET("/:id", controller.GetVideo)
	router.GET("/", controller.GetVideos)
	router.DELETE("/:id", controller.DeleteVideo)

}
