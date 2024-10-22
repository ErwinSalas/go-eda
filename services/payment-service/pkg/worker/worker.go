package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.comErwinSalas/go-eda/services/payment-service/pkg/app"
	"github.comErwinSalas/go-eda/services/payment-service/pkg/types"
)

type PaymentWorker interface {
	ProcessOrderPayment(ctx context.Context) error
}

type Payment struct {
	app *app.App
}

func NewPaymentWorker(app *app.App) PaymentWorker {
	return &Payment{
		app: app,
	}
}

func (p *Payment) ProcessOrderPayment(ctx context.Context) error {
	for {
		hasMessages, err := p.app.OrdersConsumer.ReceiveMessage(ctx, func(message []byte, receipt string) error {

			payment := &types.Payment{}
			err := json.Unmarshal(message, &payment)
			if err != nil {
				return err
			}

			_, err = p.app.Datastore.InsertPayment(ctx, payment)
			if err != nil {
				return err
			}

			err = p.app.OrdersConsumer.DeleteMessage(ctx, &receipt)
			if err != nil {
				return err
			}
			return nil

		})
		if err != nil {
			return err
		}
		if !hasMessages {
			time.Sleep(3 * time.Second)
		}

	}
}
