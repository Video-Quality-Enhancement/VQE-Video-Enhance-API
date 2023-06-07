package app

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/middlewares"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/routes"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/validations"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

// set up user
// set up admin - admin is getting his own repository
// set up developer - developer should not be able to access any kind of video

func SetUpUserVideo(router *gin.RouterGroup, collection *mongo.Collection, ch *amqp.Channel, firebaseClient config.FirebaseClient) {

	repository := repositories.NewVideoEnhanceRepository(collection)
	producer := producers.NewVideoEnhanceProducer(ch)
	service := services.NewVideoEnhanceService(repository, producer)
	controller := controllers.NewVideoEnhanceController(service)
	validations.RegisterVideoValidations()
	authorization := middlewares.Authorization(firebaseClient)
	routes.RegisterUserVideoRoutes(router, authorization, controller)

}

func SetUpUserVideoRepositoryIndexes(collection *mongo.Collection) {

	repository := repositories.NewVideoEnhanceRepositorySetup(collection)
	repository.MakeUserIdIndex()
	repository.MakeRequestIdUniqueIndex()

}
