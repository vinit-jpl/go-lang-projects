package routes

import (
	"net/http"

	"github.com/vinit-jpl/blog-api/internal/controllers"
)

func RegisterPostRoutes(router *http.ServeMux, postController *controllers.PostController) {
	router.HandleFunc("/createPost", postController.Create)
}

func RegisterViewPostRoutes(router *http.ServeMux, postController *controllers.PostController) {
	router.HandleFunc("/viewPost", postController.ViewPostById)
}

func RegisterViewAllPostsRoutes(router *http.ServeMux, postController *controllers.PostController) {
	router.HandleFunc("/viewAllPost", postController.GetAllBlogPosts)
}

func RegisterUpdatePostRoutes(router *http.ServeMux, postController *controllers.PostController) {
	router.HandleFunc("/updatePost", postController.UpdateBlogPost)
}
