package services

import (
	"thuggeeks-api/models"
	"thuggeeks-api/repository"
)

type BlogService struct {
	blogRepository repository.BlogRepository
}

func NewBlogService(blogRepository repository.BlogRepository) *BlogService {
	return &BlogService{blogRepository: blogRepository}
}

type BlogServiceInterface interface {
	CreateBlogService(blogGeeks models.BlogGeeks) (models.BlogGeeks, error)
}

func (blogService BlogService) CreateBlogService(blogGeeks models.BlogGeeks) (models.BlogGeeks, error) {
	return blogService.blogRepository.CreateBlogRepository(blogGeeks)
}
