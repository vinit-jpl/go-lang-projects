package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vinit-jpl/blog-api/internal/models"
	"github.com/vinit-jpl/blog-api/internal/services"
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
