package repository

import (
	"database/sql"
	"fmt"
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
	CreateBlogRepository(blogGeeks *models.BlogGeeks) (*models.BlogGeeks, error)
	GetBlogByID(id int) (*models.BlogGeeks, error)
	DeleteBlogById(id int) error
	UpdateBlog(blogGeeks *models.BlogGeeks) error
}

func (rp BlogRepository) CreateBlogRepository(blogGeeks *models.BlogGeeks) (*models.BlogGeeks, error) {
	blogCreateQuery := `
		INSERT INTO thug_geeks(title,image,description_short,description_long,tag,like_count)
		VALUES ($1,$2,$3,$4,$5,$6)
`
	_, err := rp.db.Exec(blogCreateQuery, blogGeeks.Title, blogGeeks.Image, blogGeeks.DescriptionShort,
		blogGeeks.DescriptionLong, blogGeeks.Tag, blogGeeks.LikeCount)
	if err != nil {
		return nil, err
	}
	return blogGeeks, nil
}

func (rp BlogRepository) GetBlogByID(id int) (*models.BlogGeeks, error) {
	query := `
		SELECT * FROM thug_geeks WHERE id = $1
`
	row := rp.db.QueryRow(query, id)

	var blogGeeks models.BlogGeeks
	err := row.Scan(
		&blogGeeks.ID,
		&blogGeeks.Title,
		&blogGeeks.Image,
		&blogGeeks.DescriptionShort,
		&blogGeeks.DescriptionLong,
		&blogGeeks.Tag,
		&blogGeeks.LikeCount)
	if err != nil {
		return nil, err
	}
	return &blogGeeks, nil
}

func (rp BlogRepository) DeleteBlogById(id int) error {
	query := `
		DELETE FROM thug_geeks WHERE id = $1
`
	out, err := rp.db.Exec(query, id)
	fmt.Println(out)
	return err
}

func (rp BlogRepository) UpdateBlog(blogGeeks *models.BlogGeeks) error {
	query := `
        UPDATE thug_geeks
        SET title = $1, image = $2, description_short = $3, description_long = $4, tag = $5, like_count = $6
        WHERE id = $7
    `
	_, err := rp.db.Exec(query,
		blogGeeks.Title, blogGeeks.Image, blogGeeks.DescriptionShort,
		blogGeeks.DescriptionLong, blogGeeks.Tag, blogGeeks.LikeCount, blogGeeks.ID)
	return err
}
