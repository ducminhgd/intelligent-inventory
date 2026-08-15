package stockedvehicle

import (
	"context"
	"time"

	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type StockedVehicleService struct {
	repo port.StockedVehicleRepository
}

func NewStockedVehicleService(repo port.StockedVehicleRepository) *StockedVehicleService {
	return &StockedVehicleService{repo: repo}
}

func (s *StockedVehicleService) Create(ctx context.Context, req CreateStockedVehicleRequest) (CreateStockedVehicleResponse, error) {
	action := req.Action
	if action == "" {
		action = domain.ActionNone
	}

	vehicle := &domain.StockedVehicle{
		VIN:       req.VIN,
		ModelID:   req.ModelID,
		Name:      req.Name,
		Price:     req.Price,
		Action:    action,
		CreatedBy: req.CreatedBy,
	}

	err := s.repo.Create(ctx, vehicle)
	if err != nil {
		return CreateStockedVehicleResponse{}, err
	}

	return CreateStockedVehicleResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: vehicle,
	}, nil
}

func (s *StockedVehicleService) GetByID(ctx context.Context, req GetStockedVehicleRequest) (GetStockedVehicleResponse, error) {
	r, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return GetStockedVehicleResponse{}, err
	}

	return GetStockedVehicleResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: r,
	}, nil
}

func (s *StockedVehicleService) Update(ctx context.Context, req UpdateStockedVehicleRequest) (UpdateStockedVehicleResponse, error) {
	vehicle := &domain.StockedVehicle{
		ID:        req.ID,
		VIN:       req.VIN,
		ModelID:   req.ModelID,
		Name:      req.Name,
		Price:     req.Price,
		Action:    req.Action,
		UpdatedBy: req.UpdatedBy,
		UpdatedAt: time.Now(),
	}

	err := s.repo.Update(ctx, vehicle)
	if err != nil {
		return UpdateStockedVehicleResponse{}, err
	}

	vehicle, err = s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return UpdateStockedVehicleResponse{}, err
	}

	return UpdateStockedVehicleResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: vehicle,
	}, nil
}

func (s *StockedVehicleService) List(ctx context.Context, req ListStockedVehicleRequest) (ListStockedVehicleResponse, error) {
	var filter port.StockedVehicleFilter

	if len(req.ID_In) > 0 {
		filter.ID_In = req.ID_In
	}

	if len(req.ModelID_In) > 0 {
		filter.ModelID_In = req.ModelID_In
	}

	if req.VIN != "" {
		filter.VIN = req.VIN
	}

	if req.Name_iLike != "" {
		filter.Name_iLike = req.Name_iLike
	}
	filter.Limit = req.GetPageSize()
	filter.Offset = req.GetOffset()

	r, err := s.repo.List(ctx, filter)
	if err != nil {
		return ListStockedVehicleResponse{}, err
	}

	return ListStockedVehicleResponse{
		ListResponse: http.ListResponse{
			HttpResponse: http.HttpResponse{
				Error: http.ErrorSuccess,
			},
		},
		Data: ListStockedVehicleDataResponse{
			Page:     req.GetPage(),
			PageSize: req.GetPageSize(),
			Records:  r,
		},
	}, nil
}

func (s *StockedVehicleService) Delete(ctx context.Context, req DeleteStockedVehicleRequest) (DeleteStockedVehicleResponse, error) {
	now := time.Now()
	vehicle := &domain.StockedVehicle{
		ID:        req.ID,
		DeletedBy: &req.DeletedBy,
		DeletedAt: &now,
	}
	err := s.repo.Delete(ctx, vehicle)
	if err != nil {
		return DeleteStockedVehicleResponse{}, err
	}

	return DeleteStockedVehicleResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
	}, nil
}
