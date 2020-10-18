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

func (Rating *Rating) Create() map[string]interface{} {
	if resp, ok := Rating.Validate(); !ok {
		return resp
	}

	GetDB().Create(Rating)

	resp := u.Message(true, "success")
	resp["Rating"] = Rating
	return resp
}
