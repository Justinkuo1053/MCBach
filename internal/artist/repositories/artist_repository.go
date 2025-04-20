package repositories

import (
	"mcbach/internal/artist/models"  // 修改這邊

	"gorm.io/gorm"
)

type ArtistRepository struct {
	db *gorm.DB
}

func NewArtistRepository(db *gorm.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) FindBySpotifyId(spotifyId string) (*models.Artist, error) {
	var artist models.Artist
	// 使用 Preload 載入 Albums 並只選取部分欄位，模擬 NestJS 中 select 的效果
	if err := r.db.
		Preload("Albums", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "title", "spotify_album_id")
		}).
		Where("spotify_artist_id = ?", spotifyId).
		First(&artist).Error; err != nil {
		return nil, err
	}
	return &artist, nil
}

func (r *ArtistRepository) CreateMany(artists []models.Artist) error {
	return r.db.Create(&artists).Error
}
