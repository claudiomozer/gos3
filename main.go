package main

import (
	"log"

	"github.com/claudiomozer/gos3/src/services"
)

func main() {
	configLoader := services.NewAwsSdkConfigLoaderService()
	config, err := configLoader.LoadAwsSdkConfig()

	if err != nil {
		handleError(err)
	}

	s3ClientGetter := services.NewGetS3ClientService(config)
	s3ClientGetter.Get()
}

func handleError(err error) {
	log.Fatalf("Error: %v\n", err)
}
