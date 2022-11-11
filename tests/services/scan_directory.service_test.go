package servicestest

import (
	"testing"

	"github.com/claudiomozer/gos3/src/services"
)

func TestShouldReturnAnErrorIfAnInvalidIsGiven(t *testing.T) {
	sut := services.NewScanDirectoryService()
	_, err := sut.Scan("")

	if err == nil {
		t.Error("Should return an error if a invalid path is given")
	}
}

func TestShouldReturnAListOfFilesOnSuccess(t *testing.T) {
	sut := services.NewScanDirectoryService()
	files, err := sut.Scan("../..")

	if err != nil {
		t.Error("Should not return an error on success")
	}

	if len(files) == 0 {
		t.Error("Should return an array of FileInfo on success")
	}
}
