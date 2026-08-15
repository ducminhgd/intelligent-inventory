package manufacturer

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"github.com/google/uuid"
)

type CreateManufacturerRequest struct {
	Name string `json:"name"`

	CreatedBy uuid.UUID `json:"created_by"`
}

type CreateManufacturerResponse struct {
	http.HttpResponse

	Data *domain.Manufacturer `json:"data"`
}
