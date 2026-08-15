package model

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"github.com/google/uuid"
)

type UpdateModelRequest struct {
	ID             uint32 `json:"id"`
	ManufacturerID uint32 `json:"manufacturer_id"`
	Name           string `json:"name"`

	UpdatedBy uuid.UUID `json:"updated_by"`
}

type UpdateModelResponse struct {
	http.HttpResponse
	Data *domain.Model `json:"data"`
}

type DeleteModelRequest struct {
	ID        uint32    `json:"id"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}

type DeleteModelResponse struct {
	http.HttpResponse
}
