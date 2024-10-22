package datastore

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.comErwinSalas/go-eda/services/payment-service/pkg/types"
)

type PaymentStorer interface {
	InsertPayment(ctx context.Context, payment *types.Payment) (int64, error)
	GetLastPayment(ctx context.Context) (*types.Payment, error)
	Migrate(ctx context.Context) error
	GetPayments(ctx context.Context) ([]*types.Payment, error)
	UpdatePaymentAmount(ctx context.Context, id int, amount float64) error
}

type Payment struct {
	database *sql.DB
}

func NewPaymentStore(db *sql.DB) (*Payment, error) {
	return &Payment{database: db}, nil
}

const (
	insertPaymentQuery  = "INSERT INTO payments (order_id, transaction_id, amount, timestamp) VALUES (?, ?, ?, ?)"
	getLastPaymentQuery = "SELECT id, order_id, transaction_id, amount, timestamp FROM payments ORDER BY id DESC LIMIT 1"
	migrateQuery        = `
        CREATE TABLE IF NOT EXISTS payments (
            id INT AUTO_INCREMENT PRIMARY KEY,
            order_id VARCHAR(255),
            transaction_id VARCHAR(255),
            amount DECIMAL(10, 2),
            timestamp TIMESTAMP
        )
    `
	getPaymentsQuery         = "SELECT id, order_id, transaction_id, amount, timestamp FROM payments"
	updatePaymentAmountQuery = "UPDATE payments SET amount = ? WHERE id = ?"
)

func (d *Payment) InsertPayment(ctx context.Context, payment *types.Payment) (int64, error) {
	result, err := d.database.ExecContext(ctx, insertPaymentQuery, payment.OrderID, payment.TransactionID, payment.Amount, payment.Timestamp)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Payment) GetLastPayment(ctx context.Context) (*types.Payment, error) {
	payment := &types.Payment{}
	var timestamp sql.NullTime

	err := d.database.QueryRowContext(ctx, getLastPaymentQuery).Scan(&payment.ID, &payment.OrderID, &payment.TransactionID, &payment.Amount, &timestamp)
	if err != nil {
		return nil, err
	}

	if timestamp.Valid {
		payment.Timestamp = timestamp.Time
	}

	return payment, nil
}

func (d *Payment) Migrate(ctx context.Context) error {
	_, err := d.database.ExecContext(ctx, migrateQuery)
	return err
}

func (d *Payment) GetPayments(ctx context.Context) ([]*types.Payment, error) {
	rows, err := d.database.QueryContext(ctx, getPaymentsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*types.Payment

	for rows.Next() {
		payment := &types.Payment{}
		var timestamp sql.NullTime

		err = rows.Scan(&payment.ID, &payment.OrderID, &payment.TransactionID, &payment.Amount, &timestamp)
		if err != nil {
			return nil, err
		}

		if timestamp.Valid {
			payment.Timestamp = timestamp.Time
		}

		payments = append(payments, payment)
	}
	return payments, nil
}

func (d *Payment) UpdatePaymentAmount(ctx context.Context, id int, amount float64) error {
	_, err := d.database.ExecContext(ctx, updatePaymentAmountQuery, amount, id)
	return err
}
