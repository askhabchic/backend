package models

import (
	"github.com/google/uuid"
)

type Address struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Country string
	City    string
	Street  string
}
