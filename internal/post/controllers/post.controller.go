package controllers

import (
	"net/http"

	"github.com/Deven-14/gin-mongo-trial/internal/interfaces"
	"github.com/Deven-14/gin-mongo-trial/internal/post/models"
	"github.com/Deven-14/gin-mongo-trial/internal/post/services"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	interfaces.Controller
	CreatePost(c *gin.Context)
	GetPosts(c *gin.Context)
}

type postController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return &postController{service}
}

func (c *postController) CreatePost(ctx *gin.Context) {

	var post models.Post

	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.service.CreatePost(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, post)

}

func (c *postController) GetPosts(ctx *gin.Context) {

	posts, err := c.service.GetPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}
