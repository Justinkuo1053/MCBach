package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	albumModels "mcbach/internal/album/models"
	artistModels "mcbach/internal/artist/models"
	authControllers "mcbach/internal/auth/controllers"
	authModels "mcbach/internal/auth/models"
	authRepositories "mcbach/internal/auth/repositories"
	authServices "mcbach/internal/auth/services"
	commentControllers "mcbach/internal/comment/controllers"
	commentRepositories "mcbach/internal/comment/repositories"
	commentServices "mcbach/internal/comment/services"
	config "mcbach/internal/config"
	relations "mcbach/internal/relations"
	spotifyControllers "mcbach/internal/spotify/controllers"
	spotifyServices "mcbach/internal/spotify/services"
	userControllers "mcbach/internal/user/controllers"
	userRepositories "mcbach/internal/user/repositories"
	userServices "mcbach/internal/user/services"

	"github.com/gin-gonic/gin"

	albumControllers "mcbach/internal/album/controllers"
	albumServices "mcbach/internal/album/services"

	albumRepositories "mcbach/internal/album/repositories"
	artistControllers "mcbach/internal/artist/controllers"
	artistRepositories "mcbach/internal/artist/repositories"
	artistServices "mcbach/internal/artist/services"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "未提供授權碼", http.StatusBadRequest)
		return
	}

	clientID := "your_client_id"                    // 替換為你的 Client ID
	clientSecret := "your_client_secret"            // 替換為你的 Client Secret
	redirectURI := "http://127.0.0.1:8080/callback" // 與 Spotify 開發者平台設定一致

	// 準備請求資料
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, "無法建立請求", http.StatusInternalServerError)
		return
	}

	// 設定標頭
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	// 發送請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "無法發送請求", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 解析回應
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "無法讀取回應", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Spotify API 錯誤: %s", body), http.StatusInternalServerError)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "無法解析回應", http.StatusInternalServerError)
		return
	}

	// 顯示存取權杖
	accessToken := result["access_token"].(string)
	fmt.Fprintf(w, "Access Token: %s", accessToken)
}

// func getUserProfile(accessToken string) {
// 	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
// 	if err != nil {
// 		log.Fatalf("無法建立請求: %v", err)
// 	}

// 	req.Header.Set("Authorization", "Bearer "+accessToken)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("無法發送請求: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("Spotify API 錯誤: %v", resp.Status)
// 	}

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println("使用者資料:", string(body))
// }

func main() {
	// 初始化資料庫
	db := config.ConnectDB()

	// 測試資料庫連線
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection successful!")

	// 自動遷移資料庫模型
	if err := db.AutoMigrate(
		&artistModels.Artist{},
		&albumModels.Album{},
		&relations.AlbumArtist{},
		&authModels.User{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration successful!")

	// 初始化 Spotify 服務
	spotifyService := spotifyServices.NewSpotifyService()
	spotifyController := spotifyControllers.NewSpotifyController(spotifyService)

	// 初始化 User 服務
	userRepo := userRepositories.NewUserRepository(db)
	userService := userServices.NewUserService(userRepo)
	userController := userControllers.NewUserController(userService)

	// 初始化 Comment 服務
	commentRepo := commentRepositories.NewCommentRepository(db)
	commentService := commentServices.NewCommentService(commentRepo)
	commentController := commentControllers.NewCommentController(commentService)

	// 初始化 Auth 服務
	authRepo := authRepositories.NewUserRepository(db)
	authService := authServices.NewAuthService(authRepo, "your_jwt_secret", time.Hour*24)
	authController := authControllers.NewAuthController(authService)

	// 初始化 Album 服務
	albumRepo := albumRepositories.NewAlbumRepository(db)    // 建立 AlbumRepository
	albumService := albumServices.NewAlbumService(albumRepo) // 傳入 AlbumRepository
	albumController := albumControllers.NewAlbumController(albumService)

	// 初始化 Artist 服務
	artistRepo := artistRepositories.NewArtistRepository(db) // 建立 ArtistRepository
	artistService := artistServices.NewArtistService(artistRepo)
	artistController := artistControllers.NewArtistController(artistService)

	// 初始化 Gin 路由
	r := gin.Default()

	// 註冊 Spotify 路由
	spotifyRoutes := r.Group("/spotify")
	{
		spotifyRoutes.GET("/new-releases", spotifyController.GetNewReleases)
	}

	// 註冊 User 路由
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/me", userController.GetMe)
		userRoutes.PUT("/me", userController.EditUser)
	}

	// 註冊 Comment 路由
	commentRoutes := r.Group("/comments")
	{
		commentRoutes.GET("", commentController.GetCommentsByAlbumID)
		commentRoutes.GET("/:id", commentController.GetCommentByID)
		commentRoutes.POST("", commentController.CreateComment)
		commentRoutes.PUT("/:id", commentController.EditComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
	}

	// 註冊 Auth 路由
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", authController.Signup)
		authRoutes.POST("/signin", authController.Signin)
	}

	// 註冊 Album 路由
	albumRoutes := r.Group("/albums")
	{
		albumRoutes.GET("", albumController.GetAlbums)
	}

	// 註冊 Artist 路由
	artistRoutes := r.Group("/artists")
	{
		artistRoutes.GET("/:spotifyId", artistController.GetArtistBySpotifyId)
		artistRoutes.POST("", artistController.CreateManyArtists)
	}

	// 註冊 Callback 路由
	http.HandleFunc("/callback", callbackHandler)

	// 啟動伺服器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("當前時間: %s", time.Now())
	log.Printf("當前時區: %s", time.Now().Location())
}
