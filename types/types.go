package types

import "gorm.io/gorm"

type Artist struct {
	gorm.Model `json:"-"`
	ArtistID   string `gorm:"primaryKey" json:"artist_id"`
	Name       string `json:"name"`
}

type Song struct {
	gorm.Model `json:"-"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	AlbumURI   string    `json:"album_uri"`
	Artist     []*Artist `gorm:"many2many:artist_song;" json:"artist"`
}
