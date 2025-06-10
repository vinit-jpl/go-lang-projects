package repository

import (
	"context"
	"time"

	"github.com/vinit-jpl/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	Collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) *BlogRepository {
	return &BlogRepository{
		Collection: db.Collection("posts"),
	}
}

func (r *BlogRepository) CreateBlog(ctx context.Context, post *models.BlogPost) error {
	post.ID = models.NewObjectID()
	post.CreatedAt = time.Now()

	_, err := r.Collection.InsertOne(ctx, post)
	return err
}
