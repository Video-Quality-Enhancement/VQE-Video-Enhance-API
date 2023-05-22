package main

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/app"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	client := config.NewMongoClient()
	database := client.ConnectToDB()
	defer client.Disconnect()

	app.SetUpRepositoryIndexes(database)

}
