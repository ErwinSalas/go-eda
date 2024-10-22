package broker

import (
	"context"
	"fmt"

	"github.com/ErwinSalas/go-eda/common/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSPublisherAWS struct {
	client *sns.Client
	topic  string
}

func NewSNSPublisherAWS(topic, region, endpoint string) (*SNSPublisherAWS, error) {
	snsClient, err := utils.GetSNSClient(region, endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}
	if topic == "" {
		return nil, fmt.Errorf("SNS_TOPIC_ARN is not set")
	}

	return &SNSPublisherAWS{
		client: snsClient,
		topic:  topic,
	}, nil
}

// Publish publica un mensaje en un tópico SNS
func (p *SNSPublisherAWS) Publish(message string) error {
	_, err := p.client.Publish(context.TODO(), &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(p.topic),
	})
	if err != nil {
		return fmt.Errorf("failed to publish message to SNS: %v", err)
	}
	fmt.Println("Message published to SNS successfully")
	return nil
}
