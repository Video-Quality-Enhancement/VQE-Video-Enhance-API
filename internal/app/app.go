package app

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpApp(router *gin.Engine, database *mongo.Database, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

	collection := database.Collection(os.Getenv("VIDEO_ENHANCE_COLLECTION"))
	videoRouter := router.Group("/api/videos/enhance")
	SetUpVideo(videoRouter, collection, conn, firebaseClient)

}

func OneTimeSetUpApp(database *mongo.Database) {

	// collection := database.Collection(os.Getenv("VIDEO_ENHANCE_COLLECTION"))
	// SetUpVideoRepositoryIndexes(collection)

	// SetUpStorageCORS()

}
