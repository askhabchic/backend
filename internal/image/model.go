package image

import (
	"github.com/google/uuid"
)

type Image struct {
	ID    uuid.UUID `gorm:"primaryKey"`
	Image string
}
