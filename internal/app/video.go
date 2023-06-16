package app

import (
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/middlewares"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/routes"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/validations"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// set up user
// set up admin - admin is getting his own repository
// set up developer - developer should not be able to access any kind of video

func SetUpVideo(router *gin.RouterGroup, collection *mongo.Collection, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

	repository := repositories.NewVideoEnhanceRepository(collection)
	producer := producers.NewVideoEnhanceProducer(conn)
	service := services.NewVideoEnhanceService(repository, firebaseClient, producer)
	controller := controllers.NewVideoEnhanceController(service)
	validations.RegisterVideoValidations()
	authorization := middlewares.Authorization(firebaseClient)
	routes.RegisterVideoRoutes(router, authorization, controller)

}

func SetUpVideoRepositoryIndexes(collection *mongo.Collection) {

	repository := repositories.NewVideoEnhanceRepositorySetup(collection)
	repository.MakeUserIdIndex()
	repository.MakeRequestIdUniqueIndex()

}
