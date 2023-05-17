package main

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/app"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
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

	ampq := config.NewAMQPconnection()
	defer ampq.Disconnect()

	router := gin.New()

	app.SetUpUserVideo(router, database, ampq)

	router.Run()

}
