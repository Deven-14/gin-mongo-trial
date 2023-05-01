package main

import (
	"net/http"

	"github.com/Deven-14/gin-mongo-trial/internal/config"
	"github.com/Deven-14/gin-mongo-trial/internal/controllers"
	"github.com/Deven-14/gin-mongo-trial/internal/repositories"
	"github.com/Deven-14/gin-mongo-trial/internal/routes"
	"github.com/Deven-14/gin-mongo-trial/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Env struct {
	database *mongo.Database
	router   *gin.Engine
}

type Repositories interface {
	repositories.PostRepository // Add other repositories here
}

type Services interface {
	services.PostService // Add other services here
}

type Controllers interface {
	controllers.PostController // Add other controllers here
}

type Routes interface {
	routes.PostRouter // Add other routes here
}

func SetUpModel[R Routes, C Controllers, S Services, Re Repositories](env Env, newRepository func(*mongo.Database) Re, newService func(Re) S, newController func(S) C, newRouter func(*gin.Engine, C) R) {

	repository := newRepository(env.database)
	service := newService(repository)
	controller := newController(service)
	router := newRouter(env.router, controller)
	router.SetupRoutes()

}

func init() {
	config.LoadEnvVariables()
}

func main() {

	client := config.NewMongoClient()

	database := client.Connect()
	defer client.Disconnect()

	// Create a new instance of the environment
	env := Env{
		database: database,
		router:   gin.Default(),
	}

	env.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	SetUpModel(env, repositories.NewPostRepository, services.NewPostService, controllers.NewPostController, routes.NewPostRouter)

	// Run the server
	env.router.Run()

}
