package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Author    string             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}
