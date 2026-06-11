package handler

import (
	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.JSON(dto.StandardResponse{
		Status:  "success",
		Message: "Service is healthy",
	})
}
