package main

import (
	"net/http"

	"github.com/Deven-14/gin-mongo-trial/internal/config"
	"github.com/Deven-14/gin-mongo-trial/internal/post"
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

	database := client.Connect()
	defer client.Disconnect()

	env := Env{
		database: database,
		router:   gin.Default(),
	}

	env.router.GET("/", helloController)

	post.SetUpPostModel(env.router.Group("/posts"), database)

	env.router.Run()

}
