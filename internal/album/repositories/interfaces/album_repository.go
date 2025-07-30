package interfaces

import (
    "mcbach/internal/album/models"
)

type AlbumRepository interface {
    // 基本 CRUD
    Create(album *models.Album) error
    FindByID(id uint) (*models.Album, error)
    Update(album *models.Album) error
    Delete(id uint) error

    // 特定查詢
    FindNewReleases(limit, offset int) ([]models.Album, error)
    FindBySpotifyId(spotifyId string) (*models.Album, error)
    CreateMany(albums []models.Album) error
}
