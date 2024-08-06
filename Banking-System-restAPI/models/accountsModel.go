package Models

import (
	"time"
)

type Accounts struct {
	AccountId      uint      `json:"accountId"`
	AccountNumber  uint      `json:"accountNumber"`
	AccountType    uint      `json:"accountType"`
	AccountBalance float64   `json:"accountBalance"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
