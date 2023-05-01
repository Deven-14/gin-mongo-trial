package services

import (
	"github.com/Deven-14/gin-mongo-trial/internal/models"
	"github.com/Deven-14/gin-mongo-trial/internal/repositories"
)

type PostService interface {
	// Create a new post
	CreatePost(post *models.Post) error
	// Get all posts
	GetPosts() ([]models.Post, error)
}

type postService struct {
	repository repositories.PostRepository
}

func NewPostService(repository repositories.PostRepository) PostService {
	return &postService{repository}
}

func (s *postService) CreatePost(post *models.Post) error {
	err := s.repository.Create(post)
	if err != nil {
		return err
	}
	// additional business logic
	return nil
}

func (s *postService) GetPosts() ([]models.Post, error) {
	posts, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	// additional business logic
	return posts, nil
}
