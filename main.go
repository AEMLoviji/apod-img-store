package main

import (
	"apod-img-store/api"
	"apod-img-store/config"
	"apod-img-store/database"
	apodprovider "apod-img-store/internal/apod_provider"
	"apod-img-store/internal/http_client"
	"apod-img-store/internal/image"
	"apod-img-store/job"
	"log"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	database.Initialize(config.DBDriver, config.DBSource)

	// run cron Job
	apodProvider := apodprovider.NewApodProvider(http_client.NewHttpClient())
	imageService := image.NewService(image.NewRepository(database.DB))
	job := job.NewApodJob(imageService, apodProvider)
	go job.Run()

	// run API Server
	server := api.NewServer(config.ServerPort)
	server.Start()
}
