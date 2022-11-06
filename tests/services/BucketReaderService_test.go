package servicestest

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/claudiomozer/gos3/src/services"
)

func TestShouldReturnAnErrorIfInvalidClientIsProvided(t *testing.T) {
	sut := services.NewBucketReaderService(nil)

	err := sut.Read("")

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}

func TestShouldReturnAnErrorIfAnInvalidBucketIsGiven(t *testing.T) {
	sut := services.NewBucketReaderService(&s3.Client{})

	err := sut.Read("")

	if err == nil {
		t.Errorf("Should return an error if invalid bucket is given")
	}
}

func TestShouldReturnAnErrorIfClientReturnsError(t *testing.T) {
	sut := services.NewBucketReaderService(&s3.Client{})

	err := sut.Read("teste")

	if err == nil {
		t.Errorf("Should return an error if client returns error")
	}
}
