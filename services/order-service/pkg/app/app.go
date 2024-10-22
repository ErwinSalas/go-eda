package app

import (
	"github.com/ErwinSalas/go-eda/common/broker"
	"github.com/ErwinSalas/go-eda/common/queue"
	"github.com/ErwinSalas/go-eda/services/order-service/pkg/datastore"
)

type App struct {
	Datastore        datastore.IDatastore
	OrderPublisher   *broker.SNSPublisherAWS
	PaymentsConsumer *queue.SQSService
}

func NewApp(datastore datastore.IDatastore, publisher *broker.SNSPublisherAWS, consumer *queue.SQSService) *App {
	return &App{
		Datastore:        datastore,
		OrderPublisher:   publisher,
		PaymentsConsumer: consumer,
	}
}
