package Models

import "time"

type Investment struct {
	InvestmentId          uint      `json:"InvestmentId"`
	InvestmentNumber      uint      `json:"InvestmentNumber"`
	InvestmentBalance     uint      `json:"InvestmentBalance"`
	InvestmentType        string    `json:"InvestmentType"`
	InvestmentAccountType string    `json:"InvestmentAccountType"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
