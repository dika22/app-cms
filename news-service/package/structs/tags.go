package structs

import "time"

type Tag struct {
	ID            int64       `gorm:"primaryKey"`
	Name          string      `gorm:"uniqueIndex;not null"`
	UsageCount    int64       `gorm:"default:0"` // across published versions
	TrendingScore float64     `gorm:"default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}