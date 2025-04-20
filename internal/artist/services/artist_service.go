package services

import (
    "mcbach/internal/artist/models"
    "mcbach/internal/artist/repositories"  // 將 import 調整為正確路徑
)

type ArtistService struct {
    repo *repositories.ArtistRepository
}

func NewArtistService(repo *repositories.ArtistRepository) *ArtistService {
    return &ArtistService{repo: repo}
}

func (s *ArtistService) GetArtistBySpotifyId(spotifyId string) (*models.Artist, error) {
    return s.repo.FindBySpotifyId(spotifyId)
}

func (s *ArtistService) CreateManyArtists(artists []models.Artist) error {
    return s.repo.CreateMany(artists)
}