package dto

import "github.com/google/uuid"

type CreateImageDTO struct {
	ID    uuid.UUID `json:"id"`
	Image string    `json:"image"`
}
