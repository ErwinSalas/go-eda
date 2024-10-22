package app

import (
	"github.com/ErwinSalas/go-eda/common/broker"
	"github.com/ErwinSalas/go-eda/common/queue"
	"github.comErwinSalas/go-eda/services/payment-service/pkg/datastore"
)

type App struct {
	PaymentPublisher *broker.SNSPublisherAWS
	OrdersConsumer   *queue.SQSService
	Datastore        *datastore.Payment
}

func NewApp(storer *datastore.Payment, publisher *broker.SNSPublisherAWS, consumer *queue.SQSService) *App {
	return &App{
		Datastore:        storer,
		PaymentPublisher: publisher,
		OrdersConsumer:   consumer,
	}
}
