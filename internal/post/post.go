package post

import (
	"github.com/Deven-14/gin-mongo-trial/internal/post/controllers"
	"github.com/Deven-14/gin-mongo-trial/internal/post/repositories"
	"github.com/Deven-14/gin-mongo-trial/internal/post/routes"
	"github.com/Deven-14/gin-mongo-trial/internal/post/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpPostModel(routerGroup *gin.RouterGroup, database *mongo.Database) {

	repository := repositories.NewPostRepository(database)
	service := services.NewPostService(repository)
	controller := controllers.NewPostController(service)
	postRouter := routes.NewPostRouter(routerGroup, controller)
	postRouter.RegisterRoutes()

}
