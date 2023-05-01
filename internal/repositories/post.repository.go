package repositories

import (
	"context"
	"fmt"

	"github.com/Deven-14/gin-mongo-trial/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	// Create a new post
	Create(post *models.Post) error
	// Get all posts
	FindAll() ([]models.Post, error)
}

type postRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(database *mongo.Database) PostRepository {

	collection := database.Collection("posts")

	return &postRepository{collection}
}

func (r *postRepository) Create(post *models.Post) error {

	inserted, err := r.collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}

	fmt.Println("Inserted 1 post: ", inserted.InsertedID)
	return nil

}

func (r *postRepository) FindAll() ([]models.Post, error) {

	var posts []models.Post

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var post models.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}
