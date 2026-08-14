package port

import (
	"context"

	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type ManufacturerRepository interface {
	Create(ctx context.Context, manufacturer *domain.Manufacturer) error
	GetByID(ctx context.Context, id uint32) (*domain.Manufacturer, error)
	Update(ctx context.Context, manufacturer *domain.Manufacturer) error
	Delete(ctx context.Context, id uint32) error
	List(ctx context.Context) ([]*domain.Manufacturer, error)
}
