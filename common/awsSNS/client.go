package awsSNS

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type customEndpointResolver struct {
	localstackEndpoint string
	awsRegion          string
}

func (r *customEndpointResolver) ResolveEndpoint(service string, options sns.EndpointResolverOptions) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:               r.localstackEndpoint,
		SigningRegion:     r.awsRegion,
		HostnameImmutable: true,
	}, nil
}

func getClient() (*sns.Client, error) {
	awsRegion := os.Getenv("AWS_REGION")
	localstackEndpoint := os.Getenv("LOCALSTACK_ENDPOINT")

	resolver := &customEndpointResolver{
		localstackEndpoint: localstackEndpoint,
		awsRegion:          awsRegion,
	}

	sqsClient := sns.New(sns.Options{
		Region:           awsRegion,
		EndpointResolver: resolver,
	})
	return sqsClient, nil
}
