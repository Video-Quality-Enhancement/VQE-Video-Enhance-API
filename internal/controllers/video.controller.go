package controllers

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/utils"
	"github.com/gin-gonic/gin"
)

type VideoEnhanceController interface {
	EnhanceVideo(c *gin.Context)
	GetVideoEnhance(*gin.Context)
	GetAllVideoEnhance(*gin.Context)
	DeleteVideoEnhance(*gin.Context)
}

type videoEnhanceController struct {
	service services.VideoEnhanceService
}

func NewVideoEnhanceController(service services.VideoEnhanceService) VideoEnhanceController {
	return &videoEnhanceController{service}
}

func (controller *videoEnhanceController) EnhanceVideo(c *gin.Context) {

	var request models.VideoEnhanceRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.UserId, err = utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// make this correlation id
	// request.RequestId, err = utils.GetRequestID(c)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	videoEnhance, err := controller.service.EnhanceVideo(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, videoEnhance)

}

func (controller *videoEnhanceController) GetVideoEnhance(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var requestId = c.Param("id")
	if requestId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong requestID"})
		return
	}

	video, err := controller.service.GetVideoEnhance(userId, requestId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)

}

func (controller *videoEnhanceController) GetAllVideoEnhance(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	videos, err := controller.service.GetAllVideoEnhance(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)

}

func (controller *videoEnhanceController) DeleteVideoEnhance(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var requestId = c.Param("id")
	if requestId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "requestId is required"})
		return
	}

	err = controller.service.DeleteVideoEnhance(userId, requestId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully"})

}
