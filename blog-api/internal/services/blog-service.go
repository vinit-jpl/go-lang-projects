package services

import (
	"context"
	"fmt"
	"time"

	"github.com/vinit-jpl/blog-api/internal/models"
	"github.com/vinit-jpl/blog-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogService struct {
	Repo *repository.BlogRepository
}

// new blog service creates and returns a new BlogService

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{Repo: repo}
}

// CreateBlogPost handles logic for creating a blog post

func (s *BlogService) CreateBlogPost(ctx context.Context, post *models.BlogPost) error {

	post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()

	return s.Repo.CreateBlog(ctx, post)
}

func (s *BlogService) GetBlogById(ctx context.Context, id primitive.ObjectID) (*models.BlogPost, error) {
	blogPost, err := s.Repo.GetBlogById(ctx, id)

	if err != nil {
		return nil, err
	}

	return blogPost, nil
}

func (s *BlogService) GetAllBlogPosts(ctx context.Context) ([]*models.BlogPost, error) {
	blogPosts, err := s.Repo.GetAllBlogs(ctx)
	fmt.Println("blogPosts:", blogPosts)
	if err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func (s *BlogService) UpdateBlog(ctx context.Context, id primitive.ObjectID, updateData bson.M) error {
	return s.Repo.UpdateBlog(ctx, id, updateData)

}

func (s *BlogService) DeleteBlog(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.Repo.DeleteBlog(ctx, id)
}
