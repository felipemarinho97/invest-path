package util

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// GetEnvOrFail - helper for get environment variable or exit with code 1
func GetEnvOrFail(envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		log.Fatalf("error | empty variable %s", envName)
	}
	return env
}

func endpointResolver(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:               "http://localstack:4566",
		HostnameImmutable: true,
		PartitionID:       "aws",
		SigningRegion:     "us-east-1",
	}, nil
}

// LoadAWSConfig - create a awsConfig using default config, also add endpointResolver in Development
func LoadAWSConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return aws.Config{}, err
	}
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(endpointResolver)
	}
	return cfg, nil
}
