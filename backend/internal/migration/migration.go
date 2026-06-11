package migration

import (
	"github.com/fathur/cek-ongkir-resi/backend/internal/domain"
	"gorm.io/gorm"
)

// Run executes the database auto-migrations
func Run(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.SearchHistory{},
		&domain.Courier{},
	)
	return err
}
