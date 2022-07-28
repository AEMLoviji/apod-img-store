package job

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunApodImageRetreiverCronJob() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().Do(func() {
		// TODO get image of the day
	})

	s.StartBlocking()
}
