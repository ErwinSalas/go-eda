package types

import "time"

// Payment represents a payment in the system.
type Payment struct {
	ID            int       `json:"id"`
	OrderID       string    `json:"order_id"`
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}
