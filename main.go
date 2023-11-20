package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saminahbab/media-group-test/config"
	"github.com/saminahbab/media-group-test/database"
	"github.com/saminahbab/media-group-test/service"
	"github.com/saminahbab/media-group-test/spotifyClient"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed To initialise Config. ", err.Error())
	}

	spotifyClient, err := spotifyClient.NewSpotifyClient(configuration)
	if err != nil {
		log.Fatal("Failed to initialise spotify Client. ", err.Error())
	}

	DatabaseClient, err := database.NewSongDatabaseClient(configuration)
	if err != nil {
		log.Fatal("Failed to initialise database client. ", err.Error())
	}

	svc := &service.SpotifySongService{
		Database:      DatabaseClient,
		SpotifyClient: spotifyClient,
	}

	router := gin.Default()

	router.PUT("/songs/:isrc", svc.SaveSong)
	router.GET("/songs/:isrc", svc.GetSongByID)
	router.GET("/artists/similar/:string", svc.GetArtistsLikeString)

	router.Run(":8080")
}
