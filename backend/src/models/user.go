package models

import (
	u "musicreviewtool/utils"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
	Token    string `json:"token" sql:"-"`
}

func (model *User) Validate() (map[string]interface{}, bool) {
	var validate *validator.Validate = validator.New()

	validateErr := validate.Struct(model)

	if validateErr != nil {
		for _, e := range validateErr.(validator.ValidationErrors) {
			return u.Message(false, e.Error()), false
		}
	}

	temp := &User{}

	err := GetDB().Where("email = ?", model.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another model."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (model *User) Create() map[string]interface{} {
	if resp, ok := model.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	model.Password = string(hashedPassword)

	GetDB().Create(model)

	if model.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	tk := &Token{UserId: model.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	model.Token = tokenString

	model.Password = ""

	response := u.Message(true, "User has been created")
	response["user"] = model
	return response
}

func Login(email, password string) map[string]interface{} {
	model := &User{}
	err := GetDB().Where("email = ?", email).First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	model.Password = ""

	tk := &Token{UserId: model.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	model.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["model"] = model
	return resp
}

func GetUser(u uint) *User {
	model := &User{}
	GetDB().First(model, u)
	if model.Email == "" {
		return nil
	}

	model.Password = ""
	return model
}
