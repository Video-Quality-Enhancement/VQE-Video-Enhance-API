package main

import (
	"fmt"
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Env struct {
	database *mongo.Database
	router   *gin.Engine
}

func init() {
	config.LoadEnvVariables()
}

func helloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {

	logFile := config.SetupSlogOutputFile()
	defer logFile.Close()

	client := config.NewMongoClient()
	database := client.ConnectToDB()
	defer client.Disconnect()

	gin.DefaultWriter = config.NewSlogInfoWriter()
	gin.DefaultErrorWriter = config.NewSlogErrorWriter()

	router := gin.New()

	router.Use(middlewares.JSONlogger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Authorization())
	router.Use(middlewares.SetRequestID()) // TODO: can move the request id inside and use it only for the create video endpoint

	env := Env{
		database: database,
		router:   router,
	}

	fmt.Println(env)

	router.GET("/", helloController)

	router.Run()

}
