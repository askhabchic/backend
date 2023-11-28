package dto

import "github.com/google/uuid"

type CreateAddressDTO struct {
	ID      uuid.UUID `json:"id"`
	Country string    `json:"country"`
	City    string    `json:"city"`
	Street  string    `json:"street"`
}
