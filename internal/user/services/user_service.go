package services

import (
	"errors"
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

func (s *UserService) IsValidMember(userID uint) (bool, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return false, err
	}
	if user == nil || !user.IsMember() {
		return false, errors.New("user is not a valid member")
	}
	return true, nil
}

func (s *UserService) CheckMembership(user *models.User) bool {
	// 呼叫 IsMember 方法
	return user.IsMember()
}
