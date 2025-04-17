package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Wallet    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Role      string `gorm:"type:enum('user','admin');default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
