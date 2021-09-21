// Package models contains models of the domain
package models

import "github.com/google/uuid"

// Price describes a price of a share
type Price struct {
	ID          uuid.UUID `json:"id"`
	CompanyName string    `json:"name"`
	Price       float32   `json:"buy_price"`
}
