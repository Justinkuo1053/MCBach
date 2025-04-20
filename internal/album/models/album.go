package models

import (
	"time"
)

type Album struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string    `gorm:"not null"`
	AlbumType      string    `gorm:"not null;default:'ALBUM'"`
	ReleaseAt      time.Time `gorm:"not null"`
	SpotifyAlbumId *string   `gorm:"type:varchar(255);uniqueIndex"` // 修改為 varchar(255)
}
