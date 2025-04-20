package controllers

import (
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
	userID := c.GetUint("userID") // 假設 JWT 中的 userID 已解碼
	if err := cc.commentService.CreateComment(userID, req.AlbumID, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment created"})
}

func (cc *CommentController) EditComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.commentService.EditComment(uint(commentID), req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated"})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	if err := cc.commentService.DeleteComment(uint(commentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
