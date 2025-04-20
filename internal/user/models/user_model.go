package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"uniqueIndex;not null"`
	Hash      string `gorm:"not null"`
	FirstName string
	LastName  string
}

