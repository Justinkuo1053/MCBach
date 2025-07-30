package main

import (
	"log"
	"mcbach/internal/artist/repositories"
	"mcbach/internal/artist/services"
	"mcbach/internal/config"
	spotifyServices "mcbach/internal/spotify/services"
	"os"
	"testing"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("無法取得工作目錄: %v", err)
	}
	log.Printf("當前工作目錄: %s", wd)
}
func TestArtistService(t *testing.T) {
	db := config.ConnectDB()

	artistRepo := repositories.NewArtistRepository(db)
	artistService := services.NewArtistService(artistRepo)

	artist, err := artistService.GetArtistBySpotifyId("some-spotify-id")
	if err != nil {
		log.Printf("Artist not found: %v", err)
	} else {
		log.Printf("Found artist: %+v", artist)
	}
}

func TestDatabaseConnection(t *testing.T) {
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("無法取得資料庫連線: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("無法連接到資料庫: %v", err)
	}
	log.Println("資料庫連線測試成功！")
}

func TestSpotifyNewReleases(t *testing.T) {
	// 初始化 SpotifyService
	spotifyService := spotifyServices.NewSpotifyService()

	// 測試取得新專輯
	albums, err := spotifyService.GetNewReleases(5, 0)
	if err != nil {
		t.Fatalf("Failed to fetch new releases: %v", err)
	}

	t.Logf("Fetched albums: %+v", albums)
}
