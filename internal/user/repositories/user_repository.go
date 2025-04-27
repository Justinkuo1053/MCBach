package repositories

import (
	"mcbach/internal/user/models" // 修改 import 路徑

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(userID uint, updates map[string]interface{}) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&user).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	// 實現邏輯，例如從資料庫獲取數據
	return nil, nil
}
