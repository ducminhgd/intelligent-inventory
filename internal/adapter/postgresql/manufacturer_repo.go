package postgresql

import (
	"context"

	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"gorm.io/gorm"
)

type ManufacturerRepository struct {
	db *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{db: db}
}

func (r *ManufacturerRepository) Create(ctx context.Context, manufacturer *domain.Manufacturer) error {
	model := ManufacturerModel{
		Name: manufacturer.Name,
		PostgresModel: PostgresModel{
			CreatedBy: manufacturer.CreatedBy,
		},
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	manufacturer.ID = model.ID
	manufacturer.CreatedAt = model.CreatedAt
	manufacturer.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *ManufacturerRepository) GetByID(ctx context.Context, id uint32) (*domain.Manufacturer, error) {
	var manufacturer domain.Manufacturer
	err := r.db.WithContext(ctx).First(&manufacturer, id).Error
	if err != nil {
		return nil, err
	}
	return &manufacturer, nil
}

func (r *ManufacturerRepository) Update(ctx context.Context, manufacturer *domain.Manufacturer) error {
	var m ManufacturerModel
	err := r.db.WithContext(ctx).First(&m, manufacturer.ID).Error
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Model(&m).Updates(ManufacturerModel{
		Name: manufacturer.Name,
		PostgresModel: PostgresModel{
			UpdatedAt: manufacturer.UpdatedAt,
			UpdatedBy: manufacturer.UpdatedBy,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ManufacturerRepository) Delete(ctx context.Context, manufacturer *domain.Manufacturer) error {
	var m ManufacturerModel
	err := r.db.WithContext(ctx).First(&m, manufacturer.ID).Error
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Model(&m).Updates(ManufacturerModel{
		Name: manufacturer.Name,
		PostgresModel: PostgresModel{
			DeletedAt: manufacturer.DeletedAt,
			DeletedBy: manufacturer.DeletedBy,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ManufacturerRepository) List(ctx context.Context, filter port.ManufacturerFilter) ([]*domain.Manufacturer, error) {
	var manufacturers []*domain.Manufacturer
	query := r.db.WithContext(ctx)

	if len(filter.ID_In) > 0 {
		query = query.Where("id IN ?", filter.ID_In)
	}

	if filter.Name_iLike != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name_iLike+"%")
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
		Find(&manufacturers).Error
	if err != nil {
		return nil, err
	}
	return manufacturers, nil
}
