package services

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type RemoveObjectsService struct {
	client *s3.Client
}

func NewRemoveObjectsService(client *s3.Client) *RemoveObjectsService {
	return &RemoveObjectsService{
		client: client,
	}
}

func (service *RemoveObjectsService) Remove(objects []types.Object) error {

	if service.client == nil {
		return errors.New("Cliente fornecido é inválido")
	}

	return nil
}
