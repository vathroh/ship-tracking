package middleware

import (
	"log/slog"

	"github.com/fathur/cek-ongkir-resi/backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	if code >= 500 {
		slog.Error("Server error", "path", c.Path(), "error", err.Error())
	}

	return c.Status(code).JSON(dto.StandardResponse{
		Status:  "error",
		Message: message,
	})
}
