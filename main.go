package main

import (
	"apod-img-store/api"
	"apod-img-store/config"
	"apod-img-store/database"
	"apod-img-store/job"
	"log"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	database.Initialize(config.DBDriver, config.DBSource)

	go job.RunApodImageRetreiverCronJob()

	server := api.NewServer(config.ServerPort)
	server.Start()
}
