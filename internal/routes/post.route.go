package routes

import (
	"github.com/Deven-14/gin-mongo-trial/internal/controllers"
	"github.com/gin-gonic/gin"
)

type PostRouter interface {
	SetupRoutes()
}

type postRouter struct {
	router     *gin.RouterGroup
	controller controllers.PostController
}

func NewPostRouter(router *gin.Engine, controller controllers.PostController) PostRouter {
	return &postRouter{
		router:     router.Group("/posts"),
		controller: controller,
	}
}

func (r *postRouter) SetupRoutes() {

	r.router.POST("/", r.controller.CreatePost)
	r.router.GET("/", r.controller.GetPosts)

}
