// Package models contains models of the domain
package models

import "github.com/google/uuid"

// Position describes a position on the market
type Position struct {
	ID            uuid.UUID `json:"id"`
	ServerID      int       `json:"server_id"`
	CompanyName   string    `json:"name"`
	Open          bool      `json:"open"`
	BuyPrice      float32   `json:"buy_price"`
	SellPrice     float32   `json:"sell_price"`
	ProfitAndLoss float32   `json:"pnl"`
}
