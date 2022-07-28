package image

import (
	"apod-img-store/database"
	"apod-img-store/internal/entity"
	"context"
	"database/sql"
	"errors"
	"time"
)

// Repository encapsulates the logic to access images from the data source.
type Repository interface {
	GetByDate(ctx context.Context, date time.Time) (*entity.Image, error)
	List(ctx context.Context) ([]entity.Image, error)
	Create(ctx context.Context, image entity.Image) error
}

// repository persists images in database
type repository struct {
	db *sql.DB
}

// NewRepository creates a new image repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

// Get reads the image with the specified date from the database.
func (r repository) GetByDate(ctx context.Context, date time.Time) (*entity.Image, error) {
	var id string
	var title string
	var copyright string
	var url string

	row := r.db.QueryRow(`SELECT id, title, copyright, url FROM image WHERE created_at=$1;`, date)

	switch err := row.Scan(&id, &title, &copyright, &url); err {
	case nil:
		return &entity.Image{
			ID:        id,
			Title:     title,
			Copyright: copyright,
			Url:       url,
		}, nil
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	default:
		return nil, errors.New("Error occurred image found")
	}
}

// Query retrieves the image records from the database.
func (r repository) List(ctx context.Context) ([]entity.Image, error) {
	rows, err := database.DB.Query(`SELECT id, title, copyright, url FROM image`)
	if err != nil {
		return nil, err
	}

	var id string
	var title string
	var copyright string
	var url string

	var images []entity.Image
	for rows.Next() {
		if err = rows.Scan(&id, &title, &copyright, &url); err != nil {
			return nil, err
		}

		img := entity.Image{
			ID:        id,
			Title:     title,
			Copyright: copyright,
			Url:       url,
		}

		images = append(images, img)
	}

	return images, nil
}

func (r repository) Create(ctx context.Context, image entity.Image) error {
	sqlStatement := `
	INSERT INTO image (id, title, copyright, url, created_at)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, sqlStatement,
		entity.GenerateID(), image.Title, image.Copyright, image.Url, image.CreatedAt)

	return err
}
