package relations

import "time"

// AlbumArtist 定義多對多關聯表
type AlbumArtist struct {
	AlbumID   uint `gorm:"primaryKey"`
	ArtistID  uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
