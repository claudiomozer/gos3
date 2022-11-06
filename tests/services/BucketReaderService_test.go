package servicestest

import (
	"testing"

	"github.com/claudiomozer/gos3/src/services"
)

func TestShouldReturnAnErrorIfInvalidClientIsProvided(t *testing.T) {
	sut := services.NewBucketReaderService(nil)

	err := sut.Read("")

	if err == nil {
		t.Errorf("Should return an error if invalid client is given")
	}
}
