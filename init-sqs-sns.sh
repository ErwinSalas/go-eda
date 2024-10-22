#!/bin/sh

# Variables
ENDPOINT_URL=http://localstack:4566

# Crear colas SQS
aws --endpoint-url=$ENDPOINT_URL sqs create-queue --queue-name payments
aws --endpoint-url=$ENDPOINT_URL sqs create-queue --queue-name orders
# aws --endpoint-url=$ENDPOINT_URL sqs create-queue --queue-name food-queue
# aws --endpoint-url=$ENDPOINT_URL sqs create-queue --queue-name delivery-queue

# Crear un único tema SNS
aws --endpoint-url=$ENDPOINT_URL sns create-topic --name order-topic
aws --endpoint-url=$ENDPOINT_URL sns create-topic --name payment-topic


echo "Queue and topic have been created and linked successfully."
