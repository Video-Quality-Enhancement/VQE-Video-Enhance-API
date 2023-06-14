package app

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Video-API/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpApp(router *gin.Engine, database *mongo.Database, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

	collection := database.Collection(os.Getenv("VIDEO_COLLECTION"))
	userVideoRouter := router.Group("/api/user/videos")
	SetUpUserVideo(userVideoRouter, collection, conn, firebaseClient)

}

func SetUpRepositoryIndexes(database *mongo.Database) {

	collection := database.Collection(os.Getenv("VIDEO_COLLECTION"))
	SetUpUserVideoRepositoryIndexes(collection)

}
