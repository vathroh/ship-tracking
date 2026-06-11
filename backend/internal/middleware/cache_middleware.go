package middleware

import (
	"fmt"
	"log/slog"
	"regexp"
	"time"

	"github.com/fathur/cek-ongkir-resi/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

var alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func CacheMiddleware(cacheSvc service.CacheService, duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		courier := c.Query("courier")
		trackingNumber := c.Query("tracking_number")

		// Only apply caching for requests that have tracking parameters
		if courier == "" || trackingNumber == "" {
			return c.Next()
		}

		// Sanitize input to prevent cache poisoning (allow only alphanumeric)
		if !alphanumericRegex.MatchString(courier) || !alphanumericRegex.MatchString(trackingNumber) {
			// If input is maliciously invalid, let handler catch it
			return c.Next()
		}

		key := fmt.Sprintf("tracking:%s:%s", courier, trackingNumber)

		// 1. Check cache
		cachedData, err := cacheSvc.Get(c.Context(), key)
		if err != nil {
			slog.Warn("Cache error", "error", err)
		} else if cachedData != "" {
			// Cache hit
			c.Set("Content-Type", "application/json")
			c.Set("X-Cache", "HIT")
			return c.SendString(cachedData)
		}

		// 2. Call handler (BinderByte mapping logic) if cache miss
		if err := c.Next(); err != nil {
			return err
		}

		// 3. Save cache if successful response
		if c.Response().StatusCode() == fiber.StatusOK {
			body := c.Response().Body()
			if len(body) > 0 {
				err := cacheSvc.Set(c.Context(), key, body, duration)
				if err != nil {
					slog.Warn("Failed to save cache", "error", err)
				}
				c.Set("X-Cache", "MISS")
			}
		}

		return nil
	}
}
