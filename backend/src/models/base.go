package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("APP_DB_USER")
	password := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("APP_DB_NAME")
	dbHost := os.Getenv("APP_DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	connection, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	connection.AutoMigrate(&User{}, &Artist{}, &Album{}, &Rating{})
}

func GetDB() *gorm.DB {
	return db
}
