package dto

type CourierResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type BinderByteCourierListResponse []struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
