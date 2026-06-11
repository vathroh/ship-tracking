package handler

import (
	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/fathur/cek-ongkir-resi/backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	svc      service.HistoryService
	validate *validator.Validate
}

func NewHistoryHandler(svc service.HistoryService, val *validator.Validate) *HistoryHandler {
	return &HistoryHandler{
		svc:      svc,
		validate: val,
	}
}

func (h *HistoryHandler) GetHistory(c *fiber.Ctx) error {
	var req dto.HistoryQueryRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Status:  "error",
			Message: "Invalid query parameters",
		})
	}

	if err := h.validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Status:  "error",
			Message: "Validation failed: " + err.Error(),
		})
	}

	resp, err := h.svc.GetLatestSearches(c.Context(), req)
	if err != nil {
		code := fiber.StatusInternalServerError
		if err == domain.ErrInvalidInput {
			code = fiber.StatusBadRequest
		}

		return c.Status(code).JSON(dto.StandardResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(dto.StandardResponse{
		Status: "success",
		Data:   resp,
	})
}
