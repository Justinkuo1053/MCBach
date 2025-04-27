package services

import (
	"errors"
	"log"
	"mcbach/internal/comment/models" // 為 user/models 使用別名
	"mcbach/internal/comment/repositories"
	"mcbach/internal/user/services" // 假設 UserService 在這個路徑
	"time"
)

type CommentService struct {
	repo        *repositories.CommentRepository
	userService *services.UserService // 新增 UserService 依賴
}

// NewCommentService: 建構函數，注入依賴
func NewCommentService(repo *repositories.CommentRepository, userService *services.UserService) *CommentService {
	return &CommentService{
		repo:        repo,
		userService: userService,
	}
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

func (s *CommentService) EditComment(userID, commentID uint, newContent string) error {
	// 確認該評論是否屬於該使用者
	comment, err := s.repo.GetCommentByID(commentID)
	log.Printf("Comment: %+v, Error: %v", comment, err)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		return errors.New("you can only edit your own comments")
	}

	// 更新評論內容、重置讚數並記錄編輯時間
	comment.Content = newContent
	comment.LikesCount = 0
	comment.ProLikesCount = 0
	comment.UpdatedAt = time.Now()
	comment.IsEdited = true

	log.Printf("Updating Comment: %+v", comment)

	return s.repo.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(userID, commentID uint) error {
	// 確認該評論是否屬於該使用者
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		return errors.New("you can only delete your own comments")
	}

	// 刪除評論及相關按讚記錄
	if err := s.repo.DeleteComment(commentID); err != nil {
		return err
	}
	return s.repo.DeleteLikesByCommentID(commentID)
}

// addLikeInternal: 私有輔助函數，處理按讚邏輯
func (s *CommentService) addLikeInternal(userID, commentID uint, isPro bool) error {
	// 驗證按讚者是否為有效會員
	isValid, err := s.userService.IsValidMember(userID)
	if err != nil {
		return errors.New("failed to validate member status: " + err.Error())
	}
	if !isValid {
		return errors.New("only valid members can like comments")
	}

	// 創建 Like 記錄
	like := &models.Like{
		UserID:    userID,
		CommentID: commentID,
		IsPro:     isPro,
	}

	// 嘗試新增 Like 記錄
	if err := s.repo.AddLike(like); err != nil {
		return err
	}

	// 更新讚數
	return s.repo.IncrementLikeCount(commentID, isPro)
}

// AddLike: 處理普通按讚請求
func (s *CommentService) AddLike(userID, commentID uint) error {
	return s.addLikeInternal(userID, commentID, false)
}

// AddProLike: 處理專業按讚請求
func (s *CommentService) AddProLike(userID, commentID uint) error {
	return s.addLikeInternal(userID, commentID, true)
}

func (s *CommentService) GetTopUsersByLikes() ([]map[string]interface{}, error) {
	// 計算當月的開始和結束時間
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 1, 0)

	// 調用 Repository 方法
	return s.repo.GetTopUsersByLikes(start, end, 10)
}
