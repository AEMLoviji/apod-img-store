package job

import (
	apodprovider "apod-img-store/internal/apod_provider"
	"apod-img-store/internal/entity"
	"apod-img-store/internal/image"
	"context"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type ApodJob struct {
	ImgSvc       image.Service
	ApodProvider apodprovider.ApodProvider
}

func NewApodJob(imageSvc image.Service, apodProvider apodprovider.ApodProvider) *ApodJob {
	return &ApodJob{ImgSvc: imageSvc, ApodProvider: apodProvider}
}

func (c *ApodJob) Run() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().Do(func() {
		img, err := c.ApodProvider.GetImage()
		if err != nil {
			log.Println("error occurred while retreving apod image", err.Error())
		}

		if err = c.ImgSvc.CreateIfNotExist(context.Background(), image.Image{
			Image: entity.Image{
				Title:     img.Title,
				Copyright: img.Copyright,
				Url:       img.Url,
				CreatedAt: time.Time(img.Date),
			},
		}); err != nil {
			log.Println("error occurred while storing image in DB")
		}
	})

	s.StartBlocking()
}
