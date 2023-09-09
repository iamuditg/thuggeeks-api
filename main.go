package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"thuggeeks-api/handler"
	"thuggeeks-api/repository"
	"thuggeeks-api/services"
)

func main() {
	router := mux.NewRouter()

	// create connection between handler, services, repo in blog
	blogRepository := repository.NewBlogRepository(nil)
	blogService := services.NewBlogService(*blogRepository)
	blogHandler := handler.NewBlogHandler(*blogService)
	// create blog
	router.HandleFunc("/createBlog", blogHandler.CreateBlogHandler)

	http.ListenAndServe(":8080", router)
}
