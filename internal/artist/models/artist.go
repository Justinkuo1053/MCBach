package models

import (
	"time"
)

// 更新 Artist 模型，加入多對多關聯
type Artist struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string `gorm:"not null"`
	SpotifyArtistId string `gorm:"type:varchar(255);uniqueIndex;not null"` // 修改為 varchar(255)

}
