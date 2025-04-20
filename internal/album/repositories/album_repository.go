package repositories

import (
	"mcbach/internal/album/models" // 修正匯入路徑

	"gorm.io/gorm"
)

type AlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) FindNewReleases(limit, offset int) ([]models.Album, error) {
	var albums []models.Album
	if err := r.db.Preload("Artists").
		Order("release_at asc").
		Limit(limit).
		Offset(offset).
		Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *AlbumRepository) FindBySpotifyId(spotifyId string) (*models.Album, error) {
	var album models.Album
	if err := r.db.Preload("Artists").
		Where("spotify_album_id = ?", spotifyId).
		First(&album).Error; err != nil {
		return nil, err
	}
	return &album, nil
}

func (r *AlbumRepository) CreateMany(albums []models.Album) error {
	return r.db.Create(&albums).Error
}
