package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mcbach/internal/spotify/models"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var ctx = context.Background() // 定義全域的 context

type SpotifyService struct {
	clientID     string
	clientSecret string
	baseURL      string
	tokenURL     string
	cache        sync.Map // 使用內存快取
}

func NewSpotifyService() *SpotifyService {
	// 載入 .env 檔案
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env 檔案未找到，將使用系統環境變數: %v", err)
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatalf("SPOTIFY_CLIENT_ID 或 SPOTIFY_CLIENT_SECRET 未設定")
	}

	return &SpotifyService{
		clientID:     clientID,
		clientSecret: clientSecret,
		baseURL:      "https://api.spotify.com/v1",
		tokenURL:     "https://accounts.spotify.com/api/token",
	}
}

// RenewToken 更新 Spotify 的存取權杖
func (s *SpotifyService) RenewToken() (string, error) {
	// 準備請求資料
	requestBody := []byte("grant_type=client_credentials")
	req, err := http.NewRequest("POST", s.tokenURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// 設定標頭
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(s.clientID+":"+s.clientSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authHeader)

	// 發送請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to renew token")
	}

	// 解析回應
	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// 將 Token 存入快取
	expiration := time.Duration(result.ExpiresIn) * time.Second
	s.cache.Store("accessToken", result.AccessToken)
	s.cache.Store("accessTokenExpiration", time.Now().Add(expiration))

	return result.AccessToken, nil
}

// GetAccessToken 從內存快取獲取存取權杖，若不存在則更新
func (s *SpotifyService) GetAccessToken() (string, error) {
	// 檢查快取是否有有效的 Token
	if token, ok := s.cache.Load("accessToken"); ok {
		if expiration, ok := s.cache.Load("accessTokenExpiration"); ok {
			if time.Now().Before(expiration.(time.Time)) {
				return token.(string), nil
			}
		}
	}

	// 如果沒有有效的 Token，重新取得
	return s.RenewToken()
}

// GetNewReleases 從 Spotify API 獲取新專輯
func (s *SpotifyService) GetNewReleases(limit, offset int) ([]models.SpotifyAlbum, error) {
	token, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/browse/new-releases?limit=%d&offset=%d", s.baseURL, limit, offset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch new releases")
	}

	var result struct {
		Albums struct {
			Items []models.SpotifyAlbum `json:"items"`
		} `json:"albums"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Albums.Items, nil
}
