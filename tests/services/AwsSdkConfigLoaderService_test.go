package servicestest

import (
	"testing"

	"github.com/claudiomozer/gos3/src/services"
)

func TestShouldLoadConfigWithCorrectUser(t *testing.T) {
	sut := services.NewAwsSdkConfigLoaderService()
	_, err := sut.LoadAwsSdkConfig()

	if err != nil {
		t.Error("Should not return ")
	}
}
