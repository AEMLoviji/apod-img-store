package api

import (
	"net/http"

	"apod-img-store/database"
	"apod-img-store/internal/image"
)

func NewApi() *http.ServeMux {
	mux := http.NewServeMux()

	// api per resource
	imageSvc := image.NewService(image.NewRepository(database.DB))
	image.RegisterHandlers(mux, imageSvc)

	return mux
}
