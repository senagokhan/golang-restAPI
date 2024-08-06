package Models

import (
	"time"
)

type Customer struct {
	CustomerId uint      `json:"customerId" gorm:"type:uint;primary_key;AUTO_INCREMENT"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Passcode   string    `json:"passcode"`
}
