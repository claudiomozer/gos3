package services

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketReaderService struct {
	client *s3.Client
}

func NewBucketReaderService(client *s3.Client) *BucketReaderService {
	return &BucketReaderService{
		client: client,
	}
}

func (service *BucketReaderService) Read(bucketName string) error {
	if service.client == nil {
		return errors.New("Cliente fornecido é inválido")
	}

	if bucketName == "" {
		return errors.New("Bucket fornecido é inválido")
	}

	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	_, err := service.client.ListObjectsV2(context.Background(), params)

	return err
}
