package manufacturer

import (
	"context"
	"time"

	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type ManufacturerService struct {
	repo port.ManufacturerRepository
}

func NewManufacturerService(repo port.ManufacturerRepository) *ManufacturerService {
	return &ManufacturerService{repo: repo}
}

func (s *ManufacturerService) Create(ctx context.Context, req CreateManufacturerRequest) (CreateManufacturerResponse, error) {
	manufacturer := &domain.Manufacturer{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
	}

	err := s.repo.Create(ctx, manufacturer)
	if err != nil {
		return CreateManufacturerResponse{}, err
	}

	return CreateManufacturerResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: manufacturer,
	}, nil
}

func (s *ManufacturerService) GetByID(ctx context.Context, req GetManufacturerRequest) (GetManufacturerResponse, error) {
	r, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return GetManufacturerResponse{}, err
	}

	return GetManufacturerResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: r,
	}, nil
}

func (s *ManufacturerService) Update(ctx context.Context, req UpdateManufacturerRequest) (UpdateManufacturerResponse, error) {
	manufacturer := &domain.Manufacturer{
		ID:        req.ID,
		Name:      req.Name,
		UpdatedBy: req.UpdatedBy,
		UpdatedAt: time.Now(),
	}

	err := s.repo.Update(ctx, manufacturer)
	if err != nil {
		return UpdateManufacturerResponse{}, err
	}

	manufacturer, err = s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return UpdateManufacturerResponse{}, err
	}

	return UpdateManufacturerResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: manufacturer,
	}, nil
}

func (s *ManufacturerService) List(ctx context.Context, req ListManufacturerRequest) (ListManufacturerResponse, error) {
	var filter port.ManufacturerFilter

	if len(req.ID_In) > 0 {
		filter.ID_In = req.ID_In
	}

	if req.Name_iLike != "" {
		filter.Name_iLike = req.Name_iLike
	}
	filter.Limit = req.GetPageSize()
	filter.Offset = req.GetOffset()

	r, err := s.repo.List(ctx, filter)
	if err != nil {
		return ListManufacturerResponse{}, err
	}

	return ListManufacturerResponse{
		ListResponse: http.ListResponse{
			HttpResponse: http.HttpResponse{
				Error: http.ErrorSuccess,
			},
		},
		Data: ListManufacturerDataResponse{
			Page:     req.GetPage(),
			PageSize: req.GetPageSize(),
			Records:  r,
		},
	}, nil
}

func (s *ManufacturerService) Delete(ctx context.Context, req DeleteManufacturerRequest) (DeleteManufacturerResponse, error) {
	now := time.Now()
	manufacturer := &domain.Manufacturer{
		ID:        req.ID,
		DeletedBy: &req.DeletedBy,
		DeletedAt: &now,
	}
	err := s.repo.Delete(ctx, manufacturer)
	if err != nil {
		return DeleteManufacturerResponse{}, err
	}

	return DeleteManufacturerResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
	}, nil
}
