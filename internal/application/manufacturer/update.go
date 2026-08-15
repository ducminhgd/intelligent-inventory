package manufacturer

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"github.com/google/uuid"
)

type UpdateManufacturerRequest struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`

	UpdatedBy uuid.UUID `json:"updated_by"`
}

type UpdateManufacturerResponse struct {
	http.HttpResponse
	Data *domain.Manufacturer `json:"data"`
}

type DeleteManufacturerRequest struct {
	ID        uint32    `json:"id"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}

type DeleteManufacturerResponse struct {
	http.HttpResponse
}
