package servicestest

import (
	"testing"

	"github.com/claudiomozer/gos3/src/services"
)

func TestShouldGetAnS3Client(t *testing.T) {
	configLoader := services.NewAwsSdkConfigLoaderService()
	config, err := configLoader.LoadAwsSdkConfig()

	if err != nil {
		t.Errorf("Should not return an error on load config")
		return
	}

	sut := services.NewGetS3ClientService(config)
	client := sut.Get()

	if client == nil {
		t.Errorf("Should return a valid client")
	}
}
