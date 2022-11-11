package services

import (
	"errors"
	"io/fs"
	"io/ioutil"
)

type ScanDirectoryServices struct {
}

func NewScanDirectoryService() *ScanDirectoryServices {
	return &ScanDirectoryServices{}
}

func (service *ScanDirectoryServices) Scan(path string) ([]fs.FileInfo, error) {
	if path == "" {
		return []fs.FileInfo{}, errors.New("Um caminho válido deve ser informado")
	}

	return ioutil.ReadDir(path)
}
