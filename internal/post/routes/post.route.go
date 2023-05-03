package routes

import (
	"github.com/Deven-14/gin-mongo-trial/internal/post/controllers"
	"github.com/gin-gonic/gin"
)

type PostRouter interface {
	RegisterRoutes()
}

type postRouter struct {
	router     *gin.RouterGroup
	controller controllers.PostController
}

func NewPostRouter(router *gin.RouterGroup, controller controllers.PostController) PostRouter {
	return &postRouter{router, controller}
}

func (r *postRouter) RegisterRoutes() {

	r.router.POST("/", r.controller.CreatePost)
	r.router.GET("/", r.controller.GetPosts)

}
