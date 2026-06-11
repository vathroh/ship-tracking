package dto

type HistoryQueryRequest struct {
	Page    int    `query:"page" validate:"omitempty,min=1"`
	Limit   int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Courier string `query:"courier" validate:"omitempty,alpha"`
}

type HistoryResponse struct {
	ID             uint   `json:"id"`
	TrackingNumber string `json:"tracking_number"`
	Courier        string `json:"courier"`
	SearchedAt     string `json:"searched_at"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
}
