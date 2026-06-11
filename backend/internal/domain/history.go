package domain

import "time"

// SearchHistory represents a search request record in the database.
type SearchHistory struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TrackingNumber string    `gorm:"index;not null" json:"tracking_number"`
	Courier        string    `gorm:"index;not null" json:"courier"`
	SearchedAt     time.Time `gorm:"autoCreateTime" json:"searched_at"`
}
