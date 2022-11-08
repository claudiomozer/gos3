package servicestest

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/claudiomozer/gos3/src/services"
)

func makeRemoveObjectsSut(client *s3.Client) *services.RemoveObjectsService {
	return services.NewRemoveObjectsService(client)
}

func TestShouldRemoveObjectsReturnAnErrorIfInvalidClientIsProvided(t *testing.T) {
	sut := makeRemoveObjectsSut(nil)

	err := sut.Remove("teste", nil)

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}

func TestShouldRemoveObjectsReturnsErrorIfEmptyObjectsIsProvided(t *testing.T) {
	sut := makeRemoveObjectsSut(&s3.Client{})

	err := sut.Remove("teste", nil)

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}

func TestShouldRemoveObjectsReturnAnErrorIfClientReturnsError(t *testing.T) {
	sut := makeRemoveObjectsSut(&s3.Client{})

	err := sut.Remove("teste", &[]types.Object{{}})

	if err == nil {
		t.Errorf("Should return an error if client returns error")
	}
}

func TestShouldRemoveObjectsDeleteObjectsOnSuccess(t *testing.T) {
	config, err := services.NewAwsSdkConfigLoaderService().LoadAwsSdkConfig()
	if err == nil {
		client := services.NewGetS3ClientService(config).Get()
		listObjectsService := makeListObjectsSut(client)
		sut := makeRemoveObjectsSut(client)

		objects, err := listObjectsService.Read("datatracking-web")

		if err != nil {
			t.Errorf("Should not return an error on success")
		}

		if len(objects) == 0 {
			t.Error("Should not return an empty array on success")
		}

		err = sut.Remove("datatracking-web", &objects)

		if err != nil {
			t.Errorf("Should not return an error on success")
		}
	}
}
