package models

import (
	u "musicreviewtool/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Rating struct {
	gorm.Model
	Rating        uint   `json:"rating" validate:"required,gte=0,lte=5"`
	ExternalTable string `json:"-"`
	ExternalId    uint   `json:"external_id" gorm:"primaryKey;autoIncrement:false" validate:"required"`
	UserId        uint   `json:"user_id" gorm:"primaryKey;autoIncrement:false" validate:"required"`
}

func (model *Rating) Validate() (map[string]interface{}, bool) {
	var validate *validator.Validate = validator.New()

	validateErr := validate.Struct(model)
	if validateErr != nil {
		for _, e := range validateErr.(validator.ValidationErrors) {
			return u.Message(false, e.Error()), false
		}
	}

	return u.Message(true, "success"), true
}

func (model *Rating) RateArtist(user_id uint, id uint) map[string]interface{} {
	if resp, ok := model.Validate(); !ok {
		return resp
	}

	temp := &Rating{}
	err := GetDB().Where("external_table='artists' AND external_id = ?", id).Find(temp).Error
	if err != nil {
		return nil
	}

	if temp.ID > 0 {
		temp.Rating = model.Rating
		GetDB().Save(temp)
	} else {
		GetDB().Save(model)
	}

	resp := u.Message(true, "success")
	resp["rating"] = model
	return resp
}

func (model *Rating) RateAlbum(id uint8) map[string]interface{} {
	if resp, ok := model.Validate(); !ok {
		return resp
	}

	temp := &Rating{}
	err := GetDB().Where("external_table='albums' AND external_id = ?", id).Find(temp).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "success")
	resp["rating"] = model
	return resp
}
