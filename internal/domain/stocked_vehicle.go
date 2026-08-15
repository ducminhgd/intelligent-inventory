package domain

import (
	"time"

	"github.com/google/uuid"
)

// VehicleAction is the action/status taken on a stocked vehicle.
type VehicleAction string

const (
	ActionNone                  VehicleAction = "NONE"
	ActionPriceReductionPlanned VehicleAction = "PRICE_REDUCTION_PLANNED"
	ActionPriceReduced          VehicleAction = "PRICE_REDUCED"
	ActionDestroyed             VehicleAction = "DESTROYED"
)

// Valid reports whether a is one of the known vehicle actions.
func (a VehicleAction) Valid() bool {
	switch a {
	case ActionNone, ActionPriceReductionPlanned, ActionPriceReduced, ActionDestroyed:
		return true
	default:
		return false
	}
}

// StockedVehicle represents a vehicle in stock.
type StockedVehicle struct {
	ID  uint32 `json:"id"`
	VIN string `json:"vin"`

	ModelID uint32 `json:"model_id"`

	Name  string  `json:"name"`
	Price float64 `json:"price"`

	Action VehicleAction `json:"action"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`

	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`

	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by"`
}
