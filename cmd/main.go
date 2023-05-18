package main

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/app"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	gin.DefaultWriter = config.NewSlogInfoWriter()
	gin.DefaultErrorWriter = config.NewSlogErrorWriter()
}

func main() {

	logFile := config.SetupSlogOutputFile()
	defer logFile.Close()

	client := config.NewMongoClient()
	database := client.ConnectToDB()
	defer client.Disconnect()

	// Try to keep the connection/channel count low. Use separate connections to publish and consume.
	// Ideally, you should have one connection per process, and then use one channel per thread in your application.
	ampq := config.NewAMQPconnection()
	defer ampq.DisconnectAll()

	router := gin.New()
	router.Use(middlewares.JSONlogger())
	router.Use(gin.Recovery())

	app.SetUpApp(router, database, ampq)

	router.Run()

}
