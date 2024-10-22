package app

import (
	"github.com/ErwinSalas/go-eda/common/awsSNS"
	"github.com/ErwinSalas/go-eda/services/order-service/pkg/datastore"
)

type App struct {
	Datastore      datastore.IDatastore
	OrderPublisher awsSNS.SNSPublisherAWS
}

func NewApp(datastore datastore.IDatastore) *App {
	return &App{
		Datastore: datastore,
	}
}
