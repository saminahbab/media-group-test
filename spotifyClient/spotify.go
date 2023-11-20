package spotifyClient

import (
	"context"

	"github.com/saminahbab/media-group-test/config"
	"github.com/saminahbab/media-group-test/types"
	"github.com/zmb3/spotify"
	oauth "golang.org/x/oauth2/clientcredentials"
)

type SpotifyClient struct {
	Client spotify.Client
}

func NewSpotifyClient(cfg *config.Config) (*SpotifyClient, error) {

	authConfig := &oauth.Config{
		ClientID:     cfg.SpotifyClientID,
		ClientSecret: cfg.SpotifyClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	// Get a token
	token, err := authConfig.Token(context.Background())
	if err != nil {
		return nil, err
	}

	// Create a new Spotify client
	client := spotify.Authenticator{}.NewClient(token)

	return &SpotifyClient{
		Client: client,
	}, nil

}

func (c *SpotifyClient) GetSongByID(isrc string) (*types.Song, error) {
	result, err := c.Client.Search("isrc:"+isrc, spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}

	if len(result.Tracks.Tracks) == 0 {
		return nil, nil
	}

	song := parseSpotifyResult(result, isrc)
	return song, nil
}

func parseSpotifyResult(result *spotify.SearchResult, isrc string) *types.Song {
	// if the result is more than one, use the most popular
	index := 0
	if len(result.Tracks.Tracks) > 1 {
		index = findMostPopular(result)
	}

	toParse := result.Tracks.Tracks[index]
	artists := make([]*types.Artist, 0)
	for _, artist := range toParse.Artists {
		artists = append(artists, &types.Artist{Name: artist.Name, ArtistID: string(artist.ID)})
	}
	song := &types.Song{
		ID:       isrc,
		Name:     toParse.Name,
		AlbumURI: toParse.Album.Images[0].URL,
		Artist:   artists,
	}

	return song
}

func findMostPopular(result *spotify.SearchResult) int {
	index := 0
	popularity := result.Tracks.Tracks[0].Popularity
	for i := 1; i < len(result.Tracks.Tracks); i++ {
		if result.Tracks.Tracks[i].Popularity > popularity {
			popularity = result.Tracks.Tracks[index].Popularity
			index = i
		}
	}

	return index
}
