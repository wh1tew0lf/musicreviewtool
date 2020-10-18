package models

import (
	"fmt"
	u "musicreviewtool/utils"

	"github.com/jinzhu/gorm"
)

type AlbumRating struct {
	gorm.Model
	Rating  uint `json:"rating" validate:"required,gte=0,lte=5"`
	AlbumId uint `json:"album_id" gorm:"primaryKey;autoIncrement:false" validate:"required"`
	UserId  uint `json:"user_id" gorm:"primaryKey;autoIncrement:false" validate:"required"`
}

func (AlbumRating *AlbumRating) Validate() (map[string]interface{}, bool) {
	if AlbumRating.Rating <= 0 {
		return u.Message(false, "Incorrect rating"), false
	}

	if AlbumRating.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if AlbumRating.AlbumId <= 0 {
		return u.Message(false, "Album is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (AlbumRating *AlbumRating) Create() map[string]interface{} {
	if resp, ok := AlbumRating.Validate(); !ok {
		return resp
	}

	GetDB().Create(AlbumRating)

	resp := u.Message(true, "success")
	resp["AlbumRating"] = AlbumRating
	return resp
}

func GetAlbumRating(id uint) *AlbumRating {
	AlbumRating := &AlbumRating{}
	err := GetDB().Table("album_rating").Where("id = ?", id).First(AlbumRating).Error
	if err != nil {
		return nil
	}
	return AlbumRating
}

func GetAlbumRatings(user uint) []*AlbumRating {
	AlbumRatings := make([]*AlbumRating, 0)
	err := GetDB().Table("album_rating").Where("user_id = ?", user).Find(&AlbumRatings).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return AlbumRatings
}
