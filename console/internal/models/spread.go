package models

import "github.com/google/uuid"

// Spread describes profit or loss from one share
type Spread struct {
	Id uuid.UUID
	CompanyName string
	Spread float32
}