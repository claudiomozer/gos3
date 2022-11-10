package services

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
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

func (service *RemoveObjectsService) Remove(bucketName string, objects *[]types.Object) error {

	if service.client == nil {
		return errors.New("Cliente fornecido é inválido")
	}

	if bucketName == "" {
		return errors.New("Nome do bucket é inválido")
	}

	if len(*objects) == 0 {
		return errors.New("Nenhum objeto fornecido para a remoção")
	}

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{
			Objects: getObjectsIdetifiers(objects),
		},
	}

	_, err := service.client.DeleteObjects(context.Background(), input)

	if err != nil {
		return err
	}

	return nil
}

func getObjectsIdetifiers(objects *[]types.Object) []types.ObjectIdentifier {
	identifiers := []types.ObjectIdentifier{}

	for _, object := range *objects {
		identifiers = append(identifiers, types.ObjectIdentifier{Key: object.Key})
	}
	return identifiers
}
