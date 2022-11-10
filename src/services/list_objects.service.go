package services

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListObjectsService struct {
	client *s3.Client
}

func NewListObjectsService(client *s3.Client) *ListObjectsService {
	return &ListObjectsService{
		client: client,
	}
}

func (service *ListObjectsService) Read(bucketName string) ([]types.Object, error) {
	if service.client == nil {
		return nil, errors.New("Cliente fornecido é inválido")
	}

	if bucketName == "" {
		return nil, errors.New("Bucket fornecido é inválido")
	}

	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	output, err := service.client.ListObjectsV2(context.Background(), params)

	if err != nil {
		return nil, err
	}

	objects := output.Contents

	return objects, err
}
