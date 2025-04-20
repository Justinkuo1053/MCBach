package controllers

import (
	"net/http"

	"mcbach/internal/artist/models"
	"mcbach/internal/artist/services"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	service *services.ArtistService
}

func NewArtistController(service *services.ArtistService) *ArtistController {
	return &ArtistController{service: service}
}

// GET /artist/:spotifyId
func (ac *ArtistController) GetArtistBySpotifyId(c *gin.Context) {
	spotifyId := c.Param("spotifyId")
	artist, err := ac.service.GetArtistBySpotifyId(spotifyId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		return
	}
	c.JSON(http.StatusOK, artist)
}

// POST /artist
// 預期傳入的 JSON 為陣列，例如：[{"name": "Artist name", "spotifyArtistId": "xxx"},...]
func (ac *ArtistController) CreateManyArtists(c *gin.Context) {
	var artists []models.Artist
	if err := c.ShouldBindJSON(&artists); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ac.service.CreateManyArtists(artists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "artists created"})
}
