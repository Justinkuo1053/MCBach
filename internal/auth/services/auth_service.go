// filepath: auth/services/auth_service.go
package services

import (
	"errors"
	"mcbach/internal/auth/models"       // 修改為正確的路徑
	"mcbach/internal/auth/repositories" // 修改為正確的路徑
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo       *repositories.UserRepository
	jwtSecret  string
	jwtExpires time.Duration
}

func NewAuthService(repo *repositories.UserRepository, jwtSecret string, jwtExpires time.Duration) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret, jwtExpires: jwtExpires}
}

func (s *AuthService) Signup(email, password string) (*models.User, error) {
	// 檢查是否已存在相同 email 的使用者
	if _, err := s.repo.FindByEmail(email); err == nil {
		return nil, errors.New("email already in use")
	}

	// 雜湊密碼
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 建立新使用者
	user := &models.User{
		Email: email,
		Hash:  string(hash),
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Signin(email, password string) (string, error) {
	// 查詢使用者
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 驗證密碼
	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// 簽發 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(s.jwtExpires).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
