package main

import (
	"fmt"
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/config"
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

	client := config.NewMongoClient()

	database := client.ConnectToDB()
	defer client.Disconnect()

	router := gin.Default()

	env := Env{
		database: database,
		router:   router,
	}

	fmt.Println(env)

	router.GET("/", helloController)

	router.Run()

}
