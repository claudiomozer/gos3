package services

import "errors"

type ScanDirectoryServices struct {
}

func NewScanDirectoryService() *ScanDirectoryServices {
	return &ScanDirectoryServices{}
}

func (service *ScanDirectoryServices) Scan(path string) ([]string, error) {
	if path == "" {
		return []string{}, errors.New("Um caminho válido deve ser informado")
	}

	return []string{}, nil
}
