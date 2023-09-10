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
	CreateBlogService(blogGeeks *models.BlogGeeks) (*models.BlogGeeks, error)
	GetBlogServiceById(id int) (*models.BlogGeeks, error)
	DeleteBlogServiceById(id int) error
	UpdateBlog(blogGeek *models.BlogGeeks) error
}

func (blogService BlogService) CreateBlogService(blogGeeks *models.BlogGeeks) (*models.BlogGeeks, error) {
	return blogService.blogRepository.CreateBlogRepository(blogGeeks)
}

func (blogService BlogService) GetBlogServiceById(id int) (*models.BlogGeeks, error) {
	return blogService.blogRepository.GetBlogByID(id)
}

func (blogService BlogService) DeleteBlogServiceById(id int) error {
	return blogService.blogRepository.DeleteBlogById(id)
}

func (blogService BlogService) UpdateBlog(blogGeek *models.BlogGeeks) error {
	return blogService.blogRepository.UpdateBlog(blogGeek)
}
