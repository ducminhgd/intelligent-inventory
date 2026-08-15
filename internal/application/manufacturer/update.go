package manufacturer

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type UpdateManufacturerRequest struct {
	ID   uint32 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateManufacturerResponse struct {
	http.HttpResponse
	Data *domain.Manufacturer `json:"data"`
}

type DeleteManufacturerRequest struct {
	ID uint32 `json:"id" binding:"required"`
}

type DeleteManufacturerResponse struct {
	http.HttpResponse
}
