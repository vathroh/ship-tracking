package repository

import (
	"context"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CourierRepository interface {
	GetActive(ctx context.Context) ([]domain.Courier, error)
	UpsertCouriers(ctx context.Context, couriers []domain.Courier) error
}

type courierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(db *gorm.DB) CourierRepository {
	return &courierRepository{db: db}
}

func (r *courierRepository) GetActive(ctx context.Context) ([]domain.Courier, error) {
	var couriers []domain.Courier
	if err := r.db.WithContext(ctx).Where("is_active = ?", true).Order("name ASC").Find(&couriers).Error; err != nil {
		return nil, err
	}
	return couriers, nil
}

func (r *courierRepository) UpsertCouriers(ctx context.Context, couriers []domain.Courier) error {
	if len(couriers) == 0 {
		return nil
	}
	
	// Upsert on 'code' conflict
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "is_active"}),
	}).Create(&couriers).Error
}
