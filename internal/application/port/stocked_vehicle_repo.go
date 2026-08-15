package port

import (
	"context"
	"time"

	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type StockedVehicleFilter struct {
	ID_In []uint32

	ModelID_In []uint32

	VIN string

	Name_iLike string

	Action_In []domain.VehicleAction

	AgingDay_Gte *int
	AgingDay_Lte *int

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

type StockedVehicleRepository interface {
	Create(ctx context.Context, vehicle *domain.StockedVehicle) error
	GetByID(ctx context.Context, id uint32) (*domain.StockedVehicle, error)
	Update(ctx context.Context, vehicle *domain.StockedVehicle) error
	Delete(ctx context.Context, vehicle *domain.StockedVehicle) error
	List(ctx context.Context, query StockedVehicleFilter) ([]*domain.StockedVehicle, error)
}
