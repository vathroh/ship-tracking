package domain

import "time"

type Courier struct {
	ID        uint      `gorm:"primaryKey"`
	Code      string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
