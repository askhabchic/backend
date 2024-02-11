package supplier

import (
	"github.com/google/uuid"
)

type Supplier struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	AddressId   uuid.UUID
	PhoneNumber string
}
