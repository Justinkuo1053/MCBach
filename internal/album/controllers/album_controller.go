package controllers

import (
	"mcbach/internal/album/services" // 修正匯入路徑
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumService *services.AlbumService
}

func NewAlbumController(albumService *services.AlbumService) *AlbumController {
	return &AlbumController{albumService: albumService}
}

func (ac *AlbumController) GetNewReleases(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	albums, err := ac.albumService.GetNewReleases(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (ac *AlbumController) GetAlbumBySpotifyId(c *gin.Context) {
	spotifyId := c.Param("spotifyId")

	album, err := ac.albumService.GetAlbumBySpotifyId(spotifyId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (ac *AlbumController) GetAlbums(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	albums, err := ac.albumService.GetNewReleases(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}
