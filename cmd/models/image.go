package models

import (
	"github.com/google/uuid"
)

type Image struct {
	ID    uuid.UUID `json:"id"`
	Image string    `json:"image"`
}
