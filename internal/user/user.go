package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpUserModel(routerGroup *gin.RouterGroup, database *mongo.Database) {
	// helpers.SetUpModel(
	// 	database,
	// 	routerGroup,
	// 	repositories.NewUserRepository,
	// 	services.NewUserService,
	// 	controllers.NewUserController,
	// 	routes.NewUserRouter,
	// )

	// proof that repositories.NewUserRepository() cannot be used if UserRepository doesn't implement interfaces.Repository, same for services.NewUserService() and controllers.NewUserController()
}
