package models

import "github.com/google/uuid"

// Price describes a price of a share
type Price struct {
	ID        uuid.UUID `json:"id"`
	CompanyName string    `json:"name"`
	BuyPrice  float32   `json:"buy_price"`
	SellPrice float32   `json:"sell_price"`
}
