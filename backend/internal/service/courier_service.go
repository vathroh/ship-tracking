package service

import (
	"context"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/fathur/cek-ongkir-resi/backend/internal/repository"
)

type CourierService interface {
	GetActiveCouriers(ctx context.Context) ([]dto.CourierResponse, error)
	SyncFromExternal(ctx context.Context) error
}

type courierService struct {
	repo    repository.CourierRepository
	client  *http.Client
	baseURL string
	apiKey  string
}

func NewCourierService(repo repository.CourierRepository, baseURL, apiKey string) CourierService {
	return &courierService{
		repo:    repo,
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (s *courierService) GetActiveCouriers(ctx context.Context) ([]dto.CourierResponse, error) {
	couriers, err := s.repo.GetActive(ctx)
	if err != nil {
		return nil, domain.ErrInternal
	}

	var dtos []dto.CourierResponse
	for _, c := range couriers {
		dtos = append(dtos, dto.CourierResponse{
			Code: c.Code,
			Name: c.Name,
		})
	}

	if dtos == nil {
		dtos = []dto.CourierResponse{}
	}

	return dtos, nil
}

func (s *courierService) SyncFromExternal(ctx context.Context) error {
	targetURL := fmt.Sprintf("%s/list_courier?api_key=%s", s.baseURL, s.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("external api returned status: %d", resp.StatusCode)
	}

	var bbResp dto.BinderByteCourierListResponse
	if err := json.NewDecoder(resp.Body).Decode(&bbResp); err != nil {
		return err
	}

	if len(bbResp) == 0 {
		return fmt.Errorf("external api returned an empty list of couriers")
	}

	var domainCouriers []domain.Courier
	for _, c := range bbResp {
		domainCouriers = append(domainCouriers, domain.Courier{
			Code:     c.Code,
			Name:     c.Description,
			IsActive: true,
		})
	}

	return s.repo.UpsertCouriers(ctx, domainCouriers)
}
