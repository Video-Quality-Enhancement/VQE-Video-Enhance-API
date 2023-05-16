package controllers

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/utils"
	"github.com/gin-gonic/gin"
)

type VideoEnhanceController interface {
	UserVideoEnhanceController
	AdminVideoEnhanceController
}

type UserVideoEnhanceController interface {
	EnhanceVideo(c *gin.Context)
	// GetVideo(*gin.Context)
	// GetVideos(*gin.Context)
	// delete video enhance request
	// delete both the enhanced video and then uploaded video
}

type AdminVideoEnhanceController interface {
	GetVideoByRequestId(c *gin.Context)
	GetVideosByUserId(c *gin.Context)
	DeleteVideo(c *gin.Context)
	// add video without quota to user
	// send notification to user again
}

// developer interface

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

	video.RequestId = utils.GetRequestID(c)

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

func (controller *videoEnhanceController) GetVideosByUserId(c *gin.Context) {

	var userId string
	err := c.ShouldBindJSON(&userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videos, err := controller.service.GetVideosByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)

}

func (controller *videoEnhanceController) DeleteVideo(c *gin.Context) {
	// TODO: implement delete video controller
}
