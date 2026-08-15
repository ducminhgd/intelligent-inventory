package model

import (
	"context"
	"time"

	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type ModelService struct {
	repo port.ModelRepository
}

func NewModelService(repo port.ModelRepository) *ModelService {
	return &ModelService{repo: repo}
}

func (s *ModelService) Create(ctx context.Context, req CreateModelRequest) (CreateModelResponse, error) {
	model := &domain.Model{
		ManufacturerID: req.ManufacturerID,
		Name:           req.Name,
		CreatedBy:      req.CreatedBy,
	}

	err := s.repo.Create(ctx, model)
	if err != nil {
		return CreateModelResponse{}, err
	}

	return CreateModelResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: model,
	}, nil
}

func (s *ModelService) GetByID(ctx context.Context, req GetModelRequest) (GetModelResponse, error) {
	r, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return GetModelResponse{}, err
	}

	return GetModelResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: r,
	}, nil
}

func (s *ModelService) Update(ctx context.Context, req UpdateModelRequest) (UpdateModelResponse, error) {
	model := &domain.Model{
		ID:             req.ID,
		ManufacturerID: req.ManufacturerID,
		Name:           req.Name,
		UpdatedBy:      req.UpdatedBy,
		UpdatedAt:      time.Now(),
	}

	err := s.repo.Update(ctx, model)
	if err != nil {
		return UpdateModelResponse{}, err
	}

	model, err = s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return UpdateModelResponse{}, err
	}

	return UpdateModelResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
		Data: model,
	}, nil
}

func (s *ModelService) List(ctx context.Context, req ListModelRequest) (ListModelResponse, error) {
	var filter port.ModelFilter

	if len(req.ID_In) > 0 {
		filter.ID_In = req.ID_In
	}

	if len(req.ManufacturerID_In) > 0 {
		filter.ManufacturerID_In = req.ManufacturerID_In
	}

	if req.Name_iLike != "" {
		filter.Name_iLike = req.Name_iLike
	}
	filter.Limit = req.GetPageSize()
	filter.Offset = req.GetOffset()

	r, err := s.repo.List(ctx, filter)
	if err != nil {
		return ListModelResponse{}, err
	}

	return ListModelResponse{
		ListResponse: http.ListResponse{
			HttpResponse: http.HttpResponse{
				Error: http.ErrorSuccess,
			},
		},
		Data: ListModelDataResponse{
			Page:     req.GetPage(),
			PageSize: req.GetPageSize(),
			Records:  r,
		},
	}, nil
}

func (s *ModelService) Delete(ctx context.Context, req DeleteModelRequest) (DeleteModelResponse, error) {
	now := time.Now()
	model := &domain.Model{
		ID:        req.ID,
		DeletedBy: &req.DeletedBy,
		DeletedAt: &now,
	}
	err := s.repo.Delete(ctx, model)
	if err != nil {
		return DeleteModelResponse{}, err
	}

	return DeleteModelResponse{
		HttpResponse: http.HttpResponse{
			Error: http.ErrorSuccess,
		},
	}, nil
}
