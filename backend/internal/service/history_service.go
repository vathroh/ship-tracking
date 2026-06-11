package service

import (
	"context"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/fathur/cek-ongkir-resi/backend/internal/repository"
)

type HistoryService interface {
	SaveSearch(ctx context.Context, trackingNumber, courier string) error
	GetLatestSearches(ctx context.Context, req dto.HistoryQueryRequest) (*dto.PaginatedResponse, error)
}

type historyService struct {
	repo repository.HistoryRepository
}

func NewHistoryService(repo repository.HistoryRepository) HistoryService {
	return &historyService{repo: repo}
}

func (s *historyService) SaveSearch(ctx context.Context, trackingNumber, courier string) error {
	history := &domain.SearchHistory{
		TrackingNumber: trackingNumber,
		Courier:        courier,
	}
	return s.repo.Save(ctx, history)
}

func (s *historyService) GetLatestSearches(ctx context.Context, req dto.HistoryQueryRequest) (*dto.PaginatedResponse, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}

	limit := req.Limit
	if limit < 1 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	histories, total, err := s.repo.GetLatest(ctx, req.Courier, page, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	var dtos []dto.HistoryResponse
	for _, h := range histories {
		dtos = append(dtos, dto.HistoryResponse{
			ID:             h.ID,
			TrackingNumber: h.TrackingNumber,
			Courier:        h.Courier,
			SearchedAt:     h.SearchedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	// ensure data is always an array
	if dtos == nil {
		dtos = []dto.HistoryResponse{}
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &dto.PaginatedResponse{
		Data:       dtos,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}
