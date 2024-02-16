package client

import (
	"github.com/google/uuid"
)

type Client struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	Name             string
	Surname          string
	Birthday         string
	Gender           bool
	RegistrationDate string
	AddressId        uuid.UUID
}
