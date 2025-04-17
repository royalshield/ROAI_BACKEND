package models

import "time"

type Event struct {
	ID                uint   `gorm:"primaryKey"`
	Picture           string `gorm:"type:varchar(255)"`
	Title             string `gorm:"type:varchar(100);not null"`
	Description       string `gorm:"type:text"`
	Place             string `gorm:"type:varchar(100)"`
	AttendeesExpected int    `json:"attendees_expected"`
	Date              time.Time
	MembersLimit      int    `json:"members_limit"`
	Attendees         []User `gorm:"many2many:event_attendees"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
