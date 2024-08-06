package Models

import "time"

type Transactions struct {
	TransactionId   uint      `json:"transaction_id"`
	TransactionType string    `json:"transaction_type"`
	Description     string    `json:"description"`
	Amount          uint      `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}
