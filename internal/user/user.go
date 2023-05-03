package user

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/helpers"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/controllers"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/routes"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/user/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpUserModel(routerGroup *gin.RouterGroup, database *mongo.Database) {
	helpers.SetUpModel(
		database,
		routerGroup,
		repositories.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
		routes.NewUserRouter,
	)
}
