package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsSdkConfigLoaderService struct{}

func NewAwsSdkConfigLoaderService() *AwsSdkConfigLoaderService {
	return &AwsSdkConfigLoaderService{}
}

func (service *AwsSdkConfigLoaderService) LoadAwsSdkConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("s3-user"),
	)

	return cfg, err
}
