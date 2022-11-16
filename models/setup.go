package models

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectToDatabase(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env %s",err.Error())
	}
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	DB, err = gorm.Open("mysql", DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database mysql")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database mysql", )
	}
	DB.AutoMigrate(&User{})

}