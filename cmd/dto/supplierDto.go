package dto

import "github.com/google/uuid"

type CreateSupplierDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	AddressId   uuid.UUID `json:"address_id"`
	PhoneNumber string    `json:"phone_number"`
}
