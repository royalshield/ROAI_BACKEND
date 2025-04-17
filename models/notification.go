package models

import "time"

type Notification struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
