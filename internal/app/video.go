package app

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/routes"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// set up user
// set up admin - admin is getting his own repository
// set up developer - developer should not be able to access any kind of video

func SetUpUserVideo(router *gin.Engine, database *mongo.Database, ampq config.AMQPconnection) {

	collection := database.Collection("VIDEO_COLLECTION")
	repository := repositories.NewVideoEnhanceRepository(collection)
	service := services.NewVideoEnhanceService(repository, ampq)
	controller := controllers.NewVideoEnhanceController(service)

	userVideoRouter := router.Group("/api/user/videos")
	routes.RegisterUserVideoRoutes(userVideoRouter, controller)

}
