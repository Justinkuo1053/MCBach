package services

import (
	"mcbach/internal/album/models"       // 修正匯入路徑
	"mcbach/internal/album/repositories" // 修正匯入路徑
)

type AlbumService struct {
	repo *repositories.AlbumRepository
}

func NewAlbumService(repo *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) GetNewReleases(limit, offset int) ([]models.Album, error) {
	return s.repo.FindNewReleases(limit, offset)
}

func (s *AlbumService) GetAlbumBySpotifyId(spotifyId string) (*models.Album, error) {
	return s.repo.FindBySpotifyId(spotifyId)
}

func (s *AlbumService) CreateManyAlbums(albums []models.Album) error {
	return s.repo.CreateMany(albums)
}
