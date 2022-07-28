package image

import (
	"apod-img-store/internal/entity"
	"context"
	"database/sql"
	"time"
)

// Service encapsulates logic for images.
type Service interface {
	GetByDate(ctx context.Context, date time.Time) (*Image, error)
	List(ctx context.Context) ([]Image, error)
	CreateIfNotExist(ctx context.Context, date time.Time, image Image) error
}

// Image represents the data about an image.
type Image struct {
	entity.Image
}

type service struct {
	repo Repository
}

// NewService creates a new Image service.
func NewService(repo Repository) Service {
	return &service{repo}
}

// GetByDate returns the image with the specified date.
func (s *service) GetByDate(ctx context.Context, date time.Time) (*Image, error) {
	img, err := s.repo.GetByDate(ctx, date)
	if err != nil {
		return nil, err
	}

	return &Image{*img}, nil
}

// List returns the images.
func (s *service) List(ctx context.Context) ([]Image, error) {
	items, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := []Image{}
	for _, item := range items {
		result = append(result, Image{item})
	}

	return result, nil
}

// List returns the images.
func (s *service) CreateIfNotExist(ctx context.Context, date time.Time, image Image) error {
	_, err := s.repo.GetByDate(ctx, date)
	if err != sql.ErrNoRows {
		return nil
	}

	err = s.repo.Create(ctx, image.Image)
	if err != nil {
		return err
	}

	return nil
}
