package controllers

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/services"
	"github.com/gin-gonic/gin"
)

type VideoEnhanceController interface {
	UserVideoEnhanceController
	AdminVideoEnhanceController
}

type UserVideoEnhanceController interface {
	EnhanceVideo(*gin.Context)
	// GetVideo(*gin.Context)
	// GetVideos(*gin.Context)
}

type AdminVideoEnhanceController interface {
	GetVideoByRequestId(*gin.Context)
	GetVideosByEmail(*gin.Context)
	OnVideoEnhancementComplete(*gin.Context)
	DeleteVideo(*gin.Context)
}

type videoEnhanceController struct {
	service services.VideoEnhanceService
}

func NewVideoEnhanceController(service services.VideoEnhanceService) VideoEnhanceController {
	return &videoEnhanceController{service}
}

func (controller *videoEnhanceController) EnhanceVideo(c *gin.Context) {

	var video models.VideoEnhance

	err := c.ShouldBindJSON(&video)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.EnhanceVideo(&video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, video)

}

func (controller *videoEnhanceController) GetVideoByRequestId(c *gin.Context) {

	var requestId string
	err := c.ShouldBindJSON(&requestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video, err := controller.service.GetVideoByRequestId(requestId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)

}

func (controller *videoEnhanceController) GetVideosByEmail(c *gin.Context) {

	var email string
	err := c.ShouldBindJSON(&email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videos, err := controller.service.GetVideosByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)

}

func (controller *videoEnhanceController) OnVideoEnhancementComplete(c *gin.Context) {

	var response models.VideoEnhanceResponse
	err := c.ShouldBindJSON(&response) // TODO: check if this is the correct way to bind json
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.OnVideoEnhancementComplete(&response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Notification is sent to the account"})

}

func (controller *videoEnhanceController) DeleteVideo(c *gin.Context) {
	// TODO: implement delete video controller
}
