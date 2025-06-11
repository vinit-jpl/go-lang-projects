package routes

import (
	"net/http"

	"github.com/vinit-jpl/blog-api/internal/controllers"
)

func RegisterPostRoutes(router *http.ServeMux, postController *controllers.PostController) {
	router.HandleFunc("/post", postController.Create)
}
