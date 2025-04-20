package models

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	IsDeleted bool   `gorm:"default:false"`
	UserID    uint   `gorm:"not null"`
	AlbumID   uint   `gorm:"not null"`
	Likes     []Like `gorm:"foreignKey:CommentID"`
}

type Like struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint `gorm:"primaryKey"`
	CommentID uint `gorm:"primaryKey"`
}
