package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vinit-jpl/blog-api/internal/models"
	"github.com/vinit-jpl/blog-api/internal/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController struct {
	Service *services.BlogService
}

func NewPostController(service *services.BlogService) *PostController {
	return &PostController{Service: service}
}

// New post controller returns a new PostController

func (pc *PostController) Create(w http.ResponseWriter, r *http.Request) {
	var post models.BlogPost
	fmt.Println("in post controller")
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := pc.Service.CreateBlogPost(r.Context(), &post)

	if err != nil {
		http.Error(w, "Failed to create a post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (pc *PostController) ViewPostById(w http.ResponseWriter, r *http.Request) {

	// first parse the query param "id"
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// convert string to mongoDb object iD

	objId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	// call the service to fetch the blog post

	post, err := pc.Service.GetBlogById(context.Background(), objId)

	if err != nil {
		http.Error(w, "blog post not found", http.StatusNotFound)
		return
	}

	// respond with the blog post as JSON

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(post)

}

func (pc *PostController) GetAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	posts, err := pc.Service.GetAllBlogPosts(ctx)
	fmt.Println("posts:", posts)

	if err != nil {
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)

}

func (pc *PostController) UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Creates a generic map to hold JSON key-value pairs.

	// string is the key (like "title"), and interface{} is the value (which could be any type — string, int, etc.).
	var body map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
	}

	updateFields := bson.M{} // Creates an empty MongoDB update map (type alias of map[string]interface{}).

	/*
		value := body["title"]
		title, ok := value.(string)

		if ok {
			updateFields["title"] = title
		}
	*/
	if title, ok := body["Title"].(string); ok {
		updateFields["Title"] = title
	}

	if content, ok := body["Content"].(string); ok {
		updateFields["Content"] = content
	}

	if author, ok := body["Author"].(string); ok {
		updateFields["Author"] = author
	}

	if len(updateFields) == 0 {
		http.Error(w, "No valid fields to update", http.StatusBadRequest)
		return
	}

	err = pc.Service.UpdateBlog(r.Context(), objectID, updateFields)
	if err != nil {
		http.Error(w, "Failed to update blog post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post updated successfully"))

}
