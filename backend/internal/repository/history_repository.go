package repository

import (
	"context"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	Save(ctx context.Context, history *domain.SearchHistory) error
	GetLatest(ctx context.Context, courier string, page, limit int) ([]domain.SearchHistory, int64, error)
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (r *historyRepository) Save(ctx context.Context, history *domain.SearchHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *historyRepository) GetLatest(ctx context.Context, courier string, page, limit int) ([]domain.SearchHistory, int64, error) {
	var histories []domain.SearchHistory
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.SearchHistory{})

	if courier != "" {
		query = query.Where("courier = ?", courier)
	}

	// Get total count for pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("searched_at DESC").Offset(offset).Limit(limit).Find(&histories).Error; err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}
