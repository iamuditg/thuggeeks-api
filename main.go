package main

import (
	"fmt"
	"net/http"
	"thuggeeks-api/config"
	"thuggeeks-api/handler"
	"thuggeeks-api/repository"
	"thuggeeks-api/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// create db connection
	dbConnection, err := config.CreateDbConnection()
	if err != nil {
		return
	}
	fmt.Println("connection establish: ", dbConnection)

	// create connection between handler, services, repo in blog
	blogRepository := repository.NewBlogRepository(dbConnection)
	blogService := services.NewBlogService(*blogRepository)
	blogHandler := handler.NewBlogHandler(*blogService)
	// blog api
	router.HandleFunc("/createBlog", blogHandler.CreateBlogHandler)
	router.HandleFunc("/getBlogById/{id:[0-9]+}", blogHandler.GetBlogHandlerByID)
	router.HandleFunc("/updateBlog/{id:[0-9]+}", blogHandler.UpdateBlogHandler)
	router.HandleFunc("/deleteBlog/{id:[0-9]+}", blogHandler.DeleteBlogHandler)

	http.ListenAndServe(":8080", router)
}
