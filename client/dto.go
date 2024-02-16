package client

import "github.com/google/uuid"

type CreateClientDTO struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"client_name"`
	Surname          string    `json:"client_surname"`
	Birthday         string    `json:"birthday"`
	Gender           bool      `json:"gender"`
	RegistrationDate string    `json:"registration_date"`
	AddressId        uuid.UUID `json:"address_id"`
}
