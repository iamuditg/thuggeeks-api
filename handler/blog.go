package handler

import (
	"net/http"
	"thuggeeks-api/services"
)

type BlogHandler struct {
	blogService services.BlogService
}

func NewBlogHandler(blogService services.BlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

type BlogHandlerInterface interface {
	CreateBlogHandler(writer http.ResponseWriter, request *http.Request)
}

func (handler BlogHandler) CreateBlogHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}
