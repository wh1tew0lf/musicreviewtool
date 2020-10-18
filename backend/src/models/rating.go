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

	for _, e := range validateErr.(validator.ValidationErrors) {
		return u.Message(false, e.Error()), false
	}

	return u.Message(true, "success"), true
}

func (Rating *Rating) RateArtist(id uint8) map[string]interface{} {
	if resp, ok := Rating.Validate(); !ok {
		return resp
	}

	// model := &Rating{}
	// err := GetDB().Where("external_table='artists' AND external_id = ?", id).Find(model).Error
	// if err != nil {
	// 	return nil
	// }

	resp := u.Message(true, "success")
	// resp["rating"] = model
	return resp
}

func (Rating *Rating) RateAlbum(id uint8) map[string]interface{} {
	if resp, ok := Rating.Validate(); !ok {
		return resp
	}

	// model := &Rating{}
	// err := GetDB().Where("external_table='albums' AND external_id = ?", id).Find(model).Error
	// if err != nil {
	// 	return nil
	// }

	resp := u.Message(true, "success")
	// resp["rating"] = model
	return resp
}
