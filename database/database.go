package database

import (
	"errors"
	"fmt"

	"github.com/saminahbab/media-group-test/config"
	"github.com/saminahbab/media-group-test/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SongDatabaseClient struct {
	DB *gorm.DB
}

func NewSongDatabaseClient(cfg *config.Config) (*SongDatabaseClient, error) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Note: just for test: set up schema on startup
	err = db.AutoMigrate(&types.Song{}, &types.Artist{})
	if err != nil {
		return nil, err
	}

	return &SongDatabaseClient{
		DB: db,
	}, nil
}

func (s *SongDatabaseClient) SaveSong(song *types.Song) error {
	for _, artist := range song.Artist {
		err := s.DB.FirstOrCreate(artist, artist).Error
		if err != nil {
			return err
		}
	}
	err := s.DB.FirstOrCreate(song, song).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *SongDatabaseClient) GetSongByID(id string) (*types.Song, error) {
	var song types.Song
	err := s.DB.Preload("Artist").First(&song, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &song, nil

}

func (s *SongDatabaseClient) GetArtistByString(name string) ([]*types.Artist, error) {
	var artists []*types.Artist
	err := s.DB.Where("name LIKE ?", "%"+name+"%").Find(&artists).Error
	if err != nil {
		return nil, err
	}

	return artists, err
}
