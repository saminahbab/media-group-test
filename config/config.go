package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresPassword    string
	PostgresDatabase    string
	PostgresUsername    string
	SpotifyClientID     string
	SpotifyClientSecret string
}

func LoadConfig() (*Config, error) {
	// Initialize a new Viper instance
	viper := viper.New()

	// Set environment variable prefixes
	viper.AutomaticEnv()

	// Set default values or required keys
	// NOTE: THESE CREDENTIALS SHOULD NOT BE HARD CODED IN PRODUCTION
	viper.SetDefault("POSTGRES_PASSWORD", "test")
	viper.SetDefault("POSTGRES_DATABASE", "spotify")
	viper.SetDefault("POSTGRES_USERNAME", "test")
	viper.SetDefault("SPOTIFY_CLIENT_ID", "")
	viper.SetDefault("SPOTIFY_CLIENT_SECRET", "")

	// Bind the environment variables to the Config struct
	config := &Config{
		PostgresPassword:    viper.GetString("POSTGRES_PASSWORD"),
		PostgresDatabase:    viper.GetString("POSTGRES_DATABASE"),
		PostgresUsername:    viper.GetString("POSTGRES_USERNAME"),
		SpotifyClientID:     viper.GetString("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: viper.GetString("SPOTIFY_CLIENT_SECRET"),
	}

	if len(config.SpotifyClientID) == 0 || len(config.SpotifyClientSecret) == 0 {
		return nil, errors.New("No spotify credentials set")
	}
	return config, nil

}
