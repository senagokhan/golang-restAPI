package Models

import "time"

type Card struct {
	CardId         uint      `json:"cardId"`
	CardLimit      uint      `json:"cardLimit"`
	CurrentBalance uint      `json:"currentBalance"`
	CardNumber     string    `json:"cardNumber"`
	CardType       string    `json:"cardType"`
	ExpirationDate string    `json:"expirationDate,omitempty"`
	CVV            string    `json:"cvv"`
	CreatedAt      time.Time `json:"created_at"`
}
