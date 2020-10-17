package models

import (
	"fmt"
	u "musicreviewtool/utils"

	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	ArtistId uint   `json:"artost_id"`
	UserId   uint   `json:"user_id"`
}

func (album *Album) Validate() (map[string]interface{}, bool) {
	if album.Name == "" {
		return u.Message(false, "Album name should be on the payload"), false
	}

	if album.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if album.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if album.ArtistId <= 0 {
		return u.Message(false, "Artist is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (album *Album) Create() map[string]interface{} {
	if resp, ok := album.Validate(); !ok {
		return resp
	}

	GetDB().Create(album)

	resp := u.Message(true, "success")
	resp["album"] = album
	return resp
}

func GetAlbum(id uint) *Album {
	album := &Album{}
	err := GetDB().Table("albums").Where("id = ?", id).First(album).Error
	if err != nil {
		return nil
	}
	return album
}

func GetAlbums(user uint) []*Album {
	albums := make([]*Album, 0)
	err := GetDB().Table("albums").Where("user_id = ?", user).Find(&albums).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return albums
}
