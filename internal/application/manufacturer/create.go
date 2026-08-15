package manufacturer

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type CreateManufacturerRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateManufacturerResponse struct {
	http.HttpResponse

	Data *domain.Manufacturer `json:"data"`
}
