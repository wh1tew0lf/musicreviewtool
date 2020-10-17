package models

import (
	"fmt"
	u "musicreviewtool/utils"

	"github.com/jinzhu/gorm"
)

type ArtistRating struct {
	gorm.Model
	Rating   uint `json:"rating"`
	ArtistId uint `json:"artost_id" gorm:"primaryKey;autoIncrement:false"`
	UserId   uint `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
}

func (artistRating *ArtistRating) Validate() (map[string]interface{}, bool) {
	if artistRating.Rating <= 0 {
		return u.Message(false, "Incorrect rating"), false
	}

	if artistRating.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if artistRating.ArtistId <= 0 {
		return u.Message(false, "Artist is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (artistRating *ArtistRating) Create() map[string]interface{} {
	if resp, ok := artistRating.Validate(); !ok {
		return resp
	}

	GetDB().Create(artistRating)

	resp := u.Message(true, "success")
	resp["artistRating"] = artistRating
	return resp
}

func GetArtistRating(id uint) *ArtistRating {
	artistRating := &ArtistRating{}
	err := GetDB().Table("artist_rating").Where("id = ?", id).First(artistRating).Error
	if err != nil {
		return nil
	}
	return artistRating
}

func GetArtistRatings(user uint) []*ArtistRating {
	artistRatings := make([]*ArtistRating, 0)
	err := GetDB().Table("artist_rating").Where("user_id = ?", user).Find(&artistRatings).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return artistRatings
}
