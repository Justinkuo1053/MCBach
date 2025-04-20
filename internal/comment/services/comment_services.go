package services

import (
	"mcbach/internal/comment/models"
	"mcbach/internal/comment/repositories"
)

type CommentService struct {
	repo *repositories.CommentRepository
}

func NewCommentService(repo *repositories.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) GetCommentsByAlbumID(albumID uint) ([]models.Comment, error) {
	return s.repo.FindByAlbumID(albumID)
}

func (s *CommentService) GetCommentByID(commentID uint) (*models.Comment, error) {
	return s.repo.FindByID(commentID)
}

func (s *CommentService) CreateComment(userID uint, albumID uint, content string) error {
	comment := &models.Comment{
		UserID:  userID,
		AlbumID: albumID,
		Content: content,
	}
	return s.repo.Create(comment)
}

func (s *CommentService) EditComment(commentID uint, content string) error {
	updates := map[string]interface{}{
		"content": content,
	}
	return s.repo.Update(commentID, updates)
}

func (s *CommentService) DeleteComment(commentID uint) error {
	return s.repo.SoftDelete(commentID)
}
