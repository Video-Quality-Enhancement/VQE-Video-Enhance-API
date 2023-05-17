package app

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/routes"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/services"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

// set up user
// set up admin - admin is getting his own repository
// set up developer - developer should not be able to access any kind of video

func SetUpUserVideo(router *gin.RouterGroup, collection *mongo.Collection, ch *amqp.Channel) {

	repository := repositories.NewVideoEnhanceRepository(collection)
	service := services.NewVideoEnhanceService(repository, ch)
	controller := controllers.NewVideoEnhanceController(service)
	routes.RegisterUserVideoRoutes(router, controller)

}
