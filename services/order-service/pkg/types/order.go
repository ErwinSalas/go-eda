package types

import "time"

type Order struct {
	ID              string    `json:"id"`
	CustomerDetails string    `json:"customer_details"`
	Items           []string  `json:"items"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
