package stockedvehicle

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"github.com/google/uuid"
)

type UpdateStockedVehicleRequest struct {
	ID  uint32 `json:"id"`
	VIN string `json:"vin"`

	ModelID uint32 `json:"model_id"`

	Name  string  `json:"name"`
	Price float64 `json:"price"`

	Action domain.VehicleAction `json:"action"`

	UpdatedBy uuid.UUID `json:"updated_by"`
}

type UpdateStockedVehicleResponse struct {
	http.HttpResponse
	Data *domain.StockedVehicle `json:"data"`
}

type DeleteStockedVehicleRequest struct {
	ID        uint32    `json:"id"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}

type DeleteStockedVehicleResponse struct {
	http.HttpResponse
}
