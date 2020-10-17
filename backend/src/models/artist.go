package models

import (
	"fmt"
	u "musicreviewtool/utils"

	"github.com/jinzhu/gorm"
)

type Artist struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (artist *Artist) Validate() (map[string]interface{}, bool) {
	if artist.Name == "" {
		return u.Message(false, "Artist name should be on the payload"), false
	}

	if artist.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if artist.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (artist *Artist) Create() map[string]interface{} {
	if resp, ok := artist.Validate(); !ok {
		return resp
	}

	GetDB().Create(artist)

	resp := u.Message(true, "success")
	resp["artist"] = artist
	return resp
}

func GetArtist(id uint) *Artist {
	artist := &Artist{}
	err := GetDB().Table("artists").Where("id = ?", id).First(artist).Error
	if err != nil {
		return nil
	}
	return artist
}

func GetArtists(user uint) []*Artist {
	artists := make([]*Artist, 0)
	err := GetDB().Table("artists").Where("user_id = ?", user).Find(&artists).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return artists
}
