package service

import (
	"context"
	"testing"
	"time"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
)

type mockHistoryRepository struct {
	histories []domain.SearchHistory
	saved     bool
}

func (m *mockHistoryRepository) Save(ctx context.Context, history *domain.SearchHistory) error {
	m.saved = true
	m.histories = append(m.histories, *history)
	return nil
}

func (m *mockHistoryRepository) GetLatest(ctx context.Context, courier string, page, limit int) ([]domain.SearchHistory, int64, error) {
	if courier != "" {
		var filtered []domain.SearchHistory
		for _, h := range m.histories {
			if h.Courier == courier {
				filtered = append(filtered, h)
			}
		}
		return filtered, int64(len(filtered)), nil
	}
	return m.histories, int64(len(m.histories)), nil
}

func TestHistoryService_SaveSearch(t *testing.T) {
	repo := &mockHistoryRepository{}
	svc := NewHistoryService(repo)

	err := svc.SaveSearch(context.Background(), "JP123", "jnt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.saved {
		t.Errorf("expected Save to be called")
	}

	if repo.histories[0].TrackingNumber != "JP123" {
		t.Errorf("expected tracking number to be saved")
	}
}

func TestHistoryService_GetLatestSearches(t *testing.T) {
	repo := &mockHistoryRepository{
		histories: []domain.SearchHistory{
			{ID: 1, TrackingNumber: "JP1", Courier: "jnt", SearchedAt: time.Now()},
			{ID: 2, TrackingNumber: "JP2", Courier: "jne", SearchedAt: time.Now()},
		},
	}
	svc := NewHistoryService(repo)

	// Test without filter
	req := dto.HistoryQueryRequest{Page: 1, Limit: 10}
	resp, err := svc.GetLatestSearches(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Total != 2 {
		t.Errorf("expected total 2, got %d", resp.Total)
	}

	// Test with courier filter
	reqFilter := dto.HistoryQueryRequest{Page: 1, Limit: 10, Courier: "jnt"}
	respFilter, err := svc.GetLatestSearches(context.Background(), reqFilter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if respFilter.Total != 1 {
		t.Errorf("expected total 1, got %d", respFilter.Total)
	}
}
