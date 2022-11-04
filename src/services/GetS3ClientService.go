package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type GetS3ClientService struct {
	config aws.Config
}

func NewGetS3ClientService(config aws.Config) *GetS3ClientService {
	return &GetS3ClientService{
		config: config,
	}
}

func (service *GetS3ClientService) Get() *s3.Client {
	client := s3.NewFromConfig(service.config)
	return client
}
