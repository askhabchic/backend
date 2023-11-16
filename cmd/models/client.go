package models

import (
	"github.com/google/uuid"
)

type Client struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"client_name"`
	Surname           string    `json:"client_surname"`
	Birthday          string    `json:"birthday"`
	Gender            bool      `json:"gender"`
	Registration_date string    `json:"registration_date"`
	Address_id        uuid.UUID `json:"address_id"`
}
