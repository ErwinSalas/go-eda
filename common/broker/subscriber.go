package broker

import (
	"context"
	"fmt"

	"github.com/ErwinSalas/go-eda/common/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// SNSSubscriberAWS implementa MessageSubscriber usando AWS SNS
type SNSSubscriberAWS struct {
	client *sns.Client
	topic  string
}

func NewSNSSubscriberAWS(topic, region, endpoint string) (*SNSSubscriberAWS, error) {
	snsClient, err := utils.GetSNSClient(region, endpoint)

	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}

	if topic == "" {
		return nil, fmt.Errorf("SNS_TOPIC_ARN is not set")
	}

	return &SNSSubscriberAWS{
		client: snsClient,
		topic:  topic,
	}, nil
}

// Subscribe suscribe un endpoint a un tópico SNS
func (s *SNSSubscriberAWS) Subscribe(endpoint string, protocol string) error {
	_, err := s.client.Subscribe(context.TODO(), &sns.SubscribeInput{
		Endpoint: aws.String(endpoint),
		Protocol: aws.String(protocol),
		TopicArn: aws.String(s.topic),
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to SNS topic: %v", err)
	}
	fmt.Printf("Successfully subscribed %s to SNS topic\n", endpoint)
	return nil
}
