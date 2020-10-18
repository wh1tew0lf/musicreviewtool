package models

import (
	"fmt"
	u "musicreviewtool/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Title    string `json:"name" validate:"required"`
	Year     uint   `json:"year" validate:"required"`
	ArtistId uint   `json:"artist_id" validate:"required"`
	UserId   uint   `json:"user_id" validate:"required"`
}

func (model *Album) Validate() (map[string]interface{}, bool) {
	var validate *validator.Validate = validator.New()

	validateErr := validate.Struct(model)

	for _, e := range validateErr.(validator.ValidationErrors) {
		return u.Message(false, e.Error()), false
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
	err := GetDB().Where("id = ?", id).First(album).Error
	if err != nil {
		return nil
	}
	return album
}

func GetAllAlbums(artistId uint, limit uint, offset uint) []*Album {
	albums := make([]*Album, 0)
	err := GetDB().Limit(limit).Offset(offset).Where("artistId = ?", artistId).Find(&albums).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return albums
}

func (model *Album) Update() map[string]interface{} {
	if resp, ok := model.Validate(); !ok {
		return resp
	}

	GetDB().Save(model)

	resp := u.Message(true, "success")
	resp["album"] = model
	return resp
}

func (model *Album) Delete() map[string]interface{} {
	GetDB().Delete(model)

	resp := u.Message(true, "success")
	return resp
}
