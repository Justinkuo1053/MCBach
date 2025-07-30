package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"` // 改為 varchar(255)
	Hash      string `gorm:"type:varchar(255);not null"`             // 密碼雜湊值也建議使用 varchar
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
}
