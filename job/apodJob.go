package job

import (
	"apod-img-store/internal/image"
	"time"

	"github.com/go-co-op/gocron"
)

type ApodJob struct {
	ImgSvc image.Service
}

func NewApodJob(imageSvc image.Service) *ApodJob {
	return &ApodJob{ImgSvc: imageSvc}
}

func (c *ApodJob) Run() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().Do(func() {
		// TODO get from apod
		// TODO store in DB
		// c.ImgSvc.CreateIfNotExist(context.Background(),nil,)
	})

	s.StartBlocking()
}
