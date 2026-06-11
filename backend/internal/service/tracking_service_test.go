package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
)

type mockHistoryService struct{}

func (m *mockHistoryService) SaveSearch(ctx context.Context, trackingNumber, courier string) error {
	return nil
}

func (m *mockHistoryService) GetLatestSearches(ctx context.Context, req dto.HistoryQueryRequest) (*dto.PaginatedResponse, error) {
	return nil, nil
}

func TestTrackShipment_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": 200,
			"message": "Successfully tracked package",
			"data": {
				"summary": {"awb": "JP123456789", "courier": "jnt", "status": "DELIVERED"},
				"detail": {"origin": "Jakarta", "destination": "Bandung"},
				"history": [
					{"date": "2023-10-01", "desc": "Terkirim", "location": "Bandung"}
				]
			}
		}`))
	}))
	defer server.Close()

	svc := NewTrackingService(server.URL, "test-key", &mockHistoryService{})
	req := dto.TrackingRequest{Courier: "jnt", TrackingNumber: "JP123456789"}
	
	resp, err := svc.TrackShipment(context.Background(), req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Summary.AWB != "JP123456789" {
		t.Errorf("expected AWB JP123456789, got %s", resp.Summary.AWB)
	}
	if len(resp.History) != 1 {
		t.Errorf("expected 1 history item, got %d", len(resp.History))
	}
}

func TestTrackShipment_RateLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	svc := NewTrackingService(server.URL, "test-key", &mockHistoryService{})
	req := dto.TrackingRequest{Courier: "jnt", TrackingNumber: "JP123456789"}

	_, err := svc.TrackShipment(context.Background(), req)

	if err != domain.ErrRateLimitExceeded {
		t.Fatalf("expected ErrRateLimitExceeded, got %v", err)
	}
}

func TestTrackShipment_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": 400, "message": "Data tidak ditemukan"}`))
	}))
	defer server.Close()

	svc := NewTrackingService(server.URL, "test-key", &mockHistoryService{})
	req := dto.TrackingRequest{Courier: "jnt", TrackingNumber: "UNKNOWN"}

	_, err := svc.TrackShipment(context.Background(), req)

	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
