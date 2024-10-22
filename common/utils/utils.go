package utils

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type customSNSEndpointResolver struct {
	localstackEndpoint string
	awsRegion          string
}

func (r *customSNSEndpointResolver) ResolveEndpoint(service string, options sns.EndpointResolverOptions) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:               r.localstackEndpoint,
		SigningRegion:     r.awsRegion,
		HostnameImmutable: true,
	}, nil
}

type customSQSEndpointResolver struct {
	localstackEndpoint string
	awsRegion          string
}

func (r *customSQSEndpointResolver) ResolveEndpoint(service string, options sqs.EndpointResolverOptions) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:               r.localstackEndpoint,
		SigningRegion:     r.awsRegion,
		HostnameImmutable: true,
	}, nil
}

func GetSNSClient(region, endpoint string) (*sns.Client, error) {

	resolver := &customSNSEndpointResolver{
		localstackEndpoint: endpoint,
		awsRegion:          region,
	}

	snsClient := sns.New(sns.Options{
		Region:           region,
		EndpointResolver: resolver,
	})
	return snsClient, nil
}

func GetSQSClient(region, endpoint string) (*sqs.Client, error) {

	resolver := &customSQSEndpointResolver{
		localstackEndpoint: endpoint,
		awsRegion:          region,
	}

	sqsClient := sqs.New(sqs.Options{
		Region:           region,
		EndpointResolver: resolver,
	})
	return sqsClient, nil
}
