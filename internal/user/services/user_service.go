package services

import (
	"mcbach/internal/user/models"       // 修改 import 路徑
	"mcbach/internal/user/repositories" // 修改 import 路徑
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.repo.FindByID(userID)
}

func (s *UserService) EditUser(userID uint, updates map[string]interface{}) (*models.User, error) {
	return s.repo.Update(userID, updates)
}
