package postgresql

import (
	"context"

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
	err := r.db.Select("name").Create(&manufacturer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ManufacturerRepository) GetByID(ctx context.Context, id uint32) (*domain.Manufacturer, error) {
	var manufacturer domain.Manufacturer
	err := r.db.First(&manufacturer, id).Error
	if err != nil {
		return nil, err
	}
	return &manufacturer, nil
}

func (r *ManufacturerRepository) Update(ctx context.Context, manufacturer *domain.Manufacturer) error {
	err := r.db.Save(&manufacturer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ManufacturerRepository) Delete(ctx context.Context, id uint32) error {
	var manufacturer domain.Manufacturer
	err := r.db.First(&manufacturer, id).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&manufacturer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ManufacturerRepository) List(ctx context.Context) ([]*domain.Manufacturer, error) {
	var manufacturers []*domain.Manufacturer
	err := r.db.Find(&manufacturers).Error
	if err != nil {
		return nil, err
	}
	return manufacturers, nil
}
