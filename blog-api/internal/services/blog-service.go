package services

import (
	"context"
	"time"

	"github.com/vinit-jpl/blog-api/internal/models"
	"github.com/vinit-jpl/blog-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
