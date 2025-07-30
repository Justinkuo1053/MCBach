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

// 定義 IsMember 方法
func (u *User) IsMember() bool {
	// 實現邏輯，例如檢查某些條件
	return true // 或根據業務邏輯返回適當的值
}
