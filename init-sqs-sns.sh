#!/bin/sh

# Create SQS Queue
aws --endpoint-url=http://localstack:4566 sqs create-queue --queue-name my-queue

# Create SNS Topic
aws --endpoint-url=http://localstack:4566 sns create-topic --name my-topic

# Subscribe SQS to SNS
aws --endpoint-url=http://localstack:4566 sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:my-topic --protocol sqs --notification-endpoint http://localstack:4566/000000000000/my-queue
