package stockedvehicle

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"github.com/google/uuid"
)

type CreateStockedVehicleRequest struct {
	VIN string `json:"vin"`

	ModelID uint32 `json:"model_id"`

	Name  string       `json:"name"`
	Price domain.Price `json:"price"`

	Action domain.VehicleAction `json:"action"`

	CreatedBy uuid.UUID `json:"created_by"`
}

type CreateStockedVehicleResponse struct {
	http.HttpResponse

	Data *domain.StockedVehicle `json:"data"`
}
