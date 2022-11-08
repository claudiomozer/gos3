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

	err := sut.Remove(nil)

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}

func TestShouldRemoveObjectsReturnsErrorIfEmptyObjectsIsProvided(t *testing.T) {
	sut := makeRemoveObjectsSut(&s3.Client{})

	err := sut.Remove(nil)

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}

func TestShouldRemoveObjectsReturnAnErrorIfClientReturnsError(t *testing.T) {
	sut := makeRemoveObjectsSut(&s3.Client{})

	err := sut.Remove(&[]types.Object{{}})

	if err == nil {
		t.Errorf("Should return an error if client returns error")
	}
}
