package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	Name           string
	Category       string
	Price          float64
	AvailableStock int
	LastUpdateDate time.Time
	SupplierId     uuid.UUID
	ImageId        uuid.UUID
}
