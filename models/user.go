package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Wallet    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Role      string `gorm:"type:enum('user','admin');default:'user'"`
	Username  string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Balance   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
