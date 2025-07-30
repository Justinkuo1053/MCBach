package controllers

import (
	"mcbach/internal/spotify/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SpotifyController struct {
	spotifyService *services.SpotifyService
}

func NewSpotifyController(spotifyService *services.SpotifyService) *SpotifyController {
	return &SpotifyController{spotifyService: spotifyService}
}

func (sc *SpotifyController) GetNewReleases(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	albums, err := sc.spotifyService.GetNewReleases(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}
