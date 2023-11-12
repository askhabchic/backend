package models

import (
	"github.com/google/uuid"
)

type Supplier struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Address_id   uuid.UUID `json:"address_id"`
	Phone_number string    `json:"phone_number"`
}
