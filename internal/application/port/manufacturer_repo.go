package port

import (
	"context"
	"time"

	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type ManufacturerFilter struct {
	ID_In []uint32

	Name_iLike string

	CreatedAt_Gte *time.Time
	CreatedAt_Lte *time.Time
	CreatedBy     *string

	UpdatedAt_Gte *time.Time
	UpdatedAt_Lte *time.Time
	UpdatedBy     *string

	DeletedAt_Gte *time.Time
	DeletedAt_Lte *time.Time
	DeletedBy     *string

	Limit  int
	Offset int
}

type ManufacturerRepository interface {
	Create(ctx context.Context, manufacturer *domain.Manufacturer) error
	GetByID(ctx context.Context, id uint32) (*domain.Manufacturer, error)
	Update(ctx context.Context, manufacturer *domain.Manufacturer) error
	Delete(ctx context.Context, manufacturer *domain.Manufacturer) error
	List(ctx context.Context, query ManufacturerFilter) ([]*domain.Manufacturer, error)
}
