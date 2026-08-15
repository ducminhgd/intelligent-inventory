package domain

import (
	"time"

	"github.com/google/uuid"
)

// Model represents a vehicle model. A model belongs to exactly one manufacturer.
type Model struct {
	ID             uint32 `json:"id"`
	ManufacturerID uint32 `json:"manufacturer_id"`
	Name           string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`

	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`

	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by"`
}
