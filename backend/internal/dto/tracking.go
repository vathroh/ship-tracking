package dto

type TrackingRequest struct {
	Courier        string `query:"courier" validate:"required,alpha"`
	TrackingNumber string `query:"tracking_number" validate:"required,alphanum"`
}

type TrackingResponseDTO struct {
	Summary TrackingSummaryDTO   `json:"summary"`
	Detail  TrackingDetailDTO    `json:"detail"`
	History []TrackingHistoryDTO `json:"history"`
}

type TrackingSummaryDTO struct {
	AWB     string `json:"awb"`
	Courier string `json:"courier"`
	Service string `json:"service"`
	Status  string `json:"status"`
	Date    string `json:"date"`
	Desc    string `json:"desc"`
	Amount  string `json:"amount"`
	Weight  string `json:"weight"`
}

type TrackingDetailDTO struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Shipper     string `json:"shipper"`
	Receiver    string `json:"receiver"`
}

type TrackingHistoryDTO struct {
	Date     string `json:"date"`
	Desc     string `json:"desc"`
	Location string `json:"location"`
}

// BinderByte response structures for parsing internal client calls
type BinderByteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Summary struct {
			AWB     string `json:"awb"`
			Courier string `json:"courier"`
			Service string `json:"service"`
			Status  string `json:"status"`
			Date    string `json:"date"`
			Desc    string `json:"desc"`
			Amount  string `json:"amount"`
			Weight  string `json:"weight"`
		} `json:"summary"`
		Detail struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Shipper     string `json:"shipper"`
			Receiver    string `json:"receiver"`
		} `json:"detail"`
		History []struct {
			Date     string `json:"date"`
			Desc     string `json:"desc"`
			Location string `json:"location"`
		} `json:"history"`
	} `json:"data"`
}
