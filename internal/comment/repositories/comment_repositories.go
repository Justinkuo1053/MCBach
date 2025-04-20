package repositories

import (
	"mcbach/internal/comment/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) FindByAlbumID(albumID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("album_id = ? AND is_deleted = ?", albumID, false).
		Order("created_at desc").
		Preload("Likes").
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) FindByID(commentID uint) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.Where("id = ? AND is_deleted = ?", commentID, false).
		Preload("Likes").
		First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) Update(commentID uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Comment{}).Where("id = ? AND is_deleted = ?", commentID, false).
		Updates(updates).Error
}

func (r *CommentRepository) SoftDelete(commentID uint) error {
	return r.db.Model(&models.Comment{}).Where("id = ? AND is_deleted = ?", commentID, false).
		Update("is_deleted", true).Error
}
