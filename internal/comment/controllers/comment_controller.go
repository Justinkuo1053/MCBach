package controllers

import (
	"log"
	"mcbach/internal/comment/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{commentService: service}
}

func (cc *CommentController) GetCommentsByAlbumID(c *gin.Context) {
	albumID, _ := strconv.Atoi(c.Query("albumId"))
	comments, err := cc.commentService.GetCommentsByAlbumID(uint(albumID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	comment, err := cc.commentService.GetCommentByID(uint(commentID))
	log.Printf("Comment: %+v, Error: %v", comment, err)
	log.Printf("Error fetching comment: %v", err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var req struct {
		AlbumID uint   `json:"albumId" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Request Body: %+v", req)
	userID := c.GetUint("userID") // 假設 JWT 中的 userID 已解碼
	if err := cc.commentService.CreateComment(userID, req.AlbumID, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment created"})
}

func (cc *CommentController) EditComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID") // 假設 JWT 中的 userID 已解碼
	log.Printf("userID: %d, commentID: %d", userID, commentID)
	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("Updating Comment: %+v", req) // 修正為 req
	if err := cc.commentService.EditComment(userID, uint(commentID), req.Content); err != nil {
		log.Printf("EditComment Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment edited"})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID := c.Param("id") // 從路由參數中獲取評論 ID

	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment ID is required"})
		return
	}

	err := cc.commentService.DeleteCommentByID(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func (cc *CommentController) AddLike(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID") // 假設 JWT 中的 userID 已解碼
	if err := cc.commentService.AddLike(userID, uint(commentID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like added"})
}

func (cc *CommentController) AddProLike(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID") // 假設 JWT 中的 userID 已解碼
	if err := cc.commentService.AddProLike(userID, uint(commentID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pro like added"})
}

func (cc *CommentController) GetTopUsers(c *gin.Context) {
	// 調用服務層方法以獲取當月前十名使用者
	topUsers, err := cc.commentService.GetTopUsersByLikes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topUsers)
}
