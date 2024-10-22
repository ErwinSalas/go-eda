package queue

import (
	"context"
	"fmt"

	"github.com/ErwinSalas/go-eda/common/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// QueueService defines the interface for SQS operations
type QueueService interface {
	SendMessage(ctx context.Context, body []byte) (string, error)
	ReceiveMessage(ctx context.Context, deserialize func([]byte, string) error) (bool, error)
	DeleteMessage(ctx context.Context, receiptHandle *string) error
}

type SQSService struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSService(queueURL, region, endpoint string) (*SQSService, error) {
	sqsClient, err := utils.GetSQSClient(region, endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}
	if queueURL == "" {
		return nil, fmt.Errorf("SNS_TOPIC_ARN is not set")
	}
	return &SQSService{
		client:   sqsClient,
		queueURL: queueURL,
	}, nil
}

func (s *SQSService) SendMessage(ctx context.Context, body []byte) (string, error) {
	messageOut, err := s.client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(body)),
		QueueUrl:    aws.String(s.queueURL),
	})
	if err != nil {
		return "", fmt.Errorf("unable to send message: %v", err)
	}
	return *messageOut.MessageId, nil
}

func (s *SQSService) ReceiveMessage(ctx context.Context, deserialize func([]byte, string) error) (bool, error) {
	output, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(s.queueURL),
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     10, // Long polling
	})
	if err != nil {
		return false, fmt.Errorf("unable to receive message: %v", err)
	}
	if len(output.Messages) > 0 {
		messageBody := []byte(*output.Messages[0].Body)
		receiptH := output.Messages[0].ReceiptHandle
		err := deserialize(messageBody, *receiptH)
		if err != nil {
			return true, fmt.Errorf("unable to deserialize message: %v", err)
		}

	}

	return false, nil
}

func (s *SQSService) DeleteMessage(ctx context.Context, receiptHandle *string) error {
	_, err := s.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(s.queueURL),
		ReceiptHandle: receiptHandle,
	})
	if err != nil {
		return fmt.Errorf("unable to delete message: %v", err)
	}
	return nil
}
