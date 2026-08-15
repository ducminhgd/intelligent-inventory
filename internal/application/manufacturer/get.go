package manufacturer

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type GetManufacturerRequest struct {
	ID uint32 `json:"id" binding:"required"`
}

type GetManufacturerResponse struct {
	http.HttpResponse
	Data *domain.Manufacturer `json:"data"`
}

type ListManufacturerResponse struct {
	http.HttpResponse
	Data []*domain.Manufacturer `json:"data"`
}

type ListManufacturerRequest struct {
	http.ListRequest

	ID_In []uint32 `json:"id_in" form:"id_in"`

	Name_iLike string `json:"name_ilike" form:"name_ilike"`
}
