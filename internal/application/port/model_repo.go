package port

import (
	"context"
	"time"

	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type ModelFilter struct {
	ID_In []uint32

	ManufacturerID_In []uint32

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

type ModelRepository interface {
	Create(ctx context.Context, model *domain.Model) error
	GetByID(ctx context.Context, id uint32) (*domain.Model, error)
	Update(ctx context.Context, model *domain.Model) error
	Delete(ctx context.Context, model *domain.Model) error
	List(ctx context.Context, query ModelFilter) ([]*domain.Model, error)
}
