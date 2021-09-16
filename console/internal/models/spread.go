package models

import "github.com/google/uuid"

type Spread struct {
	Id uuid.UUID
	CompanyName string
	Spread float32
}