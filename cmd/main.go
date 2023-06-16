package main

import (
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/app"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	gin.DefaultWriter = config.NewSlogInfoWriter()
	gin.DefaultErrorWriter = config.NewSlogErrorWriter()
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func main() {

	logFile := config.SetupSlogOutputFile()
	defer logFile.Close()

	router := gin.New()
	router.Use(middlewares.JSONlogger())
	router.Use(gin.Recovery())
	// ? router.MaxMultipartMemory = 8 << 20
	configurations := cors.DefaultConfig()
	configurations.AllowAllOrigins = true
	configurations.AllowCredentials = true
	configurations.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS"}
	configurations.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"}
	configurations.ExposeHeaders = []string{"Content-Length"}
	router.Use(cors.New(configurations))
	// router.Use(CORSMiddleware())

	client := config.NewMongoClient()
	database := client.ConnectToDB()
	defer client.Disconnect()

	// Try to keep the connection/channel count low. Use separate connections to publish and consume.
	// Ideally, you should have one connection per process, and then use one channel per thread in your application.
	conn := config.NewAMQPconnection()
	defer conn.DisconnectAll()

	firebaseClient := config.NewFirebaseClient()

	app.SetUpApp(router, database, conn, firebaseClient)

	router.Run()

}
