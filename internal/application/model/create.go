package model

import (
	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"github.com/google/uuid"
)

type CreateModelRequest struct {
	ManufacturerID uint32 `json:"manufacturer_id"`
	Name           string `json:"name"`

	CreatedBy uuid.UUID `json:"created_by"`
}

type CreateModelResponse struct {
	http.HttpResponse

	Data *domain.Model `json:"data"`
}
