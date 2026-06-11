package router

import (
	"time"

	"github.com/fathur/cek-ongkir-resi/backend/internal/handler"
	"github.com/fathur/cek-ongkir-resi/backend/internal/middleware"
	"github.com/fathur/cek-ongkir-resi/backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRoutes(app *fiber.App, healthHandler *handler.HealthHandler, trackingHandler *handler.TrackingHandler, historyHandler *handler.HistoryHandler, courierHandler *handler.CourierHandler, cacheSvc service.CacheService) {
	api := app.Group("/api/v1")

	// Health Check
	api.Get("/health", healthHandler.Check)

	// Rate Limiter for tracking
	trackingLimiter := limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "error",
				"message": "Rate limit exceeded, please try again later",
			})
		},
	})

	// Tracking with Limiter and 30 minute cache
	api.Get("/tracking", trackingLimiter, middleware.CacheMiddleware(cacheSvc, 30*time.Minute), trackingHandler.Track)

	// History
	api.Get("/history", historyHandler.GetHistory)

	// Couriers
	api.Get("/couriers", courierHandler.GetCouriers)
	api.Post("/couriers/sync", courierHandler.SyncCouriers)
}



