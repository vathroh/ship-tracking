package main

import (
	"context"
	"log/slog"

	"github.com/fathur/cek-ongkir-resi/backend/internal/handler"
	"github.com/fathur/cek-ongkir-resi/backend/internal/middleware"
	"github.com/fathur/cek-ongkir-resi/backend/internal/migration"
	"github.com/fathur/cek-ongkir-resi/backend/internal/repository"
	"github.com/fathur/cek-ongkir-resi/backend/internal/router"
	"github.com/fathur/cek-ongkir-resi/backend/internal/service"
	"github.com/fathur/cek-ongkir-resi/backend/pkg/config"
	"github.com/fathur/cek-ongkir-resi/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func main() {
	// 1. Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	// 2. Setup logger
	logger.Setup(cfg.LogLevel)

	// Setup Validator
	val := validator.New()

	// 3. Initialize Database & Repositories
	dsn := cfg.PostgresDSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Run Migrations
	if err := migration.Run(db); err != nil {
		panic("Failed to run migrations: " + err.Error())
	}

	historyRepo := repository.NewHistoryRepository(db)
	courierRepo := repository.NewCourierRepository(db)

	// Initialize Redis
	redisAddr := cfg.RedisAddr
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// 4. Initialize Services
	binderByteURL := cfg.BinderByteURL
	binderByteKey := cfg.BinderByteKey

	historySvc := service.NewHistoryService(historyRepo)
	trackingSvc := service.NewTrackingService(binderByteURL, binderByteKey, historySvc)
	cacheSvc := service.NewCacheService(redisClient)
	courierSvc := service.NewCourierService(courierRepo, binderByteURL, binderByteKey)

	// Trigger initial Courier Sync
	go func() {
		if err := courierSvc.SyncFromExternal(context.Background()); err != nil {
			slog.Error("Failed to sync couriers from external API on startup", "error", err)
		} else {
			slog.Info("Successfully synced couriers from external API")
		}
	}()

	// 5. Initialize Handlers
	healthHandler := handler.NewHealthHandler()
	trackingHandler := handler.NewTrackingHandler(trackingSvc, val)
	historyHandler := handler.NewHistoryHandler(historySvc, val)
	courierHandler := handler.NewCourierHandler(courierSvc)

	// 6. Setup Fiber App
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// 7. Global Middleware
	app.Use(recover.New())
	app.Use(fiberLogger.New())
	
	frontendURL := cfg.FrontendURL
	app.Use(cors.New(cors.Config{
		AllowOrigins: frontendURL,
	}))

	// 8. Register Routes
	router.SetupRoutes(app, healthHandler, trackingHandler, historyHandler, courierHandler, cacheSvc)

	// 9. Start Server
	slog.Info("Starting server", "port", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		slog.Error("Server failed", "error", err)
	}
}

