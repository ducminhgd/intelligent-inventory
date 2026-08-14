package manufacturer

import (
	"context"

	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
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

type GetManufacturer struct {
	repo port.ManufacturerRepository
}

func NewGetManufacturer(repo port.ManufacturerRepository) *GetManufacturer {
	return &GetManufacturer{repo: repo}
}

func (g *GetManufacturer) Execute(ctx context.Context, req GetManufacturerRequest) (GetManufacturerResponse, error) {
	manufacturer, err := g.repo.GetByID(ctx, req.ID)
	if err != nil {
		return GetManufacturerResponse{}, err
	}

	return GetManufacturerResponse{
		Data: manufacturer,
	}, nil
}

type ListManufacturer struct {
	repo port.ManufacturerRepository
}

func NewListManufacturer(repo port.ManufacturerRepository) *ListManufacturer {
	return &ListManufacturer{repo: repo}
}

func (l *ListManufacturer) Execute(ctx context.Context, req ListManufacturerRequest) (*ListManufacturerResponse, error) {
	manufacturers, err := l.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return &ListManufacturerResponse{
		Data: manufacturers,
	}, nil
}
