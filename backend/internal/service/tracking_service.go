package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
)

type TrackingService interface {
	TrackShipment(ctx context.Context, req dto.TrackingRequest) (*dto.TrackingResponseDTO, error)
}

type trackingService struct {
	client      *http.Client
	baseURL     string
	apiKey      string
	historySvc  HistoryService
}

func NewTrackingService(baseURL, apiKey string, historySvc HistoryService) TrackingService {
	return &trackingService{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseURL:     baseURL,
		apiKey:      apiKey,
		historySvc:  historySvc,
	}
}

func (s *trackingService) TrackShipment(ctx context.Context, req dto.TrackingRequest) (*dto.TrackingResponseDTO, error) {
	// Secure URL encoding to prevent Parameter Injection
	q := url.Values{}
	q.Add("api_key", s.apiKey)
	q.Add("courier", req.Courier)
	q.Add("awb", req.TrackingNumber)

	targetURL := fmt.Sprintf("%s/track?%s", s.baseURL, q.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		// Scrub explicit error to prevent URL/API key leakage
		return nil, domain.ErrExternalAPI
	}

	resp, err := s.client.Do(httpReq)
	if err != nil {
		// Scrub explicit error to prevent URL/API key leakage from network timeouts
		return nil, domain.ErrExternalAPI
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, domain.ErrRateLimitExceeded
	}

	if resp.StatusCode != http.StatusOK {
		var bbResp dto.BinderByteResponse
		if err := json.NewDecoder(resp.Body).Decode(&bbResp); err == nil {
			if bbResp.Status == 400 {
				return nil, domain.ErrNotFound
			}
		}
		return nil, domain.ErrExternalAPI
	}

	var bbResp dto.BinderByteResponse
	if err := json.NewDecoder(resp.Body).Decode(&bbResp); err != nil {
		return nil, domain.ErrExternalAPI
	}

	if bbResp.Status != 200 {
		return nil, domain.ErrNotFound
	}

	// Sanitize and map to internal DTO
	result := &dto.TrackingResponseDTO{
		Summary: dto.TrackingSummaryDTO{
			AWB:     bbResp.Data.Summary.AWB,
			Courier: bbResp.Data.Summary.Courier,
			Service: bbResp.Data.Summary.Service,
			Status:  bbResp.Data.Summary.Status,
			Date:    bbResp.Data.Summary.Date,
			Desc:    bbResp.Data.Summary.Desc,
			Amount:  bbResp.Data.Summary.Amount,
			Weight:  bbResp.Data.Summary.Weight,
		},
		Detail: dto.TrackingDetailDTO{
			Origin:      bbResp.Data.Detail.Origin,
			Destination: bbResp.Data.Detail.Destination,
			Shipper:     maskString(bbResp.Data.Detail.Shipper),
			Receiver:    maskString(bbResp.Data.Detail.Receiver),
		},
		History: make([]dto.TrackingHistoryDTO, 0, len(bbResp.Data.History)),
	}

	for _, h := range bbResp.Data.History {
		result.History = append(result.History, dto.TrackingHistoryDTO{
			Date:     h.Date,
			Desc:     h.Desc,
			Location: h.Location,
		})
	}

	// Save tracking request history asynchronously
	go func() {
		_ = s.historySvc.SaveSearch(context.Background(), req.TrackingNumber, req.Courier)
	}()

	return result, nil
}

// maskString replaces characters after the first letter of each word with asterisks
func maskString(s string) string {
	if s == "" {
		return s
	}
	words := strings.Split(s, " ")
	for i, w := range words {
		if len(w) > 1 {
			words[i] = string(w[0]) + strings.Repeat("*", len(w)-1)
		}
	}
	return strings.Join(words, " ")
}
