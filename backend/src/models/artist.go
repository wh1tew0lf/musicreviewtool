package models

import (
	"fmt"
	u "musicreviewtool/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Artist struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      uint   `json:"user_id" validate:"required"`
}

func (model *Artist) Validate() (map[string]interface{}, bool) {
	var validate *validator.Validate = validator.New()

	validateErr := validate.Struct(model)

	for _, e := range validateErr.(validator.ValidationErrors) {
		return u.Message(false, e.Error()), false
	}

	return u.Message(true, "success"), true
}

func (model *Artist) Create() map[string]interface{} {
	if resp, ok := model.Validate(); !ok {
		return resp
	}

	GetDB().Create(model)

	resp := u.Message(true, "success")
	resp["artist"] = model
	return resp
}

func GetArtist(id uint) *Artist {
	artist := &Artist{}
	err := GetDB().Where("id = ?", id).First(artist).Error
	if err != nil {
		return nil
	}
	return artist
}

func GetAllArtists(limit uint, offset uint) []*Artist {
	artists := make([]*Artist, 0)
	err := GetDB().Limit(limit).Offset(offset).Find(&artists).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return artists
}

func (artist *Artist) Update() map[string]interface{} {
	if resp, ok := artist.Validate(); !ok {
		return resp
	}

	GetDB().Save(artist)

	resp := u.Message(true, "success")
	resp["artist"] = artist
	return resp
}

func (artist *Artist) Delete() map[string]interface{} {
	GetDB().Delete(artist)

	resp := u.Message(true, "success")
	return resp
}
