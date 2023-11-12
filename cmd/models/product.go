package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Category         string    `json:"category"`
	Price            float64   `json:"price"`
	Available_stock  int       `json:"available_stock"`
	Last_update_date time.Time `json:"last_update_date"`
	Supplier_id      uuid.UUID `json:"supplier_id"`
	Image_id         uuid.UUID `json:"image_id"`
}
