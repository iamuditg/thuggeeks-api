package repository

import (
	"database/sql"
	"thuggeeks-api/models"
)

// BlogRepository struct contains instance of db sql
type BlogRepository struct {
	db *sql.DB
}

// NewBlogRepository make connection with sql db
func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

type BlogRepositoryInterface interface {
	CreateBlogRepository(blogGeeks models.BlogGeeks) (models.BlogGeeks, error)
}

func (rp BlogRepository) CreateBlogRepository(blogGeeks models.BlogGeeks) (models.BlogGeeks, error) {
	blogCreateQuery := `
		INSERT INTO thug_geeks(title,image,description_short,description_long,tag,like_count)
		VALUES ($1,$2,$3,$4,$5,$6)
`
	_, err := rp.db.Exec(blogCreateQuery, blogGeeks.Title, blogGeeks.Image, blogGeeks.DescriptionShort,
		blogGeeks.DescriptionLong, blogGeeks.Tag, blogGeeks.LikeCount)
	if err != nil {
		return models.BlogGeeks{}, err
	}
	return blogGeeks, nil
}
