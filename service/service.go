package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saminahbab/media-group-test/database"
	"github.com/saminahbab/media-group-test/spotifyClient"
)

type SpotifySongService struct {
	Database      *database.SongDatabaseClient
	SpotifyClient *spotifyClient.SpotifyClient
}

func (s *SpotifySongService) SaveSong(c *gin.Context) {
	id := c.Param("isrc")
	song, err := s.SpotifyClient.GetSongByID(id)
	if err != nil {
		// todo clean up the error handling to make it consistent with the rest of the file
		c.Error(err)
	}

	if song == nil {
		c.Status(http.StatusNotFound)
		return
	}

	err = s.Database.SaveSong(song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

}

func (s *SpotifySongService) GetSongByID(c *gin.Context) {
	id := c.Param("isrc")
	song, err := s.Database.GetSongByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if song == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": song,
	})

}

func (s *SpotifySongService) GetArtistsLikeString(c *gin.Context) {
	like := c.Param("string")
	// Code to get songs like a string
	artists, err := s.Database.GetArtistByString(like)
	if err != nil {
		log.Println("Error when getting artist by Like")
		c.Status(http.StatusInternalServerError)
		return
	}

	if len(artists) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": artists,
	})

}
