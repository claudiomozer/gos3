package servicestest

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/claudiomozer/gos3/src/services"
)

func TestShouldReturnAnErrorIfInvalidClientIsProvided(t *testing.T) {
	sut := services.NewListObjectsService(nil)

	_, err := sut.Read("")

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}

func TestShouldReturnAnErrorIfAnInvalidBucketIsGiven(t *testing.T) {
	sut := services.NewListObjectsService(&s3.Client{})

	_, err := sut.Read("")

	if err == nil {
		t.Errorf("Should return an error if invalid bucket is given")
	}
}

func TestShouldReturnAnErrorIfClientReturnsError(t *testing.T) {
	sut := services.NewListObjectsService(&s3.Client{})

	_, err := sut.Read("teste")

	if err == nil {
		t.Errorf("Should return an error if client returns error")
	}
}

func TestShouldReturnObjectsOnSuccess(t *testing.T) {
	config, err := services.NewAwsSdkConfigLoaderService().LoadAwsSdkConfig()
	if err == nil {
		client := services.NewGetS3ClientService(config).Get()
		sut := services.NewListObjectsService(client)

		objects, err := sut.Read("datatracking-web")

		if err != nil {
			t.Errorf("Should not return an error on success")
		}

		if len(objects) == 0 {
			t.Error("Should not return an empty array on success")
		}
	}
}
