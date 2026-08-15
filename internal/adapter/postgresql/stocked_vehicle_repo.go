package postgresql

import (
	"context"

	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"gorm.io/gorm"
)

// agingDayExpr computes the number of whole days a vehicle has been in stock,
// based on created_at (the stock-in date).
const agingDayExpr = "(CURRENT_DATE - created_at::date)"

type StockedVehicleRepository struct {
	db *gorm.DB
}

func NewStockedVehicleRepository(db *gorm.DB) *StockedVehicleRepository {
	return &StockedVehicleRepository{db: db}
}

func (r *StockedVehicleRepository) Create(ctx context.Context, vehicle *domain.StockedVehicle) error {
	m := StockedVehicleModel{
		VIN:     vehicle.VIN,
		ModelID: vehicle.ModelID,
		Name:    vehicle.Name,
		Price:   float64(vehicle.Price),
		Action:  string(vehicle.Action),
		PostgresModel: PostgresModel{
			CreatedBy: vehicle.CreatedBy,
		},
	}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}

	vehicle.ID = m.ID
	vehicle.CreatedAt = m.CreatedAt
	vehicle.UpdatedAt = m.UpdatedAt
	return nil
}

func (r *StockedVehicleRepository) GetByID(ctx context.Context, id uint32) (*domain.StockedVehicle, error) {
	var vehicle domain.StockedVehicle
	err := r.db.WithContext(ctx).
		Select("stocked_vehicles.*, "+agingDayExpr+" AS aging_day").
		First(&vehicle, id).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *StockedVehicleRepository) Update(ctx context.Context, vehicle *domain.StockedVehicle) error {
	var m StockedVehicleModel
	err := r.db.WithContext(ctx).First(&m, vehicle.ID).Error
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Model(&m).Updates(StockedVehicleModel{
		VIN:     vehicle.VIN,
		ModelID: vehicle.ModelID,
		Name:    vehicle.Name,
		Price:   float64(vehicle.Price),
		Action:  string(vehicle.Action),
		PostgresModel: PostgresModel{
			UpdatedAt: vehicle.UpdatedAt,
			UpdatedBy: vehicle.UpdatedBy,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *StockedVehicleRepository) Delete(ctx context.Context, vehicle *domain.StockedVehicle) error {
	var m StockedVehicleModel
	err := r.db.WithContext(ctx).First(&m, vehicle.ID).Error
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Model(&m).Updates(StockedVehicleModel{
		PostgresModel: PostgresModel{
			DeletedAt: vehicle.DeletedAt,
			DeletedBy: vehicle.DeletedBy,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *StockedVehicleRepository) List(ctx context.Context, filter port.StockedVehicleFilter) ([]*domain.StockedVehicle, error) {
	var vehicles []*domain.StockedVehicle
	query := r.db.WithContext(ctx).
		Select("stocked_vehicles.*, " + agingDayExpr + " AS aging_day")

	if len(filter.ID_In) > 0 {
		query = query.Where("id IN ?", filter.ID_In)
	}

	if len(filter.ModelID_In) > 0 {
		query = query.Where("model_id IN ?", filter.ModelID_In)
	}

	if filter.VIN != "" {
		query = query.Where("vin = ?", filter.VIN)
	}

	if filter.Name_iLike != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name_iLike+"%")
	}

	if len(filter.Action_In) > 0 {
		query = query.Where("action IN ?", filter.Action_In)
	}

	if filter.AgingDay_Gte != nil {
		query = query.Where(agingDayExpr+" >= ?", *filter.AgingDay_Gte)
	}
	if filter.AgingDay_Lte != nil {
		query = query.Where(agingDayExpr+" <= ?", *filter.AgingDay_Lte)
	}

	if filter.CreatedAt_Gte != nil {
		query = query.Where("created_at >= ?", filter.CreatedAt_Gte)
	}
	if filter.CreatedAt_Lte != nil {
		query = query.Where("created_at <= ?", filter.CreatedAt_Lte)
	}
	if filter.CreatedBy != nil {
		query = query.Where("created_by = ?", *filter.CreatedBy)
	}

	if filter.UpdatedAt_Gte != nil {
		query = query.Where("updated_at >= ?", filter.UpdatedAt_Gte)
	}
	if filter.UpdatedAt_Lte != nil {
		query = query.Where("updated_at <= ?", filter.UpdatedAt_Lte)
	}
	if filter.UpdatedBy != nil {
		query = query.Where("updated_by = ?", *filter.UpdatedBy)
	}

	if filter.DeletedAt_Gte != nil {
		query = query.Where("deleted_at >= ?", filter.DeletedAt_Gte)
	}
	if filter.DeletedAt_Lte != nil {
		query = query.Where("deleted_at <= ?", filter.DeletedAt_Lte)
	}
	if filter.DeletedBy != nil {
		query = query.Where("deleted_by = ?", *filter.DeletedBy)
	}

	err := query.Limit(filter.Limit).
		Offset(filter.Offset).
		Find(&vehicles).Error
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}
