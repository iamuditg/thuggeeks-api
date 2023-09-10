package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"thuggeeks-api/models"
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
	GetBlogHandlerByID(writer http.ResponseWriter, request *http.Request)
	UpdateBlogHandler(writer http.ResponseWriter, request *http.Request)
	DeleteBlogHandler(writer http.ResponseWriter, request *http.Request)
}

func (handler BlogHandler) CreateBlogHandler(writer http.ResponseWriter, request *http.Request) {
	var blogGeeks *models.BlogGeeks
	err := json.NewDecoder(request.Body).Decode(&blogGeeks)
	if err != nil {
		http.Error(writer, "Invalid request Payload", http.StatusBadRequest)
		return
	}
	blogGeeks, err = handler.blogService.CreateBlogService(blogGeeks)
	if err != nil {
		http.Error(writer, "Failed to create thug geek", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(blogGeeks)
	if err != nil {
		log.Fatalf("Error in Sending Response JsonEncoder: %v", err)
	}
}

func (handler BlogHandler) GetBlogHandlerByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	blogGeeks, err := handler.blogService.GetBlogServiceById(id)
	if err != nil {
		http.Error(writer, "Failed to fetch blog geek", http.StatusInternalServerError)
		return
	}

	if blogGeeks == nil {
		http.Error(writer, "Blog geek not found", http.StatusNotFound)
		return
	}

	// Serialise the blogGeek to JSON and send it as response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogGeeks)
	if err != nil {
		log.Fatalf("Error in Sending Response JsonEncoder: %v", err)
	}
}

func (handler BlogHandler) UpdateBlogHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateBlogGeek models.BlogGeeks
	if err := json.NewDecoder(request.Body).Decode(&updateBlogGeek); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updateBlogGeek.ID = id
	if err := handler.blogService.UpdateBlog(&updateBlogGeek); err != nil {
		http.Error(writer, "Failed to update blog geek", http.StatusInternalServerError)
		return
	}
	// Serialise the blogGeek to JSON and send it as response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updateBlogGeek)
	if err != nil {
		log.Fatalf("Error in Sending Response JsonEncoder: %v", err)
	}
}

func (handler *BlogHandler) DeleteBlogHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err = handler.blogService.DeleteBlogServiceById(id); err != nil {
		http.Error(writer, "Failed to delete blog geek", http.StatusInternalServerError)
		return
	}

	// Serialise the blogGeek to JSON and send it as response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("Successful deleted")
	if err != nil {
		log.Fatalf("Error in Sending Response JsonEncoder: %v", err)
	}
}
