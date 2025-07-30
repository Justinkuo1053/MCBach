package controllers

import (
	"mcbach/internal/user/services" // 修改 import 路徑
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetMe(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID")) // 假設 JWT 中的 userID 已解碼
	user, err := uc.userService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) EditUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID")) // 假設 JWT 中的 userID 已解碼
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userService.EditUser(uint(userID), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
