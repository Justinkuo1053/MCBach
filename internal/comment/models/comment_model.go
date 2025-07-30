package models

import "time"

type Comment struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Content       string
	IsDeleted     bool   `gorm:"default:false"`
	UserID        uint   `gorm:"not null"`
	AlbumID       uint   `gorm:"not null"`
	LikesCount    int    `gorm:"default:0"` // 一般讚數量
	ProLikesCount int    `gorm:"default:0"` // 專業讚數量
	Likes         []Like `gorm:"foreignKey:CommentID"`
	IsEdited      bool   `gorm:"default:false"` // 是否已編輯
}

type Like struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint `gorm:"primaryKey"`
	CommentID uint `gorm:"primaryKey"`
	IsPro     bool `gorm:"default:false"` // 是否為專業讚
}
