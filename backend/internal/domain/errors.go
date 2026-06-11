package domain

import "errors"

var (
	ErrNotFound          = errors.New("tracking data not found")
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrExternalAPI       = errors.New("external api error")
	ErrInvalidInput      = errors.New("invalid input")
	ErrInternal          = errors.New("internal server error")
)
