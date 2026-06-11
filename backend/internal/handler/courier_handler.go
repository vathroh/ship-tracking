package handler

import (
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/fathur/cek-ongkir-resi/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CourierHandler struct {
	svc service.CourierService
}

func NewCourierHandler(svc service.CourierService) *CourierHandler {
	return &CourierHandler{svc: svc}
}

func (h *CourierHandler) GetCouriers(c *fiber.Ctx) error {
	couriers, err := h.svc.GetActiveCouriers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.StandardResponse{
			Status:  "error",
			Message: "Failed to fetch couriers",
		})
	}

	return c.JSON(dto.StandardResponse{
		Status: "success",
		Data:   couriers,
	})
}

func (h *CourierHandler) SyncCouriers(c *fiber.Ctx) error {
	if err := h.svc.SyncFromExternal(c.Context()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.StandardResponse{
			Status:  "error",
			Message: "Failed to sync couriers from external API",
		})
	}

	return c.JSON(dto.StandardResponse{
		Status:  "success",
		Message: "Couriers synchronized successfully",
	})
}
