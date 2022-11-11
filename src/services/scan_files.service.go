package services

import "errors"

type ScanFilesServices struct {
}

func NewScanFilesService() *ScanFilesServices {
	return &ScanFilesServices{}
}

func (service *ScanFilesServices) Scan(path string) ([]string, error) {
	if path == "" {
		return []string{}, errors.New("Um caminho válido deve ser informado")
	}

	return []string{}, nil
}
