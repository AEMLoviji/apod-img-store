package main

import (
	"apod-img-store/api"
	"apod-img-store/config"
	"apod-img-store/database"
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

	imageService := image.NewService(image.NewRepository(database.DB))
	job := job.NewApodJob(imageService)
	go job.Run()

	server := api.NewServer(config.ServerPort)
	server.Start()
}
