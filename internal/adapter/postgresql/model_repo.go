package postgresql

import (
	"context"

	"github.com/ducminhgd/intelligent-inventory/internal/application/port"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
	"gorm.io/gorm"
)

type ModelRepository struct {
	db *gorm.DB
}

func NewModelRepository(db *gorm.DB) *ModelRepository {
	return &ModelRepository{db: db}
}

func (r *ModelRepository) Create(ctx context.Context, model *domain.Model) error {
	m := ModelModel{
		ManufacturerID: model.ManufacturerID,
		Name:           model.Name,
		PostgresModel: PostgresModel{
			CreatedBy: model.CreatedBy,
		},
	}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}

	model.ID = m.ID
	model.CreatedAt = m.CreatedAt
	model.UpdatedAt = m.UpdatedAt
	return nil
}

func (r *ModelRepository) GetByID(ctx context.Context, id uint32) (*domain.Model, error) {
	var model domain.Model
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ModelRepository) Update(ctx context.Context, model *domain.Model) error {
	var m ModelModel
	err := r.db.WithContext(ctx).First(&m, model.ID).Error
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Model(&m).Updates(ModelModel{
		ManufacturerID: model.ManufacturerID,
		Name:           model.Name,
		PostgresModel: PostgresModel{
			UpdatedAt: model.UpdatedAt,
			UpdatedBy: model.UpdatedBy,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ModelRepository) Delete(ctx context.Context, model *domain.Model) error {
	var m ModelModel
	err := r.db.WithContext(ctx).First(&m, model.ID).Error
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Model(&m).Updates(ModelModel{
		PostgresModel: PostgresModel{
			DeletedAt: model.DeletedAt,
			DeletedBy: model.DeletedBy,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ModelRepository) List(ctx context.Context, filter port.ModelFilter) ([]*domain.Model, error) {
	var models []*domain.Model
	query := r.db.WithContext(ctx)

	if len(filter.ID_In) > 0 {
		query = query.Where("id IN ?", filter.ID_In)
	}

	if len(filter.ManufacturerID_In) > 0 {
		query = query.Where("manufacturer_id IN ?", filter.ManufacturerID_In)
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
		Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}
