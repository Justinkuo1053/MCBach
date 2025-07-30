package repositories

import (
	"errors"
	"log"
	"mcbach/internal/comment/models"
	"time"

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

func (r *CommentRepository) AddLike(like *models.Like) error {
	var existing models.Like
	if err := r.db.Where("user_id = ? AND comment_id = ?", like.UserID, like.CommentID).First(&existing).Error; err == nil {
		return errors.New("already liked")
	}
	return r.db.Create(like).Error
}

func (r *CommentRepository) IncrementLikeCount(commentID uint, isPro bool) error {
	field := "likes_count"
	if isPro {
		field = "pro_likes_count"
	}
	return r.db.Model(&models.Comment{}).Where("id = ?", commentID).Update(field, gorm.Expr(field+" + ?", 1)).Error
}

func (r *CommentRepository) GetTopUsersByLikes(start, end time.Time, limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	query := `
        SELECT user_id, SUM(likes_count + pro_likes_count * 2) AS total_score
        FROM comments
        WHERE created_at BETWEEN ? AND ?
        GROUP BY user_id
        ORDER BY total_score DESC
        LIMIT ?
    `
	if err := r.db.Raw(query, start, end, limit).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *CommentRepository) DeleteComment(commentID uint) error {
	return r.db.Where("id = ?", commentID).Delete(&models.Comment{}).Error
}

func (r *CommentRepository) DeleteLikesByCommentID(commentID uint) error {
	return r.db.Where("comment_id = ?", commentID).Delete(&models.Like{}).Error
}

func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	err := r.db.Save(comment).Error
	if err != nil {
		log.Printf("Save Error: %v", err)
	}
	return err
}

func (r *CommentRepository) GetCommentByID(id uint) (*models.Comment, error) {
	// 從資料庫獲取 Comment
	var comment models.Comment
	if err := r.db.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
