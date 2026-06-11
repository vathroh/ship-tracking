package handler

import (
	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/fathur/cek-ongkir-resi/backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TrackingHandler struct {
	svc      service.TrackingService
	validate *validator.Validate
}

func NewTrackingHandler(svc service.TrackingService, val *validator.Validate) *TrackingHandler {
	return &TrackingHandler{
		svc:      svc,
		validate: val,
	}
}

func (h *TrackingHandler) Track(c *fiber.Ctx) error {
	var req dto.TrackingRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Status:  "error",
			Message: "Invalid request parameters",
		})
	}

	if err := h.validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Status:  "error",
			Message: "Validation failed: " + err.Error(),
		})
	}

	resp, err := h.svc.TrackShipment(c.Context(), req)
	if err != nil {
		code := fiber.StatusInternalServerError
		msg := "Internal server error"

		switch err {
		case domain.ErrNotFound:
			code = fiber.StatusNotFound
			msg = err.Error()
		case domain.ErrRateLimitExceeded:
			code = fiber.StatusTooManyRequests
			msg = err.Error()
		case domain.ErrExternalAPI:
			code = fiber.StatusBadGateway
			msg = err.Error()
		}

		return c.Status(code).JSON(dto.StandardResponse{
			Status:  "error",
			Message: msg,
		})
	}

	return c.JSON(dto.StandardResponse{
		Status: "success",
		Data:   resp,
	})
}
