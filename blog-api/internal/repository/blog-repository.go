package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/vinit-jpl/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *BlogRepository) GetBlogById(ctx context.Context, id primitive.ObjectID) (*models.BlogPost, error) {
	var blog models.BlogPost
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

// Find() returns a cursor, not a single document.
// The cursor represents a stream of results from the database.
// You need to iterate over them
func (r *BlogRepository) GetAllBlogs(ctx context.Context) ([]*models.BlogPost, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{}) //M unorderd representation of documnet
	fmt.Println("cursor", cursor)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var posts []*models.BlogPost

	for cursor.Next(ctx) {
		var post models.BlogPost

		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}

		posts = append(posts, &post)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}

func (r *BlogRepository) UpdateBlog(ctx context.Context, id primitive.ObjectID, updateData bson.M) error {
	filter := bson.M{"_id": id}

	update := bson.M{"$set": updateData}

	_, err := r.Collection.UpdateOne(ctx, filter, update)

	return err

}
