package main

import (
	"fmt"
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/middlewares"
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

	router.Use(middlewares.SetRequestID())
	router.Use(middlewares.JSONlogger())
	router.Use(gin.Recovery())

	env := Env{
		database: database,
		router:   router,
	}

	fmt.Println(env)

	router.GET("/", helloController)

	router.Run()

}
