package models

type SpotifyAlbum struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Artists     []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"artists"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}
